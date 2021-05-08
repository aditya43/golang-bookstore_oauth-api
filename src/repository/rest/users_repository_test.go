package rest

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/aditya43/golang-bookstore_oauth-api/src/domain/users"
	restErr "github.com/aditya43/golang-bookstore_oauth-api/src/utils/errors"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	httpmock.ActivateNonDefault(usersRESTClient.GetClient())
	defer httpmock.DeactivateAndReset()

	type invalidErrStruct struct {
		Message string
		Status  string // This must be int
		Error   string
	}

	responder, _ := httpmock.NewJsonResponder(400, &invalidErrStruct{
		Message: "Test Message",
		Status:  "400", // This must be int
		Error:   "bad_request",
	})
	httpmock.RegisterResponder("POST", "http://localhost:8080/users/login", responder)

	repository := userRepository{}
	user, err := repository.LoginUser("aditya@gmail.com", "123456789")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error interface", err.Message)
}
func TestLoginUserTimeoutFromApi(t *testing.T) {
	httpmock.ActivateNonDefault(usersRESTClient.GetClient())
	defer httpmock.DeactivateAndReset()

	responder := httpmock.NewErrorResponder(errors.New("Request timeout"))
	httpmock.RegisterResponder("POST", "http://localhost:8080/users/login", responder)

	repository := userRepository{}
	user, err := repository.LoginUser("aditya@gmail.com", "123456789")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "internal_server_error", err.Error)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	httpmock.ActivateNonDefault(usersRESTClient.GetClient())
	defer httpmock.DeactivateAndReset()

	responder, _ := httpmock.NewJsonResponder(400, restErr.NotFoundErr("Invalid login credentials"))
	httpmock.RegisterResponder("POST", "http://localhost:8080/users/login", responder)

	repository := userRepository{}
	user, err := repository.LoginUser("aditya@gmail.com", "123456789")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	httpmock.ActivateNonDefault(usersRESTClient.GetClient())
	defer httpmock.DeactivateAndReset()

	type invalidUserStruct struct {
		Id        string // This must be int
		Email     string
		FirstName string
		LastName  string
	}

	responder, _ := httpmock.NewJsonResponder(http.StatusOK, &invalidUserStruct{
		Id:        "123",
		Email:     "aditya@hajare.com",
		FirstName: "Aditya",
		LastName:  "Hajare",
	})
	httpmock.RegisterResponder("POST", "http://localhost:8080/users/login", responder)

	repository := userRepository{}
	user, err := repository.LoginUser("aditya@gmail.com", "123456789")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid JSON response received", err.Message)
}

func TestLoginUserSuccessful(t *testing.T) {
	httpmock.ActivateNonDefault(usersRESTClient.GetClient())
	defer httpmock.DeactivateAndReset()

	responder, _ := httpmock.NewJsonResponder(http.StatusOK, &users.User{
		Id:        123,
		Email:     "aditya@hajare.com",
		FirstName: "Aditya",
		LastName:  "Hajare",
	})
	httpmock.RegisterResponder("POST", "http://localhost:8080/users/login", responder)

	repository := userRepository{}
	user, err := repository.LoginUser("aditya@gmail.com", "123456789")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "aditya@hajare.com", user.Email)
	assert.EqualValues(t, "Aditya", user.FirstName)
	assert.EqualValues(t, "Hajare", user.LastName)
}
