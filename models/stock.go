package models

import "github.com/beego/beego/v2/client/orm"

type Stock struct {
	ID        int        `orm:"auto;pk"`
	Warehouse *Warehouse `orm:"rel(fk);on_delete(cascade)"`
	Product   *Product   `orm:"rel(fk);on_delete(cascade)"`
	Quantity  int        `orm:"default(0)"`
	Threshold int        `orm:"default(0)"`
	UpdatedAt string     `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(
		new(Stock),
	)
}
