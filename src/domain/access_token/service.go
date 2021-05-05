package access_token

import (
	"strings"

	"github.com/aditya43/golang-bookstore_oauth-api/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RESTErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RESTErr)
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
