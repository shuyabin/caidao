package main

import (
	"caidao/models"
	_ "caidao/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/caidao.db")
	orm.RegisterModel(new(models.Caidao))
	// orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
	
}
