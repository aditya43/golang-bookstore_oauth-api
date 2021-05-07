package db

import (
	"github.com/aditya43/golang-bookstore_oauth-api/src/domain/access_token"
	"github.com/aditya43/golang-bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RESTErr)
}

type dbRepository struct {
}

func (db *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RESTErr) {
	return nil, errors.InternalServerErr("Database connection not implemented yet..")
}