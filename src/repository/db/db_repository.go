package db

import (
	"github.com/aditya43/golang-bookstore_oauth-api/src/client/cassandra"
	"github.com/aditya43/golang-bookstore_oauth-api/src/domain/access_token"
	"github.com/aditya43/golang-bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpiry      = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RESTErr)
	Create(access_token.AccessToken) *errors.RESTErr
	UpdateExpiry(access_token.AccessToken) *errors.RESTErr
}

type dbRepository struct {
}

func (db *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RESTErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.InternalServerErr(err.Error())
	}
	defer session.Close()

	var res access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&res.AccessToken,
		&res.UserId,
		&res.ClientId,
		&res.Expires,
	); err != nil {
		return nil, errors.InternalServerErr(err.Error())
	}

	return &res, nil
}

func (db *dbRepository) Create(at access_token.AccessToken) *errors.RESTErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return errors.InternalServerErr(err.Error())
	}

	return nil
}

func (db *dbRepository) UpdateExpiry(at access_token.AccessToken) *errors.RESTErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryUpdateExpiry,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.InternalServerErr(err.Error())
	}

	return nil
}
