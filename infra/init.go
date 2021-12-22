package infra

import "os"

func ServiceInit() error {
	if err := IamDbSrv.Init(&DbConfig{
		Host:  os.Getenv("DB_HOST"),
		Name:  os.Getenv("DB_NAME"),
		User:  os.Getenv("DB_USER"),
		Pass:  os.Getenv("DB_PASSWORD"),
		Port:  os.Getenv("DB_PORT"),
		Debug: true,
	}); err != nil {
		return err
	}
	AppSrv.Init()
	return nil
}
