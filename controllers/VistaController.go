package controllers

import (
	"Base1.com/go-backend/configs"
	"Base1.com/go-backend/models"
	"github.com/gin-gonic/gin"
)

type AlmacenRequestBody struct {
	Nombre   string `json:"nombre" binding:"required"`
	Cantidad string `json:"cantidad" binding:"required"`
}

func AlmacenGet(c *gin.Context) {
	var almacens []models.Almacen
	configs.DB.Find(&almacens)
	c.JSON(200, &almacens)
}
