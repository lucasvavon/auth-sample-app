package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userID int) (string, error) {
	// Crée un objet JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,                                // L'ID de l'utilisateur
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expiration du token après 24 heures
	})

	// Signature du token avec la clé secrète
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
