package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func myRoutes(r *gin.RouterGroup) {
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "It Works !, Hello World !"})
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	myRoutes(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
