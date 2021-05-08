package http

import (
	"net/http"

	atDomain "github.com/aditya43/golang-bookstore_oauth-api/src/domain/access_token"

	"github.com/aditya43/golang-bookstore_oauth-api/src/services/access_token"
	"github.com/aditya43/golang-bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	UpdateExpiry(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at atDomain.AccessTokenRequest

	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.BadRequestErr("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if _, err := handler.service.Create(&at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}

func (handler *accessTokenHandler) UpdateExpiry(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Not implemented yet")
}
