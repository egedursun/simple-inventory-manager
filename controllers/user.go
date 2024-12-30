package controllers

import (
	"fullstack-beego-app/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type UserController struct {
	beego.Controller
}

// CreateUser handles POST /user
func (c *UserController) CreateUser() {
	user := models.User{
		Username: c.GetString("username"),
		Password: c.GetString("password"),
		Email:    c.GetString("email"),
		Role:     c.GetString("role"),
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		c.Data["Error"] = "Username, password, and email are required."
		c.TplName = "users.tpl"
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&user); err != nil {
		c.Data["Error"] = "Failed to create user."
		c.TplName = "users.tpl"
		return
	}

	c.Redirect("/users", 302)
}

// UpdateUser handles PUT /user/:id
func (c *UserController) UpdateUser() {
	// Extract the ID from the route
	id, err := c.GetInt(":id")
	if err != nil {
		c.Data["Error"] = "Invalid user ID"
		c.Redirect("/users", 303)
		return
	}

	// Initialize ORM and retrieve the user by ID
	o := orm.NewOrm()
	user := models.User{ID: id}
	if err := o.Read(&user); err != nil {
		c.Data["Error"] = "User not found"
		c.Redirect("/users", 303)
		return
	}

	// Update fields from form data
	user.Username = c.GetString("username")
	user.Email = c.GetString("email")
	user.Role = c.GetString("role")

	// Validate required fields
	if user.Username == "" || user.Email == "" || user.Role == "" {
		c.Data["Error"] = "All fields are required"
		c.Redirect("/users", 303)
		return
	}

	// Save updated user to the database
	if _, err := o.Update(&user); err != nil {
		c.Data["Error"] = "Failed to update user"
		c.Redirect("/users", 303)
		return
	}

	// Redirect to users page with success message
	c.Redirect("/users", 303)
}

// DeleteUser handles DELETE /user/:id
func (c *UserController) DeleteUser() {
	rawID := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		c.Data["Error"] = "Invalid user ID."
		c.Redirect("/users", 303)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&models.User{ID: id}); err != nil {
		c.Data["Error"] = "Failed to delete user."
		c.Redirect("/users", 303)
		return
	}

	c.Redirect("/users", 303)
}

// GetAllUsers handles GET /users
func (c *UserController) GetAllUsers() {
	o := orm.NewOrm()
	var users []*models.User
	_, err := o.QueryTable(new(models.User)).All(&users)
	if err != nil {
		c.Data["Error"] = "Failed to fetch users."
		c.TplName = "users.tpl"
		return
	}

	c.Data["Users"] = users
	c.TplName = "users.tpl"
}
