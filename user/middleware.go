package user

import (
	"demo/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func validateUserRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		newUser := UserRequest{}
		if err := c.Bind(&newUser); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := utils.ValidateRequest(newUser); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return next(c)
	}
}

func validateLoginRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		login := LoginRequest{}
		if err := c.Bind(&login); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := utils.ValidateRequest(login); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return next(c)
	}
}

func validateUpdateUserRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := UpdateUserRequest{}
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := utils.ValidateRequest(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return next(c)
	}
}