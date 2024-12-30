package controllers

import (
	"fullstack-beego-app/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type StockController struct {
	beego.Controller
}

// AddStock handles POST /stock
func (c *StockController) AddStock() {
	warehouseID, err := c.GetInt("warehouse_id")
	if err != nil || warehouseID <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{
			"error": "Invalid warehouse ID",
		}
		_ = c.ServeJSON()
		return
	}

	productID, err := c.GetInt("product_id")
	if err != nil || productID <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{
			"error": "Invalid product ID",
		}
		_ = c.ServeJSON()
		return
	}

	quantity, err := c.GetInt("quantity")
	if err != nil || quantity < 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{
			"error": "Invalid quantity",
		}
		_ = c.ServeJSON()
		return
	}

	threshold, err := c.GetInt("threshold")
	if err != nil || threshold < 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{
			"error": "Invalid threshold",
		}
		_ = c.ServeJSON()
		return
	}

	stock := models.Stock{
		Warehouse: &models.Warehouse{ID: warehouseID},
		Product:   &models.Product{ID: productID},
		Quantity:  quantity,
		Threshold: threshold,
	}

	o := orm.NewOrm()
	if _, err := o.Insert(&stock); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Failed to add stock",
		}
		_ = c.ServeJSON()
		return
	}

	c.Redirect("/stocks", 302)
}

// UpdateStock handles PUT /stock/:id
func (c *StockController) UpdateStock() {
	id, err := c.GetInt(":id")
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["Error"] = "Invalid stock ID"
		c.Redirect("/stocks", 303)
		return
	}

	o := orm.NewOrm()
	stock := models.Stock{ID: id}
	if err = o.Read(&stock); err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["Error"] = "Stock not found"
		c.Redirect("/stocks", 302)
		return
	}

	// Parse form data
	warehouseID, err := c.GetInt("warehouse_id")
	if err == nil && warehouseID > 0 {
		stock.Warehouse = &models.Warehouse{ID: warehouseID}
	}

	productID, err := c.GetInt("product_id")
	if err == nil && productID > 0 {
		stock.Product = &models.Product{ID: productID}
	}

	quantity, err := c.GetInt("quantity")
	if err == nil && quantity >= 0 {
		stock.Quantity = quantity
	}

	threshold, err := c.GetInt("threshold")
	if err == nil && threshold >= 0 {
		stock.Threshold = threshold
	}

	// Validate required fields
	if stock.Warehouse == nil || stock.Product == nil || stock.Quantity < 0 {
		c.Data["Error"] = "Warehouse, Product, and Quantity are required"
		c.Redirect("/stocks", 303)
		return
	}

	// Perform the update
	if _, err := o.Update(&stock); err != nil {
		c.Data["Error"] = "Failed to update stock"
		c.Redirect("/stocks", 303)
		return
	}

	// Redirect to the stocks page
	c.Redirect("/stocks", 303)
}

// DeleteStock handles DELETE /stock/:id
func (c *StockController) DeleteStock() {
	id, err := c.GetInt(":id")
	if err != nil || id <= 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{
			"error": "Invalid stock ID",
		}
		_ = c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&models.Stock{ID: id}); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Failed to delete stock",
		}
		_ = c.ServeJSON()
		return
	}

	// Redirect explicitly as a GET request
	c.Redirect("/stocks", 303)
}

// GetAllStocks handles GET /stocks
func (c *StockController) GetAllStocks() {
	o := orm.NewOrm()
	var stocks []*models.Stock
	if _, err := o.QueryTable("stock").RelatedSel().All(&stocks); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Failed to retrieve stocks",
		}
		_ = c.ServeJSON()
		return
	}

	var warehouses []*models.Warehouse
	if _, err := o.QueryTable("warehouse").All(&warehouses); err != nil {
		warehouses = []*models.Warehouse{}
	}

	var products []*models.Product
	if _, err := o.QueryTable("product").All(&products); err != nil {
		products = []*models.Product{}
	}

	c.Data["Stocks"] = stocks
	c.Data["Warehouses"] = warehouses
	c.Data["Products"] = products
	c.TplName = "stocks.tpl"
}
