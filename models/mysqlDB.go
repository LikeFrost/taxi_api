package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDB() {
	orm.RegisterModel(new(User), new(Car), new(Question), new(Traffic), new(Money))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	orm.RegisterDataBase("taxi", "mysql", beego.AppConfig.String("sqlconn"))
}
