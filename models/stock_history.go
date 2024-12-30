package models

import "github.com/beego/beego/v2/client/orm"

type StockHistory struct {
	ID         int    `orm:"auto;pk"`
	Stock      *Stock `orm:"rel(fk);on_delete(cascade)"`
	ChangeType string `orm:"size(50)"`
	Quantity   int    `orm:""`
	ChangedBy  *User  `orm:"rel(fk);"`
	CreatedAt  string `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(
		new(StockHistory),
	)
}
