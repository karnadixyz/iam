package cmd

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"gitlab.com/odeo/admin-iam/middlewares"
	"gitlab.com/odeo/admin-iam/routes"
	"gitlab.com/odeo/admin-iam/services"
)

func init() {
	rootCmd.AddCommand(gatewayCmd)
}

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Public gateway (REST)",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func New(mdls ...gin.HandlerFunc) *gin.Engine {
	api := gin.New()
	api.Use(mdls...)
	return api
}

func NewDefault(mdls ...gin.HandlerFunc) *gin.Engine {
	defMdls := []gin.HandlerFunc{
		middlewares.ErrorMiddleware(),
	}
	api := New(defMdls...)
	api.Use(mdls...)
	return api
}

type RegisterHandler func(engine *gin.Engine)

var RegisterHandles []RegisterHandler

func SetRoutes(router *gin.Engine) {
	for _, h := range RegisterHandles {
		h(router)
	}
}

func Run() {
	gin.SetMode(gin.DebugMode)

	router := NewDefault(cors.New(cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			for _, host := range []string{"*"} {
				if origin == host || host == "*" {
					return true
				}
			}
			return false
		},
		MaxAge: 12 * time.Hour,
	}))
	services.InitOauth2Service()
	SetRoutes(routes.Setup(router))

	if err := router.Run(); err != nil {
		log.Fatalf("run api error: %v", err)
		panic(err)
	}
}
