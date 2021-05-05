package access_token

import (
	"github.com/aditya43/golang-bookstore_oauth-api/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RESTErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RESTErr)
}
