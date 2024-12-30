package controllers

import (
	"fullstack-beego-app/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"strconv"
)

type ProductController struct {
	beego.Controller
}

// CreateProduct handles POST /product
func (c *ProductController) CreateProduct() {
	product := models.Product{
		Name: c.GetString("name"),
		SKU:  c.GetString("sku"),
	}

	// Validate required fields
	if product.Name == "" || product.SKU == "" {
		c.Data["Error"] = "Name and SKU are required"
		c.TplName = "products.tpl"
		return
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&product); err != nil {
		c.Data["Error"] = "Failed to create product"
		c.TplName = "products.tpl"
		return
	}

	// Redirect to the products page
	c.Redirect("/products", 302)
}

// GetProduct handles GET /product/:id
func (c *ProductController) GetProduct() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{
			"error": "Invalid product ID",
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

	o := orm.NewOrm()
	product := models.Product{ID: id}
	if err = o.Read(&product); err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{
			"error": "Product not found",
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

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = product
	err = c.ServeJSON()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Internal server error",
		}
		return
	}

	return
}

// UpdateProduct handles PUT /product/:id
func (c *ProductController) UpdateProduct() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid product ID"}
		_ = c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	product := models.Product{ID: id}
	if err = o.Read(&product); err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": "Product not found"}
		_ = c.ServeJSON()
		return
	}

	// Update fields from form data
	product.Name = c.GetString("name")
	product.SKU = c.GetString("sku")

	if product.Name == "" || product.SKU == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Name and SKU are required"}
		_ = c.ServeJSON()
		return
	}

	if _, err := o.Update(&product); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to update product"}
		_ = c.ServeJSON()
		return
	}

	c.Redirect("/products", 302)
}

// DeleteProduct handles DELETE /product/:id
func (c *ProductController) DeleteProduct() {
	rawID := c.Ctx.Input.Param(":id")
	log.Printf("Incoming request: Method=%s, _method=%s, ID=%s", c.Ctx.Request.Method, c.GetString("_method"), rawID)

	id, err := strconv.Atoi(rawID)
	if err != nil {
		log.Printf("Error converting ID: %v", err)
		c.Data["Error"] = "Invalid product ID"
		c.Redirect("/products", 303)
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&models.Product{ID: id}); err != nil {
		log.Printf("Error deleting product with ID %d: %v", id, err)
		c.Data["Error"] = "Failed to delete product"
		c.Redirect("/products", 303)
		return
	}

	log.Printf("Product with ID %d deleted successfully", id)
	c.Redirect("/products", 303)
}

// GetAllProducts handles GET /products
func (c *ProductController) GetAllProducts() {
	o := orm.NewOrm()
	var products []models.Product
	_, err := o.QueryTable(new(models.Product)).All(&products)
	if err != nil {
		c.Data["Error"] = "Failed to retrieve products"
		c.TplName = "products.tpl"
		return
	}
	c.Data["Products"] = products
	c.TplName = "products.tpl"
}
