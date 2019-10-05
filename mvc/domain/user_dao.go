package domain

import (
	"fmt"
	"golang-microservice-guide/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Shou", LastName: "Tseng", Email: "yakushou730@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exist", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
