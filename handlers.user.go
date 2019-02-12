package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"math/rand"
	)

//var session map[string]user
var session = make(map[string]user)


func register(c *gin.Context)  {

	username,_ := c.GetPostForm("username")
	password,_ := c.GetPostForm("password")

	if u,err := registerNewUser(username, password); err==nil {
		token := generateToken()
		session[token] = *u
		render(c, RESPONSE_CODE_SUCCESS, "SUCCESS", token)
	} else {
		render(c, RESPONSE_CODE_FAILURE, err.Error(), err.Error())
	}
}

func generateToken()  string{
	return strconv.FormatInt(rand.Int63(),32)
}

func login(c *gin.Context)  {
	username,_ := c.GetPostForm("username")
	password,_ := c.GetPostForm("password")
	if err := validUser(username,password); err ==nil {
		token := generateToken()
		session[token]=user{username,""}
		render(c, RESPONSE_CODE_SUCCESS, "SUCCESS", token)
	} else {
		render(c, RESPONSE_CODE_FAILURE, err.Error(), err.Error())
	}
}

