package web

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

func (s *Server) createRoom(c echo.Context) error {

	id := uuid.New().String()
	s.game.Start(id)
	return c.JSON(http.StatusOK, &struct {
		RoomId string `json:"roomId"`
	}{id})
}
