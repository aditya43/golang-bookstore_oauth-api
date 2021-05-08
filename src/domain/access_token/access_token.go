package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/aditya43/golang-bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/aditya43/golang-bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for grant_type="password"
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for grant_type="client_credentials"
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (atR *AccessTokenRequest) Validate() *errors.RESTErr {
	switch atR.GrantType {
	case grantTypeClientCredentials:
		break

	case grantTypePassword:
		break

	default:
		return errors.BadRequestErr("Invalid grant_type parameter")
	}

	return nil
}

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

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (accessToken AccessToken) IsExpired() bool {
	return time.Unix(accessToken.Expires, 0).Before(time.Now().UTC())
}

func (accessToken *AccessToken) Generate() {
	accessToken.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", accessToken.UserId, accessToken.Expires))
}
