package access_token

import (
	"strings"
	"time"

	"github.com/aditya43/golang-bookstore_oauth-api/src/utils/errors"
)

const expirationTime = 24

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RESTErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.BadRequestErr("Invalid access token")
	}

	if at.UserId <= 0 {
		return errors.BadRequestErr("Invalid user id")
	}

	if at.ClientId <= 0 {
		return errors.BadRequestErr("Invalid client id")
	}

	if at.Expires <= 0 {
		return errors.BadRequestErr("Invalid expiry time")
	}

	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (accessToken AccessToken) IsExpired() bool {
	return time.Unix(accessToken.Expires, 0).Before(time.Now().UTC())
}
