package rest

import (
	"encoding/json"
	"strings"

	"github.com/aditya43/golang-bookstore_oauth-api/src/domain/users"
	"github.com/aditya43/golang-bookstore_oauth-api/src/utils/errors"

	"github.com/go-resty/resty/v2"
)

var (
	usersRESTClient      = resty.New()
	userLoginAPIEndpoint = "http://localhost:8080/users/login"
)

type RESTUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RESTErr)
}

type userRepository struct{}

func NewRepository() RESTUsersRepository {
	return &userRepository{}
}

func (r *userRepository) LoginUser(email string, password string) (*users.User, *errors.RESTErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	var user users.User

	response, err := usersRESTClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&request).
		SetResult(&user).
		Post(userLoginAPIEndpoint)

	if response == nil || err != nil {
		if strings.Contains(err.Error(), "cannot unmarshal") {
			return nil, errors.InternalServerErr("Invalid JSON response received")
		}
		return nil, errors.InternalServerErr(err.Error())
	}

	if response.StatusCode() == 200 && user.Email != "" {
		return &user, nil
	}

	var restErr errors.RESTErr
	if err := json.Unmarshal(response.Body(), &restErr); err != nil {
		return nil, errors.InternalServerErr("Invalid error interface")
	}

	return nil, &restErr
}
