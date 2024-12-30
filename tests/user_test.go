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

func setup() {
	test_db.InitTestDB()
	web.BConfig.RunMode = "test"
	web.Router("/user", &controllers.UserController{}, "post:CreateUser")
	web.Router("/user/:id", &controllers.UserController{}, "put:UpdateUser;delete:DeleteUser")
	web.Router("/users", &controllers.UserController{}, "get:GetAllUsers")
}

// TestCreateUser Test CreateUser
func TestCreateUser(t *testing.T) {
	setup()

	formData := `username=testuser&password=testpass&email=testuser@example.com&role=admin`
	req, _ := http.NewRequest("POST", "/user", bytes.NewBufferString(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusFound, recorder.Code, "Expected status code 303 See Other")
}

// TestUpdateUser Test UpdateUser
func TestUpdateUser(t *testing.T) {
	setup()

	// Insert a test user
	o := orm.NewOrm()

	userId, err := o.Insert(&models.User{
		Username: "testuser",
		Password: "testpass",
		Email:    "testuser@example.com",
		Role:     "admin",
	})

	if err != nil {
		return
	}

	userIdString := strconv.Itoa(int(userId))

	updateData := `username=updateduser&email=updated@example.com&role=user`
	req, _ := http.NewRequest("PUT", "/user/"+userIdString, bytes.NewBufferString(updateData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusSeeOther, recorder.Code, "Expected status code 303 See Other")
}

// TestDeleteUser Test DeleteUser
func TestDeleteUser(t *testing.T) {
	setup()

	// Insert a test user
	o := orm.NewOrm()
	user := &models.User{
		Username: "testuser",
		Password: "testpass",
		Email:    "testuser@example.com",
		Role:     "admin",
	}
	id, err := o.Insert(user)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}
	t.Logf("Inserted user ID: %d", id)

	// Prepare DELETE request
	stringifiedUserID := strconv.Itoa(int(id))
	req, _ := http.NewRequest("DELETE", "/user/"+stringifiedUserID, nil)

	// Send request
	recorder := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(recorder, req)

	// Log response
	t.Logf("Response Code: %d, Body: %s", recorder.Code, recorder.Body.String())

	// Assertions
	assert.Equal(t, http.StatusSeeOther, recorder.Code, "Expected status code 303 See Other")

	// Verify deletion
	err = o.QueryTable("user").Filter("id", id).One(user)
	assert.Error(t, err, "Expected no user to exist")
}
