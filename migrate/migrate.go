package main

import (
	"Base1.com/go-backend/configs"
	"Base1.com/go-backend/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.AccionesRH{})
	configs.DB.AutoMigrate(&models.Usuario{})
	configs.DB.AutoMigrate(&models.Almacen{})
	configs.DB.AutoMigrate(&models.Transa01{})
}
