package main

import (
	"fmt"

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
	router.Run(":5000")
}
