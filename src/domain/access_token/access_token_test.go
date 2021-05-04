package access_token

import "testing"

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	if at.IsExpired() {
		t.Error("Brand new access token should not be expired")
	}
}
