package tests

import (
	"bytes"
	"fullstack-beego-app/controllers"
	"fullstack-beego-app/models"
	"fullstack-beego-app/test_db"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func setupProducts() {
	test_db.InitTestDB()
	web.BConfig.RunMode = "test"
	web.Router("/product", &controllers.ProductController{}, "post:CreateProduct")
	web.Router("/product/:id", &controllers.ProductController{}, "put:UpdateProduct;delete:DeleteProduct")
	web.Router("/products", &controllers.ProductController{}, "get:GetAllProducts")
}

// TestCreateProduct Test CreateProduct
func TestCreateProduct(t *testing.T) {
	setupProducts()

	formData := `name=testproduct&sku=12345`
	req, _ := http.NewRequest("POST", "/product", bytes.NewBufferString(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusFound, recorder.Code, "Expected status code 303 See Other")
}

// TestUpdateProduct Test UpdateProduct
func TestUpdateProduct(t *testing.T) {
	setupProducts()

	// Insert a test product
	o := orm.NewOrm()

	productId, err := o.Insert(&models.Product{
		Name: "testproduct",
		SKU:  "12345",
	})

	if err != nil {
		return
	}

	productIdString := strconv.Itoa(int(productId))

	updateData := `name=updatedproduct&sku=67890`
	req, _ := http.NewRequest("PUT", "/product/"+productIdString, bytes.NewBufferString(updateData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusFound, recorder.Code, "Expected status code 303 See Other")
}

// TestDeleteProduct Test DeleteProduct
func TestDeleteProduct(t *testing.T) {
	setupProducts()

	// Insert a test product
	o := orm.NewOrm()
	product := &models.Product{
		Name: "testproduct",
		SKU:  "12345",
	}
	id, err := o.Insert(product)
	if err != nil {
		t.Fatalf("Failed to insert product: %v", err)
	}
	t.Logf("Inserted product ID: %d", id)

	// Prepare DELETE request
	stringifiedProductID := strconv.Itoa(int(id))
	req, _ := http.NewRequest("DELETE", "/product/"+stringifiedProductID, nil)

	// Send request
	recorder := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	// Log response
	t.Logf("Response Code: %d, Body: %s", recorder.Code, recorder.Body.String())

	// Assertions
	assert.Equal(t, http.StatusSeeOther, recorder.Code, "Expected status code 303 See Other")

	// Verify deletion
	err = o.QueryTable("product").Filter("id", id).One(product)
	assert.Error(t, err, "Expected no product to exist")
}
