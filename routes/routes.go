package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) *gin.Engine {
	AppRouter(app)
	return app
}
