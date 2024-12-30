package test_db

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

var dbInitOnce sync.Once

func resetDatabase() {
	o := orm.NewOrm()

	// Drop all tables
	for _, table := range []string{"user", "warehouse", "stock", "product"} {
		_, err := o.Raw("DROP TABLE IF EXISTS " + table).Exec()
		if err != nil {
			log.Fatalf("Failed to drop table %s: %v", table, err)
		}
	}

	// Recreate tables
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatalf("Failed to recreate tables: %v", err)
	}
}

func InitTestDB() {
	dbInitOnce.Do(func() {
		orm.Debug = true
		err := orm.RegisterDriver("sqlite3", orm.DRSqlite)
		if err != nil {
			return
		}
		err = orm.RegisterDataBase("default", "sqlite3", "file::memory:?cache=shared")
		if err != nil {
			return
		}
		err = orm.RunSyncdb("default", false, true)
		if err != nil {
			panic(err)
		}
	})
	resetDatabase()
}
