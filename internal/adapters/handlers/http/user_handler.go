package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"remember-me/internal/domain/models"
	"remember-me/internal/domain/usecases"
	"strconv"
)

type UserHandler struct {
	us *usecases.UserService
}

func NewUserHandler(us *usecases.UserService) *UserHandler {
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

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	user, err := uh.us.GetUserByID(id)

	if err != nil {
		return c.String(404, "Not Found")
	}
	return c.Render(200, "user", user)
}

func (uh *UserHandler) PostUser(c echo.Context) error {
	u := models.UserDTO{}

	if err := c.Bind(&u); err != nil {
		return c.String(500, err.Error())
	}

	err := uh.us.CreateUser(&u)
	if err != nil {
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
	var updateUser models.UserDTO
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&updateUser); err != nil {
		return c.String(500, err.Error())
	}

	// Call the UserService to update the user.
	err := uh.us.UpdateUser(id, &updateUser)
	if err != nil {
		// Handle errors, e.g., user not found or validation errors.
		return c.String(422, err.Error())
	}

	return c.Render(200, "updateUser", updateUser)
}
