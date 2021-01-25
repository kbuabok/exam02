package user

import (
	"demo/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouterInstance(app *echo.Group, resource *database.Resource) {
	repository := newRepoInstance(resource)
	h := handlerNewInstance{repository}
	app.POST("/login", h.handleLogin, validateLoginRequest)
	app.POST("/register", h.handleNewUser, validateUserRequest)
	app.GET("/users", h.handleGetAllUser)

	app.GET("/user", h.handleGetByID, middleware.JWT([]byte("vbvb")))
	app.PUT("/user", h.handleUpdate, middleware.JWT([]byte("vbvb")), validateUpdateUserRequest)
}

type Middleware struct {
	Context echo.Context
}
