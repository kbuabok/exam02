package demo

import (
	"demo/database"
	"demo/user"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	resource, err := database.CreateResource()
	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()

	g := e.Group("/api/v1")
	user.NewRouterInstance(g, resource)
	e.Logger.Fatal(e.Start(":8080"))
}
