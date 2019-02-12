package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	)

func loadAllArticle(c *gin.Context)  {
	render(c,RESPONSE_CODE_SUCCESS, "SUCCESS", getAllArticles())

}

func loadDetail(c *gin.Context) {

	id,err := strconv.Atoi(c.Param("article_id"))

	if (err!= nil) {
		fmt.Errorf("%v",err)
	}

	if article,err := getArticleById(id); err==nil{
		render(c,RESPONSE_CODE_SUCCESS, "SUCCESS", article)
	} else {
		render(c,RESPONSE_CODE_FAILURE, err.Error(), err.Error())
	}

}

func post(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	newArticle(title,content)
	render(c, RESPONSE_CODE_SUCCESS, "SUCCESS", articleList)
}