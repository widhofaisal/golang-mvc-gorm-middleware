package controller

import (
	"github.com/labstack/echo/v4"
	"learn/mvc/config"
	"learn/mvc/middleware"
	"learn/mvc/model"
	"net/http"
)

func GetUsersController(c echo.Context) error {
	var users []model.User
	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})

}

func CreateUser(c echo.Context) error {
	// binding data
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func LoginUserController(c echo.Context) error {
	// binding data
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Where("email=? AND password=?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	userResponse := model.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    userResponse,
	})
}
