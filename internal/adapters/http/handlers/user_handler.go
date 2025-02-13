package handlers

import (
	"auth-sample-app/internal/domain/models"
	"auth-sample-app/internal/domain/usecases"
	"auth-sample-app/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	us *usecases.UserService
	ss *usecases.SessionService
}

func NewUserHandler(us *usecases.UserService, ss *usecases.SessionService) *UserHandler {
	return &UserHandler{
		us: us,
		ss: ss,
	}
}

func (uh *UserHandler) GetUsers(c echo.Context) error {
	users, err := uh.us.GetUsers()

	if err != nil {
		return c.String(404, "Not Found")
	}

	return c.Render(200, "login", users)
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	var user *models.User

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	user, err := uh.us.GetUserByEmail(user.Email)

	if err != nil {
		return c.String(404, "GetUser Not Found")
	}
	return c.Render(200, "user", user)
}

func (uh *UserHandler) PostUser(c echo.Context) error {
	u := models.User{}

	if err := c.Bind(&u); err != nil {
		return c.String(500, err.Error())
	}

	user, _ := uh.us.GetUserByEmail(u.Email)

	if user != nil {
		return c.String(422, "user with this email already exists")
	}

	if err := uh.us.CreateUser(&u); err != nil {
		return c.String(422, err.Error())
	}

	if c.Request().Header.Get("HX-Request") != "" {
		c.Response().Header().Set("HX-Redirect", "/login")
	}

	return c.NoContent(http.StatusOK)
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := uh.us.DeleteUser(uint(id))

	if err != nil {
		return c.String(500, "Internal Server Error")
	}

	return c.String(200, "User deleted!")
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	var user *models.User
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&user); err != nil {
		return c.String(500, err.Error())
	}

	// Call the UserService to update the user.
	err := uh.us.UpdateUser(uint(id), user)
	if err != nil {
		// Handle errors, e.g., user not found or validation errors.
		return c.String(422, err.Error())
	}

	return c.Render(200, "user", user)
}

func (uh *UserHandler) Login(c echo.Context) error {
	var user *models.User

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	u, err := uh.us.GetUserByEmail(user.Email)
	passwordValid := utils.ComparePassword(user.Password, u.Password)

	if err != nil || !passwordValid {
		return c.String(http.StatusUnauthorized, models.ErrInvalidCredentials.Error())
	}

	sessionID := uuid.NewString()
	err = uh.ss.CreateSession(c.Request().Context(), sessionID, u.ID)

	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(cookie)

	if c.Request().Header.Get("HX-Request") != "" {
		c.Response().Header().Set("HX-Redirect", "/")
	}

	return c.NoContent(http.StatusOK)
}

func (uh *UserHandler) Logout(c echo.Context) error {

	sessionID, _ := c.Cookie("session_id")
	err := uh.ss.InvalidateSession(c.Request().Context(), sessionID.Value)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	expiredCookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0), // Set an expiration date in the past
		MaxAge:   -1,              // Explicitly indicate that the cookie should be removed
		HttpOnly: true,
	}
	c.SetCookie(expiredCookie)

	if c.Request().Header.Get("HX-Request") != "" {
		c.Response().Header().Set("HX-Redirect", "/login")
	}

	return c.NoContent(http.StatusOK)
}
