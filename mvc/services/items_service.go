package services

import (
	"golang-microservice-guide/mvc/domain"
	"golang-microservice-guide/mvc/utils"
	"net/http"
)

type itemsService struct{}

var (
	ItemService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
