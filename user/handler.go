package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type UserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type handlerNewInstance struct {
	repo Repository
}

func (h handlerNewInstance) handleGetByID(c echo.Context) error {
	code := http.StatusOK
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	result, err := h.repo.GetByID(id)
	var res map[string]interface{}
	if err != nil {
		res = map[string]interface{}{
			"err": getErrorMessage(err),
		}
	} else {
		res = map[string]interface{}{
			"user": result,
		}
	}
	return c.JSON(code, res)
}

func (h handlerNewInstance) handleGetAllUser(c echo.Context) error {
	code := http.StatusOK
	users, err := h.repo.GetAll()
	if err != nil {
		code = http.StatusInternalServerError
	}
	if len(users) == 0 {
		code = http.StatusNotFound
	}
	return c.JSON(code, users)
}

func (h handlerNewInstance) handleNewUser(c echo.Context) error {
	code := http.StatusOK
	newUser := UserRequest{}
	_ = c.Bind(&newUser)
	user, err := h.repo.Create(newUser)
	var res map[string]interface{}
	if err != nil {
		res = map[string]interface{}{
			"err": getErrorMessage(err),
		}
	} else {
		res = map[string]interface{}{
			"user": user,
		}
	}
	return c.JSON(code, res)
}

func (h handlerNewInstance) handleUpdate(c echo.Context) error {
	code := http.StatusOK
	request := UpdateUserRequest{}
	_ = c.Bind(&request)
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	err := h.repo.Update(request,id)
	var res map[string]interface{}
	if err != nil {
		code = http.StatusInternalServerError
		res = map[string]interface{}{
			"err": getErrorMessage(err),
		}
	}
	return c.JSON(code, res)
}

func (h handlerNewInstance) handleLogin(c echo.Context) error {
	code := http.StatusOK
	login := LoginRequest{}
	_ = c.Bind(&login)
	var res map[string]interface{}
	user, err := h.repo.Login(login)
	if err != nil {
		res = map[string]interface{}{
			"err": getErrorMessage(err),
		}
	} else {
		if user != nil {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["id"] = user.Id
			claims["name"] = user.Name
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
			t, err := token.SignedString([]byte("vbvb"))
			if err != nil {
				code = http.StatusInternalServerError
				res = map[string]interface{}{
					"err": getErrorMessage(err),
				}
			} else {
				res = map[string]interface{}{
					"token": t,
				}
			}
		} else {
			code = http.StatusNotFound
			res = map[string]interface{}{
				"err": "user not found",
			}
		}
	}
	return c.JSON(code, res)
}

func getErrorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
