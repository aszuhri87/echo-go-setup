package controllers

import (
	"echo-go/app/models"
	"echo-go/app/repositories"
	"echo-go/app/utils"
	"echo-go/config"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Register godoc
// @Summary      Register User
// @Description  Register bro
// @Tags         register
// @Accept       json
// @Produce      json
// @Param data body models.Create true "The input todo struct"
// @Success      200  {object}  models.DataResponseOk
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /register [post]
func RegisterUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	data, err := repositories.RegisterUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, data)
}

// Login godoc
// @Summary      Login User
// @Description  Login bro
// @Tags         Login
// @Accept       json
// @Produce      json
// @Param data body models.Credentials true "The input todo struct"
// @Success      200  {object}  models.DataResponseOk
// @Failure      400  {object}  models.ResponseStatus
// @Failure      404  {object}  models.ResponseStatus
// @Failure      500  {object}  models.ResponseStatus
// @Router       /login [post]
func Login(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	data, err := repositories.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	password := utils.CheckPasswordHash(user.Password, data.Password)

	if !password {
		return c.JSON(http.StatusUnauthorized, password)
	}

	token := config.JwtMakeToken(data.ID)

	var expirationTime = time.Now().Add(24 * time.Hour)

	var result = models.ResponseToken{Name: data.Name, Username: data.Username, Token: "Bearer " + token}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = expirationTime
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, result)
}
