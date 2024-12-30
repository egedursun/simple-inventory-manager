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

func setupWarehouses() {
	test_db.InitTestDB()
	web.BConfig.RunMode = "test"
	web.Router("/warehouse", &controllers.WarehouseController{}, "post:CreateWarehouse")
	web.Router("/warehouse/:id", &controllers.WarehouseController{}, "put:UpdateWarehouse;delete:DeleteWarehouse")
	web.Router("/warehouses", &controllers.WarehouseController{}, "get:GetAllWarehouses")
}

// TestCreateWarehouse ensures the CreateWarehouse method works correctly
func TestCreateWarehouse(t *testing.T) {
	setupWarehouses()

	formData := `name=Test Warehouse&location=Test Location`
	req, _ := http.NewRequest("POST", "/warehouse", bytes.NewBufferString(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusFound, recorder.Code, "Expected status code 303 See Other")
}

// TestDeleteWarehouse ensures the DeleteWarehouse method works correctly
func TestDeleteWarehouse(t *testing.T) {
	setupWarehouses()

	// Insert a test warehouse
	o := orm.NewOrm()
	warehouse := &models.Warehouse{
		Name:     "Test Warehouse",
		Location: "Test Location",
	}
	id, err := o.Insert(warehouse)
	if err != nil {
		t.Fatalf("Failed to insert warehouse: %v", err)
	}
	t.Logf("Inserted warehouse ID: %d", id)

	// Prepare DELETE request
	stringifiedWarehouseID := strconv.Itoa(int(id))
	req, _ := http.NewRequest("DELETE", "/warehouse/"+stringifiedWarehouseID, nil)

	// Send request
	recorder := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	// Log response
	t.Logf("Response Code: %d, Body: %s", recorder.Code, recorder.Body.String())

	// Assertions
	assert.Equal(t, http.StatusFound, recorder.Code, "Expected status code 303 See Other")

	// Verify deletion
	err = o.QueryTable("warehouse").Filter("id", id).One(warehouse)
	assert.Error(t, err, "Expected no warehouse to exist")
}
