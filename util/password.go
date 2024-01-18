package util

import (
	"golang.org/x/crypto/bcrypt"
)

// CreatePassword crea un hash de la contraseña con bcrypt
func CreatePassword(clave string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(clave), 10)
	return string(bytes), err
}

// CheckPassword verifica si la contraseña es correcta con respecto al hash
func CheckPassword(hashedPassword, clave string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(clave))
	return err == nil
}
