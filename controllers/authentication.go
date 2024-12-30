package controllers

import (
	"fullstack-beego-app/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// AuthController operations for Authentication
type AuthController struct {
	beego.Controller
}

func (c *AuthController) HomePage() {
	c.Data["Error"] = ""
	c.TplName = "home.tpl"
}

func (c *AuthController) LoginPage() {
	c.Data["Error"] = ""
	c.TplName = "login.tpl"
}

func (c *AuthController) RegisterPage() {
	c.Data["Error"] = ""
	c.TplName = "register.tpl"
}

func (c *AuthController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")

	if username == "" || password == "" {
		c.Data["Error"] = "Username and password are required"
		c.TplName = "login.tpl"
		return
	}

	o := orm.NewOrm()
	user := models.User{Username: username}

	if err := o.Read(&user, "Username"); err != nil || user.Password != password {
		c.Data["Error"] = "Invalid username or password"
		c.TplName = "login.tpl"
		return
	}

	if err := c.SetSession("uid", user.ID); err != nil {
		c.Data["Error"] = "Failed to start session"
		c.TplName = "login.tpl"
		return
	}

	c.Redirect("/home", 303)
}

func (c *AuthController) Logout() {
	err := c.DestroySession()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Internal server error",
		}
		err := c.ServeJSON()
		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{
				"error": "Internal server error",
			}
			return
		}
		c.Data["json"] = map[string]string{
			"message": "Logout successful",
		}
	}

	c.Redirect("/", 303)
}

func (c *AuthController) ValidateSession() {
	if c.GetSession("uid") == nil {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = map[string]string{
			"error": "Unauthorized",
		}
		err := c.ServeJSON()
		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{
				"error": "Internal server error",
			}
			return
		}
		return
	}

	c.Data["json"] = map[string]string{
		"message": "Session valid",
	}
	err := c.ServeJSON()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Internal server error",
		}
		return
	}
}

func (c *AuthController) Register() {
	// Retrieve form values
	username := c.GetString("username")
	email := c.GetString("email")
	password := c.GetString("password")
	role := c.GetString("role")

	// Validate form inputs
	if username == "" || email == "" || password == "" {
		c.Data["Error"] = "All fields are required"
		c.TplName = "register.tpl"
		return
	}

	// Check for existing user
	o := orm.NewOrm()
	existingUser := models.User{Username: username}
	if err := o.Read(&existingUser, "Username"); err == nil {
		c.Data["Error"] = "User already exists"
		c.TplName = "register.tpl"
		return
	}

	// Create new user
	newUser := models.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}
	if _, err := o.Insert(&newUser); err != nil {
		c.Data["Error"] = "Failed to create user"
		c.TplName = "register.tpl"
		return
	}

	// Redirect to login page
	c.Redirect("/login", 303)
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hashed), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}
