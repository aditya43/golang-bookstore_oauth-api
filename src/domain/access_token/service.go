package access_token

import (
	"strings"

	"github.com/aditya43/golang-bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RESTErr)
	Create(*AccessToken) *errors.RESTErr
	UpdateExpiry(*AccessToken) *errors.RESTErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RESTErr)
	Create(*AccessToken) *errors.RESTErr
	UpdateExpiry(*AccessToken) *errors.RESTErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (service *service) GetById(accessTokenId string) (*AccessToken, *errors.RESTErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.BadRequestErr("Invalid access token id")
	}

	accessToken, err := service.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (service *service) Create(at *AccessToken) *errors.RESTErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return service.repository.Create(at)
}

func (service *service) UpdateExpiry(at *AccessToken) *errors.RESTErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return service.repository.UpdateExpiry(at)
}
