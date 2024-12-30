package controllers

import (
	"encoding/json"
	"fullstack-beego-app/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type StockHistoryController struct {
	beego.Controller
}

// CreateStockHistory handles POST /stock_history
func (c *StockHistoryController) CreateStockHistory() {
	var stockHistory models.StockHistory
	if err := json.Unmarshal(
		c.Ctx.Input.RequestBody,
		&stockHistory,
	); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{
			"error": "Invalid request",
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

	if stockHistory.Stock == nil || stockHistory.ChangeType == "" || stockHistory.Quantity < 0 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{
			"error": "Invalid stock, change type, or quantity",
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
	if _, err := o.Insert(&stockHistory); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Failed to create stock history",
		}
		err := c.ServeJSON()
		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{
				"error": "Internal server error",
			}
		}
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = stockHistory
	err := c.ServeJSON()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Internal server error",
		}
	}
}

// GetStockHistory handles GET /stock_history/:id
func (c *StockHistoryController) GetStockHistory() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{
			"error": "Invalid stock history ID",
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

	o := orm.NewOrm()
	stockHistory := models.StockHistory{ID: id}
	if err := o.Read(&stockHistory); err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{
			"error": "Stock history not found",
		}
		err := c.ServeJSON()
		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{
				"error": "Internal server error",
			}
		}
		return
	}

	c.Data["json"] = stockHistory
	err = c.ServeJSON()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Internal server error",
		}
	}
}

// GetAllStockHistories handles GET /stock_history
func (c *StockHistoryController) GetAllStockHistories() {
	o := orm.NewOrm()
	var stockHistories []*models.StockHistory
	_, err := o.QueryTable(new(models.StockHistory)).RelatedSel().All(&stockHistories)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Failed to get stock histories",
		}
		err := c.ServeJSON()
		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{
				"error": "Internal server error",
			}
		}
		return
	}

	c.Data["json"] = stockHistories
	err = c.ServeJSON()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Internal server error",
		}
	}
}

// DeleteStockHistory handles DELETE /stock_history/:id
func (c *StockHistoryController) DeleteStockHistory() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{
			"error": "Invalid stock history ID",
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

	o := orm.NewOrm()
	if _, err := o.Delete(&models.StockHistory{ID: id}); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Failed to delete stock history",
		}
		err := c.ServeJSON()
		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{
				"error": "Internal server error",
			}
		}
		return
	}

	c.Ctx.Output.SetStatus(204)
	c.Data["json"] = map[string]string{
		"message": "Stock history deleted",
	}
	err = c.ServeJSON()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{
			"error": "Internal server error",
		}
	}
}
