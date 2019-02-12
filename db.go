package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/houkemian/xorm"
	"github.com/golang/glog"
	"github.com/go-xorm/core"
)

var engine *xorm.Engine

func initDB()  {
	var err error
	engine,err = xorm.NewEngine("mysql","dev:distribution@tcp(192.168.10.79:3306)/xl_pub_shop")
	engine.ShowSQL(true)
	engine.SetMapper(core.GonicMapper{})
	engine.SetLogLevel(core.LOG_DEBUG)
	if err!=nil {
		glog.Error("can not init db conn")
	}
}