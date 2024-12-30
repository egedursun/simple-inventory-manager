package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Warehouse struct {
	ID        int       `orm:"auto;pk"`
	Name      string    `orm:"size(100);unique"`
	Location  string    `orm:"size(255)"`
	Manager   *User     `orm:"rel(fk);null"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(
		new(Warehouse),
	)
}
