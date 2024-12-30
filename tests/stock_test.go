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

func setupStocks() {
	test_db.InitTestDB()
	web.BConfig.RunMode = "test"
	web.Router("/stock", &controllers.StockController{}, "post:AddStock")
	web.Router("/stock/:id", &controllers.StockController{}, "put:UpdateStock;delete:DeleteStock")
	web.Router("/stocks", &controllers.StockController{}, "get:GetAllStocks")
}

// TestAddStock ensures the AddStock method works correctly
func TestAddStock(t *testing.T) {
	setupStocks()

	o := orm.NewOrm()

	// Insert related records
	warehouseID, _ := o.Insert(&models.Warehouse{Name: "Test Warehouse", Location: "Test Location"})
	productID, _ := o.Insert(&models.Product{Name: "Test Product", SKU: "SKU123"})

	formData := `warehouse_id=` + strconv.Itoa(int(warehouseID)) +
		`&product_id=` + strconv.Itoa(int(productID)) +
		`&quantity=100&threshold=10`
	req, _ := http.NewRequest("POST", "/stock", bytes.NewBufferString(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusFound, recorder.Code, "Expected status code 303 See Other")
}

// TestUpdateStock ensures the UpdateStock method works correctly
func TestUpdateStock(t *testing.T) {
	setupStocks()

	o := orm.NewOrm()

	// Insert related records
	warehouseID, _ := o.Insert(&models.Warehouse{Name: "Test Warehouse", Location: "Test Location"})
	productID, _ := o.Insert(&models.Product{Name: "Test Product", SKU: "SKU123"})

	// Insert a test stock
	stockID, _ := o.Insert(&models.Stock{
		Warehouse: &models.Warehouse{ID: int(warehouseID)},
		Product:   &models.Product{ID: int(productID)},
		Quantity:  100,
		Threshold: 10,
	})

	formData := `warehouse_id=` + strconv.Itoa(int(warehouseID)) +
		`&product_id=` + strconv.Itoa(int(productID)) +
		`&quantity=200&threshold=20`
	req, _ := http.NewRequest("PUT", "/stock/"+strconv.Itoa(int(stockID)), bytes.NewBufferString(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusSeeOther, recorder.Code, "Expected status code 303 See Other")
}

// TestDeleteStock ensures the DeleteStock method works correctly
func TestDeleteStock(t *testing.T) {
	setupStocks()

	o := orm.NewOrm()

	// Insert related records
	warehouseID, _ := o.Insert(&models.Warehouse{Name: "Test Warehouse", Location: "Test Location"})
	productID, _ := o.Insert(&models.Product{Name: "Test Product", SKU: "SKU123"})

	// Insert a test stock
	stockID, _ := o.Insert(&models.Stock{
		Warehouse: &models.Warehouse{ID: int(warehouseID)},
		Product:   &models.Product{ID: int(productID)},
		Quantity:  100,
		Threshold: 10,
	})

	req, _ := http.NewRequest("DELETE", "/stock/"+strconv.Itoa(int(stockID)), nil)
	recorder := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusSeeOther, recorder.Code, "Expected status code 303 See Other")

	// Verify deletion
	stock := models.Stock{ID: int(stockID)}
	err := o.Read(&stock)
	assert.Error(t, err, "Expected stock to be deleted")
}
