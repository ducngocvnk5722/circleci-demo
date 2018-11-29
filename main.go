package main

import (
	"net/http"

	"github.com/ducngocvnk57/circleci-demo/app"
	htmlrender "github.com/ducngocvnk57/circleci-demo/html-render"
	"github.com/ducngocvnk57/circleci-demo/services"
	"github.com/gin-gonic/gin"
)

var DB = app.Db()

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		test := services.GetAllJobType()
		c.HTML(http.StatusOK, "home/index", gin.H{
			"title": "Main website 201803121547",
			"test":  test,
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Static("/assets", "./assets")
	htmlRender := htmlrender.New()
	r.HTMLRender = htmlRender.Create()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
