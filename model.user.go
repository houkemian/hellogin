package main

import (
	"strings"
	"errors"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
} 


var userList = []user{
	user{"houkemian","123"},
	user{"sunrui","321"},
}


func registerNewUser(username, password string) (*user,error) {
	if strings.TrimSpace(password)=="" {
		return nil, errors.New("password can not be nil")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("this username is exist")
	}
	u := user{username,password}
	userList = append(userList, u)
	return &u,nil
}

func isUsernameAvailable(username string) bool {
	for _,u := range userList {
		if username==u.Username {
			return false
		}
	}
	return true;
}

func validUser(username, password string) (error)  {

	for _,u := range userList {
		if u.Username==username && u.Password == password {
			return nil
		}
	}

	return errors.New("username or password can not match anything")
}