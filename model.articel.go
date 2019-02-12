package main

import (
		"errors"
)

type article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticles() [] article {
	return articleList
}

func getArticleById(id int) (*article,error) {
	for _,ar := range articleList  {
		if ar.ID==id {
			return &ar,nil
		}
	}
	return nil,errors.New("article not found")
}

func newArticle(title, content string) [] article {
	a := article{ID:len(articleList)+1,Title:title, Content:content}
	articleList = append(articleList,a)
	return articleList
}