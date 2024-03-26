package controllers

import (
	"echo-go/app/models"
	"echo-go/app/repositories"
	"echo-go/config"
	"echo-go/response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ShowUser godoc
// @Summary      Show all users
// @Description  get string by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ListResponseOk
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /user [get]
func GetUser(c echo.Context) error {
	user := []models.User{}
	result := []models.ResponseData{}
	raw := models.ResponseData{}

	data, err := repositories.GetUser(user)
	for i := 0; i < len(data); i++ {
		raw = models.ResponseData{ID: data[i].ID, Name: data[i].Name, Username: data[i].Username}

	}
	result = append(result, raw)

	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, result)
}

// CreateUser godoc
// @Summary      Create an User
// @Description  create
// @Tags         users
// @Accept       json
// @Produce      json
// @Param data body models.Create true "The input todo struct"
// @Success      200  {object}  models.DataResponseOk
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /user [post]
func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	data, err := repositories.CreateUser(user)
	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, data)
}

// ShowUserFirst godoc
// @Security Bearer
// @Summary      Show an User By ID
// @Description  get User by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Success      200  {object}  models.DataResponseOk
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /user/{id} [get]
func GetUserByID(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	id := uuid.MustParse(c.Param("id"))

	data, err := repositories.GetUserByID(user, id)
	result := models.ResponseData{ID: data.ID, Name: data.Name, Username: data.Username}

	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, result)
}

// Profile godoc
// @Security Bearer
// @Summary      Show profile
// @Description  get User profile
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.DataResponseOk
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /user/profile [get]
func UserProfile(c echo.Context) error {
	user := models.User{}

	token := c.Get("user").(*jwt.Token)

	id := config.JwtUserID(token)

	data, err := repositories.GetUserByID(user, id)
	result := models.ResponseData{ID: data.ID, Name: data.Name, Username: data.Username}

	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, result)
}

// UpdateUser godoc
// @Security Bearer
// @Summary      Update an User
// @Description  Update
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Param data body models.Create true "The input todo struct"
// @Success      200  {object}  models.DataResponseOk
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /user/{id} [put]
func UpdateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	id := uuid.MustParse(c.Param("id"))

	data, err := repositories.UpdateUser(user, id)
	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, data)
}

// DeleteUser godoc
// @Security Bearer
// @Summary      Delete an User
// @Description  Delete
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Success      200  {object}  models.StatusOk
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /user/{id} [delete]
func DeleteUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	id := uuid.MustParse(c.Param("id"))

	data, err := repositories.DeleteUser(user, id)
	if err != nil {
		return response.InternalServerError(c)
	}
	return response.Success(c, data)
}
