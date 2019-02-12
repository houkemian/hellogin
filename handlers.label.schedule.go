package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	)

func loadLabelScheduleDetail(c *gin.Context) {

	id,_ := strconv.Atoi(c.Param("id"));
	schedule := get(id)
	render(c, RESPONSE_CODE_SUCCESS, "SUCCESS", schedule)

}

func postLabelSchedule(c *gin.Context)  {

	label := c.PostForm("label")
	skucode := c.PostForm("skucode")

	add(label, skucode)

	render(c, RESPONSE_CODE_SUCCESS, "SUCCESS", "SUCCESS")
}

func loadAll(c *gin.Context)  {

	schedules := getAll()

	render(c, RESPONSE_CODE_SUCCESS, "SUCCESS", schedules)

}