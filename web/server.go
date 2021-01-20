package web

import (
	"dixit/game"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strings"
)

type Server struct {
	port         string
	externalHost string
	externalPort string
	game         game.Game
}

func (s *Server) init() {
	s.game = game.NewGame()
	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/test", s.test)
	e.POST("/api/upper", s.upper)
	e.GET("/api/pull/:roomId", s.pull)
	e.POST("/api/room", s.createRoom)
	e.POST("/api/playerName", s.addPlayer)
	e.POST("/api/ready", s.ready)
	e.POST("/api/me-first", s.meFirst)
	e.POST("/api/choose-cards", s.chooseCard)

	e.GET("/api/qr", s.qr)

	e.Static("/", "static")
	e.Logger.Fatal(e.Start(":" + s.port))
}

func (s *Server) test(c echo.Context) error {

	return c.JSON(http.StatusOK, &struct {
		Test string `json:"test"`
	}{"ok"})
}

func (s *Server) upper(c echo.Context) error {

	var m map[string] interface{}
	err := c.Bind(&m)
	if err != nil {
		return handleError(c, err)
	}

	for k, v := range m {
		s, ok := v.(string)
		if ok {
			m[k] = strings.ToUpper(s)
		}
	}

	return c.JSON(http.StatusOK, m)
}

func handleError(c echo.Context, err error) error {
	fmt.Println("Error: ", err)
	return c.JSON(http.StatusBadRequest, &struct {
		Message string `json:"message"`
	}{err.Error()})
}

func NewServer(port, externalHost, externalPort string) *Server {
	s := &Server{port: port, externalHost: externalHost, externalPort: externalPort}
	s.init()
	return s
}
