package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"remember-me/internal/domain/services"
)

func AuthMiddleware(ss *services.SessionService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			sessionID, err := c.Cookie("session_id")

			if err != nil || sessionID.Value == "" {

				return c.Redirect(http.StatusFound, "/login")
			}

			userID, err := ss.ValidateSession(context.Background(), sessionID.Value)

			if err != nil {
				return c.Redirect(http.StatusFound, "/login")
			}

			c.Set("userID", userID)

			return next(c)
		}
	}
}
