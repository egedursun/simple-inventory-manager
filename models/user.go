package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	ID        int       `orm:"auto;pk"`
	Username  string    `orm:"size(100);unique"`
	Password  string    `orm:"size(255)"`
	Email     string    `orm:"size(255);unique"`
	Role      string    `orm:"size(50)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(
		new(User),
	)
}
