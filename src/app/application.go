package app

import (
	"github.com/ericha1981/bookstore_oauth-api/src/domain/access_token"
	"github.com/ericha1981/bookstore_oauth-api/src/http"
	"github.com/ericha1981/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication()  {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository())) // Presentation/Web layer (http)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
