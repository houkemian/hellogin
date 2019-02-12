package main

import (
	"github.com/golang/glog"
	"strconv"
)

type LabelSchedule struct {
	Id int `xorm:"pk autoincr" json:"id"`
	LabelNames string `xorm:"varchar(255) 'label_names'" json:"label"`
	SkuCodes string `xorm:"varchar(255) 'sku_codes'" json:"skucode"`
}

func get(id int) (schedule *LabelSchedule)  {
	glog.Info(id)
	schedule = new(LabelSchedule)
	has, err := engine.ID(id).Get(schedule)
	glog.Info(has)
	glog.Info(err)
	return
}

func add(label, skuCode string) (schedule *LabelSchedule) {


	session := engine.NewSession()

	defer session.Close()

	session.Begin()

	schedule = new(LabelSchedule)
	schedule.LabelNames = label
	schedule.SkuCodes = skuCode
	count,err := session.Insert(schedule)

	if _,errr := strconv.Atoi(label); errr!=nil {
		glog.Error("error, rollback")
		session.Rollback()
		return
	}

	sc := new(LabelSchedule)
	sc.LabelNames = label
	sc.SkuCodes = skuCode
	session.Insert(sc)

	session.Commit()

	glog.Info(schedule.Id)
	glog.Info(count)
	glog.Info(err)

	return
}

func getAll() ([]LabelSchedule) {

	schedules := make([]LabelSchedule,0)
	engine.Where("id>?",0).And("id<?",10000).Asc("id").Find(&schedules)

	res,err := httpClient.Get("http://localhost:8080/label/1000")
	if err!=nil {
		glog.Error(err.Error())
	}
	p := make([]byte,res.ContentLength)
	res.Body.Read(p)
	resStr := string(p[:])
	glog.Info(resStr)
	return schedules
}