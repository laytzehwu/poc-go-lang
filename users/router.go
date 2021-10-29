package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	AhLayGinEngine *gin.Engine
}

func (r *Router) RouterInit() {
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
					"message": fmt.Sprintf("Welcome %s", name),
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
