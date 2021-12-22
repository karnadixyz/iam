package services

import (
	"fmt"
	"os"

	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	ginserver "github.com/karnadixyz/gin-server"
	oauth2gorm "github.com/karnadixyz/go-oauth2-gorm"
)

func InitOauth2Service() {
	manager := manage.NewDefaultManager()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	configClient := oauth2gorm.NewConfig(dsn, 1, "oauth2_client")
	configToken := oauth2gorm.NewConfig(dsn, 1, "oauth2_token")
	clientStore := oauth2gorm.NewClientStore(configClient)
	store := oauth2gorm.NewTokenStore(configToken, 600)
	defer store.Close()
	manager.MapClientStorage(clientStore)
	manager.MapTokenStorage(store)
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)
}
