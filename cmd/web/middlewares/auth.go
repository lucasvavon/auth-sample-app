package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

// Clé secrète pour signer et vérifier les tokens JWT
var secretKey = []byte("votre_clé_secrète")

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Récupère le token dans l'en-tête Authorization
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token manquant"})
		}

		// Le format attendu est "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Format du token incorrect"})
		}

		// Récupère le token
		tokenString := parts[1]

		// Parse le token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Vérifie la méthode de signature du token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Méthode de signature invalide")
			}
			return secretKey, nil
		})

		// Vérifie si le token est valide
		if err != nil || !token.Valid {
			log.Println("Erreur lors de la validation du token:", err)
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token invalide"})
		}

		// Passe à la requête suivante si le token est valide
		return next(c)
	}
}
