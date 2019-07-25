package web

import (
	"github.com/labstack/echo"
	"net/http"
)

func (s *Server) pull(c echo.Context) error {
	roomId := c.Param("roomId")
	response, err := s.game.Pull(roomId)

	if err != nil {
		return handleError(c, err)
	}
	return c.JSON(http.StatusOK, response)
}
