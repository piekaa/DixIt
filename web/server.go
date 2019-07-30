package web

import (
	"dixit/game"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type Server struct {
	game game.Game
}

func (s *Server) init() {
	s.game = game.NewGame()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/test", s.test)
	e.GET("/api/pull/:roomId", s.pull)
	e.POST("/api/room", s.createRoom)
	e.POST("/api/name", s.addPlayer)
	e.POST("/api/ready", s.ready)
	e.POST("/api/me-first", s.meFirst)
	e.Static("/", "static")
	e.Logger.Fatal(e.Start(":6999"))
}

func (s *Server) test(c echo.Context) error {

	return c.JSON(http.StatusOK, &struct {
		Test string `json:"test"`
	}{"ok"})
}

func handleError(c echo.Context, err error) error {
	fmt.Println("Error: ", err)
	return c.JSON(http.StatusBadRequest, &struct {
		Message string `json:"message"`
	}{err.Error()})
}

func NewServer() *Server {
	s := &Server{}
	s.init()
	return s
}
