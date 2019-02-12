package main

import "testing"

func TestGetAllArticels(t *testing.T)  {
	temp := getAllArticles()

	if (len(temp)!= len(articleList)) {
		t.Fail()
	}

	for i,v := range temp {
		if v.ID!=articleList[i].ID || v.Title != articleList[i].Title || v.Content!= articleList[i].Content{
			t.Fail()
			break
		}
	}
}
