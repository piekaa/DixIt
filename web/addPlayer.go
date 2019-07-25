package web

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func (s *Server) addPlayer(c echo.Context) error {

	body := struct {
		Name   string `json:"name"`
		RoomId string `json:"roomId"`
	}{}
	err := c.Bind(&body)

	if err != nil {
		fmt.Println("Error: ", err)
		return c.JSON(http.StatusBadRequest, &struct {
			Message string `json:"message"`
		}{"Cannot deserialize request body"})
	}

	success, err := s.game.ChooseName(body.RoomId, body.Name)
	if err != nil {
		return handleError(c, err)
	}

	if !success {
		return c.JSON(http.StatusConflict, &struct {
		}{})
	}

	return c.JSON(http.StatusOK, &struct {
	}{})
}
