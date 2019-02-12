package main

import (
	"github.com/gin-gonic/gin"
		)

func ensureLogin()  gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token");
		if token == "" {
			render(c, RESPONSE_CODE_FAILURE,"login first","login first")
			c.Abort()
		} else {
			if _,ok := session[token]; ok==false{
				render(c, RESPONSE_CODE_FAILURE,"invalid token","invalid token")
				c.Abort()
			}
		}
	}
}