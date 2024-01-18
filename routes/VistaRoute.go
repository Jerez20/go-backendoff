package routes

import (
	"Base1.com/go-backend/controllers"
	"github.com/gin-gonic/gin"
)

func AlmacenRouter(router *gin.Engine) {

	routes := router.Group("api/v1/almacens")

	routes.GET("", controllers.AlmacenGet)
}
