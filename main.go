package main

import (
	"blog/utils"
	_ "blog/models"
	_ "blog/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	database := beego.AppConfig.String("database")

	db_str := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&loc=Local", username, password, host, port, database)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", db_str)

	err := orm.RunSyncdb("default", false, true) //创建表
	if err != nil {
		panic(err)
	}
}

func main() {
	beego.InsertFilter("cms/index/*", beego.BeforeRouter, utils.CmsLoginFilter)
	beego.Run()
}
