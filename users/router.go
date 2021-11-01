package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	AhLayGinEngine *gin.Engine
}

func defineUserRoutes(r *Router) {
	userGroup := r.AhLayGinEngine.Group("/user")
	{
		userGroup.GET("", func(c *gin.Context) {
			name := c.DefaultQuery("name", "Guest")
			fmt.Printf("User %s is accessing\n", name)
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Welcome %s", name),
			})
		})
		specUser := userGroup.Group("/:name")
		{
			specUser.GET("", func(c *gin.Context) {
				name := c.Param("name")
				fmt.Printf("User %s is accessing\n", name)
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("Welcome back %s XXX", name),
				})
			})
			specUserAction := specUser.Group("/*action")
			{
				specUserAction.GET("", func(c *gin.Context) {
					name := c.Param("name")
					action := c.Param("action")[1:]
					fmt.Printf("User %s is %sing\n", name, action)
					c.String(http.StatusOK, fmt.Sprintf("%s does %s", name, action))
				})
				specUserAction.POST("", func(c *gin.Context) {
					name := c.Param("name")
					action := c.Param("action")[1:]
					reqPath := c.FullPath()
					fmt.Printf("Receiving user action %s to do %sing via below path:\n\t%s\n", name, action, reqPath)
					c.JSON(200, gin.H{
						"message": fmt.Sprintf("Hi %s, we received your action: %s", name, action),
					})
				})
			}
		}
	}
}

func defineUserActionRoute(r *Router) {
	r.AhLayGinEngine.GET("/user-action", func(c *gin.Context) {
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
}

func defineUserActionForm(r *Router) {
	r.AhLayGinEngine.POST("/user-action-form", func(c *gin.Context) {
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
}

func (r *Router) RouterInit() {
	defineUserRoutes(r)
	defineUserActionRoute(r)
	defineUserActionForm(r)
}
