package access_token

import (
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("Expiration time constant is not set to 24 hrs")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	if at.IsExpired() {
		t.Error("Brand new access token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("Brand new access token should not have defined access token id")
	}

	if at.UserId != 0 {
		t.Error("Brand new access token should not have associated user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	if !at.IsExpired() {
		t.Error("Empty access token should always be identified as expired")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("Access token expiring after 3 hrs should not be identified as expired")
	}
}
