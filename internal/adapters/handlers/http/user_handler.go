package http

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"remember-me/internal/domain/models"
	"remember-me/internal/domain/services"
	"remember-me/internal/utils"
	"strconv"
	"time"
)

type UserHandler struct {
	us *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{
		us: us,
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
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	user, err := uh.us.GetUserByEmail(user.Email)

	if err != nil {
		return c.String(404, "Not Found")
	}
	return c.Render(200, "user", user)
}

func (uh *UserHandler) PostUser(c echo.Context) error {
	u := models.User{}

	if err := c.Bind(&u); err != nil {
		return c.String(500, err.Error())
	}

	exists := uh.us.ExistsByEmail(u.Email)
	if exists {
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

	err := uh.us.DeleteUser(id)

	if err != nil {
		return c.String(500, "Internal Server Error")
	}

	return c.String(200, "User deleted!")
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&user); err != nil {
		return c.String(500, err.Error())
	}

	// Call the UserService to update the user.
	err := uh.us.UpdateUser(id, &user)
	if err != nil {
		// Handle errors, e.g., user not found or validation errors.
		return c.String(422, err.Error())
	}

	return c.Render(200, "user", user)
}

func (uh *UserHandler) Login(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	u, err := uh.us.GetUserByEmail(user.Email)

	valid := utils.ComparePassword(user.Password, u.Password)

	if !valid {
		return errors.New("invalid password")
	}

	token := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	//inmemory
	models.Sessions[token] = models.Session{
		UserID:    u.ID,
		ExpiresAt: expiresAt,
	}

	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: expiresAt,
	})

	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	return c.String(http.StatusOK, token)
}
