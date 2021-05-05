package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time constant is not set to 24 hrs")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.IsExpired(), "Brand new access token should not be expired")

	assert.EqualValues(t, "", at.AccessToken, "Brand new access token should not have defined access token id")

	assert.True(t, at.UserId == 0, "Brand new access token should not have associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.IsExpired(), "Empty access token should always be identified as expired")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Access token expiring after 3 hrs should not be identified as expired")

}
