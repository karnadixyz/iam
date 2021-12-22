package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginserver "github.com/karnadixyz/gin-server"
	"gitlab.com/odeo/admin-iam/controllers"
)

func AppRouter(r *gin.Engine) {
	auth := r.Group("/oauth2")
	{
		auth.GET("/token", controllers.AppController.Auth.GetToken)
	}

	//guard auth
	api := r.Group("/api")
	{
		api.Use(ginserver.HandleTokenVerify())
		api.GET("/test", func(c *gin.Context) {
			ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
			if exists {
				c.JSON(http.StatusOK, ti)
				return
			}
			c.String(http.StatusOK, "not found")
		})
	}
}
