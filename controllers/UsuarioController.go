package controllers

import (
	"Base1.com/go-backend/configs"
	"Base1.com/go-backend/models"
	"Base1.com/go-backend/util"
	"github.com/gin-gonic/gin"
)

type UsuarioRequestBody struct {
	Usuario string `json:"usuario" binding:"required"`
	Clave   string `json:"clave" binding:"required"`
	Dire    string `json:"dire" binding:"required"`
	Tele    string `json:"tele" binding:"required"`
	Correo  string `json:"correo" binding:"required"`
}

func UsuarioCreate(c *gin.Context) {
	var body UsuarioRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if len(body.Usuario) == 0 || len(body.Clave) == 0 || len(body.Dire) == 0 || len(body.Tele) == 0 || len(body.Correo) == 0 {
		c.JSON(400, gin.H{"error": "Missing required fields"})
		return
	}

	hashedPassword, err := util.CreatePassword(body.Clave)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	usuario := &models.Usuario{
		Usuario: body.Usuario,
		Clave:   hashedPassword,
		Dire:    body.Dire,
		Tele:    body.Tele,
		Correo:  body.Correo,
	}

	result := configs.DB.Create(usuario)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to insert: " + result.Error.Error()})
		return
	}

	// Enviar el correo electrónico si es necesario

	c.JSON(200, gin.H{"message": "Usuario creado con éxito"})
}

func UsuarioGet(c *gin.Context) {
	var usuarios []models.Usuario
	configs.DB.Find(&usuarios)
	c.JSON(200, &usuarios)

}

func UsuarioGetById(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario
	configs.DB.First(&usuario, id)
	c.JSON(200, &usuario)

}

func UsuarioUpdate(c *gin.Context) {

	id := c.Param("id")
	var usuario models.Usuario
	configs.DB.First(&usuario, id)

	body := UsuarioRequestBody{}
	c.BindJSON(&body)
	data := &models.Usuario{Usuario: body.Usuario, Clave: body.Clave, Dire: body.Dire, Tele: body.Tele, Correo: body.Correo}

	result := configs.DB.Model(&usuario).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &usuario)
}

func UsuarioDelete(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario
	configs.DB.Delete(&usuario, id)
	c.JSON(200, gin.H{"deleted": true})

}
