package main

import (
	"net/http"

	"Base1.com/go-backend/configs"
	"Base1.com/go-backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}

func main() {

	r := gin.Default()
	r.Use(CORS())
	routes.UsuarioRouter(r)
	routes.AlmacenRouter(r)
	routes.TransaRouter(r)
	routes.TransacionesRouter(r)

	authorized := r.Group("/")

	authorized.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from the protected API endpoint.",
		})
	})

	r.Run()
}
