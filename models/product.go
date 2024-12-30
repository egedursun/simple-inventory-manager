package models

import "github.com/beego/beego/v2/client/orm"

type Product struct {
	ID          int     `orm:"auto;pk"`
	Name        string  `orm:"size(100);unique;"`
	Description string  `orm:"size(500);"`
	SKU         string  `orm:"size(100);unique;"`
	Price       float64 `orm:"digits(10);decimals(2);"`
	CreatedAt   string  `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   string  `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(
		new(Product),
	)
}
