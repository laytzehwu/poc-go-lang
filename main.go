package main

import (
	"fmt"
	"net/http"

	"github.com/laytzehwu/poc-go-lang/users"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Example service is up!")

	//router := gin.Default() // This line is replaced by below few lines
	router := gin.New()
	router.Use(gin.Logger())   // Logger is included by gin.Default()
	router.Use(gin.Recovery()) // Recovery is included by gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong pong",
		})
	})
	usersRouter := users.Router{
		AhLayGinEngine: router,
	}
	usersRouter.RouterInit()

	router.GET("/user-action", func(c *gin.Context) {
		name := c.Query("name")
		action := c.Query("action")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Missing query name",
			})
			return
		}
		if action == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Missing query action",
			})
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("%s does %s", name, action))
	})
	router.POST("/user-action-form", func(c *gin.Context) {
		name := c.Copy().DefaultPostForm("name", "guest")
		action := c.PostForm("action")
		if action == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Missing action",
			})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"message": fmt.Sprintf("Hi %s your request action: %s is received", name, action),
		})
	})
	router.Run(":5000")
}
