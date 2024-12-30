package controllers

import (
	"fullstack-beego-app/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/opentracing/opentracing-go/log"
)

type WarehouseController struct {
	beego.Controller
}

// CreateWarehouse handles POST /warehouse
func (c *WarehouseController) CreateWarehouse() {
	warehouse := models.Warehouse{
		Name:     c.GetString("name"),
		Location: c.GetString("location"),
	}

	// Validate required fields
	if warehouse.Name == "" || warehouse.Location == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Name and location are required"}
		_ = c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	_, err := o.Insert(&warehouse)
	if err != nil {
		// Log the detailed error
		log.Error(err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to create warehouse"}
		_ = c.ServeJSON()
		return
	}

	c.Redirect("/warehouses", 302)
}

// UpdateWarehouse handles PUT /warehouse/:id
func (c *WarehouseController) UpdateWarehouse() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid warehouse ID"}
		_ = c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	warehouse := models.Warehouse{ID: id}
	if err = o.Read(&warehouse); err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": "Warehouse not found"}
		_ = c.ServeJSON()
		return
	}

	warehouse.Name = c.GetString("name")
	warehouse.Location = c.GetString("location")

	if warehouse.Name == "" || warehouse.Location == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Name and location are required"}
		_ = c.ServeJSON()
		return
	}

	if _, err := o.Update(&warehouse); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to update warehouse"}
		_ = c.ServeJSON()
		return
	}

	c.Redirect("/warehouses", 302)
}

// DeleteWarehouse handles DELETE /warehouse/:id
func (c *WarehouseController) DeleteWarehouse() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid warehouse ID"}
		_ = c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&models.Warehouse{ID: id}); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to delete warehouse"}
		_ = c.ServeJSON()
		return
	}

	c.Redirect("/warehouses", 302)
}

// GetAllWarehouses handles GET /warehouses
func (c *WarehouseController) GetAllWarehouses() {
	o := orm.NewOrm()
	var warehouses []models.Warehouse
	_, err := o.QueryTable(new(models.Warehouse)).All(&warehouses)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to retrieve warehouses"}
		_ = c.ServeJSON()
		return
	}

	c.Data["Warehouses"] = warehouses
	c.TplName = "warehouses.tpl"
	_ = c.Render()
}
