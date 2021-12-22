package controllers

import (
	"github.com/gin-gonic/gin"
	ginserver "github.com/karnadixyz/gin-server"
	"gitlab.com/odeo/admin-iam/infra"
)

type auth struct{}

func (*auth) GetToken(c *gin.Context) {
	var request = struct {
		GrantType    string `form:"grant_type" binding:"required"`
		ClientID     string `form:"client_id" binding:"required"`
		ClientSecret string `form:"client_secret" binding:"required"`
	}{}

	if err := c.ShouldBind(&request); err != nil {
		infra.TransSrv.ErrorValidationResponse(c, err)
		return
	}
	ginserver.HandleTokenRequest(c)
}
