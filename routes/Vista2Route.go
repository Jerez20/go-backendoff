package routes

import (
	"Base1.com/go-backend/controllers"
	"github.com/gin-gonic/gin"
)

func TransaRouter(router *gin.Engine) {

	routes := router.Group("api/v1/trasa")

	routes.GET("", controllers.Transaget)
}

func TransacionesRouter(router *gin.Engine) {

	routes := router.Group("api/v1/t")

	routes.GET("", controllers.GetTransaccionesVentas)
}

/*func TransacionesRouter(r *gin.Engine) {

	r.GET("/transacciones-ventas", controllers.GetTransaccionesVentas)
}*/
