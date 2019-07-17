package web

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type Server struct {
}

func (s *Server) init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/test", s.test)
	e.Static("/", "static")
	e.Logger.Fatal(e.Start(":6999"))
}

func (s *Server) test(c echo.Context) error {

	return c.JSON(http.StatusOK, &struct {
		Test string `json:"test"`
	}{"ok"})
}

func NewServer() *Server {
	s := &Server{}
	s.init()
	return s
}
