package access_token

import (
	"strings"

	"github.com/aditya43/golang-bookstore_oauth-api/src/domain/access_token"
	"github.com/aditya43/golang-bookstore_oauth-api/src/repository/db"
	"github.com/aditya43/golang-bookstore_oauth-api/src/repository/rest"
	"github.com/aditya43/golang-bookstore_oauth-api/src/utils/errors"
	// "github.com/federicoleon/bookstore_utils-go/rest_errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RESTErr)
	Create(*access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RESTErr)
	UpdateExpirationTime(*access_token.AccessToken) *errors.RESTErr
}

type service struct {
	restUsersRepo rest.RESTUsersRepository
	dbRepo        db.DbRepository
}

func NewService(usersRepo rest.RESTUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RESTErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.BadRequestErr("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request *access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RESTErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(&at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at *access_token.AccessToken) *errors.RESTErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpiry(at)
}
