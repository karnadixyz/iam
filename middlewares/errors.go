package middlewares

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type ErrorInvalidSignature struct{}

func (*ErrorInvalidSignature) Error() string {
	return fmt.Sprintf("Invalid Signature")
}

type ErrorUnauthorizedAccess struct{}

func (*ErrorUnauthorizedAccess) Error() string {
	return fmt.Sprintf("Unauthorized Access")
}

type ErrorInvalidToken struct{}

func (*ErrorInvalidToken) Error() string {
	return fmt.Sprintf("Invalid Authorization Token")
}

type ErrorTokenData struct{}

func (*ErrorTokenData) Error() string {
	return fmt.Sprintf("Invalid Authorization Data")
}

type InvalidApiRequestError struct{}

func (*InvalidApiRequestError) Error() string {
	return fmt.Sprintf("Invalid Api Request")
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				log.Errorf("%v", err)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
