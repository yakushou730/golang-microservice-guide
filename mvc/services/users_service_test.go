package services

import (
	"golang-microservice-guide/mvc/domain"
	"golang-microservice-guide/mvc/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (user *domain.User, applicationError *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 does not exist",
			StatusCode: http.StatusNotFound,
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (user *domain.User, applicationError *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}
	user, err := UsersService.GetUser(123)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.Id)
}
