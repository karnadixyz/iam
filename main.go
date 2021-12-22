package main

import (
	"github.com/joho/godotenv"
	"gitlab.com/odeo/admin-iam/cmd"
	"gitlab.com/odeo/admin-iam/infra"
)

func main() {
	// Load env
	godotenv.Load()

	if err := infra.ServiceInit(); err != nil {
		panic(err)
	}
	cmd.Execute()
}
