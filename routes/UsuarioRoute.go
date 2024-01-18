package routes

import (
	"Base1.com/go-backend/controllers"
	"github.com/gin-gonic/gin"
)

func UsuarioRouter(router *gin.Engine) {

	routes := router.Group("api/v1/usuarios")
	routes.POST("", controllers.UsuarioCreate)
	routes.GET("", controllers.UsuarioGet)
	routes.GET("/:id", controllers.UsuarioGetById)
	routes.PUT("/:id", controllers.UsuarioUpdate)
	routes.DELETE("/:id", controllers.UsuarioDelete)

}
