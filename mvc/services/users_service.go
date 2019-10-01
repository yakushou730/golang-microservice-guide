package services

import (
	"golang-microservice-guide/mvc/domain"
	"golang-microservice-guide/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
