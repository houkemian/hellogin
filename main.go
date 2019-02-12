package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)

var router *gin.Engine
var httpClient *http.Client = new(http.Client)

func main() {
	router = gin.Default()
	initializeRoutes()
	initDB()
	router.Run()
}

func render(c *gin.Context, code int, message string,  data interface{}) {

	accept := c.Request.Header.Get("Accept")

	r := gin.H{
		"code": code,
		"message": message,
		"data": data,
	}

	switch accept {

	case "application/json":
		c.JSON(http.StatusOK, r)
	case "application/xml":
		c.XML(http.StatusOK, r)
	default:
		c.JSON(http.StatusOK, r)
	}

}
