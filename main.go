package main

import (
	"net/http"

	"github.com/ducngocvnk57/circleci-demo/app"
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
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website 201803121547",
			"test":  test,
		})
	})
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		c.JSON(200, gin.H{"user": user, "status": "no value"})
	})
	r.GET("/user/:name/detail", func(c *gin.Context) {
		user := c.Params.ByName("name")
		c.JSON(200, gin.H{"user": user, "status": "no value"})
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("views/*")
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
