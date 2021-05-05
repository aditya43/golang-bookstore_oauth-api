package app

import (
	"github.com/aditya43/golang-bookstore_oauth-api/http"
	"github.com/aditya43/golang-bookstore_oauth-api/repository/db"
	"github.com/aditya43/golang-bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080 ")
}
