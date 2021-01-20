package web

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func (s *Server) meFirst(c echo.Context) error {

	body := struct {
		Name   string `json:"playerName"`
		RoomId string `json:"roomId"`
	}{}
	err := c.Bind(&body)

	if err != nil {
		fmt.Println("Error: ", err)
		return c.JSON(http.StatusBadRequest, &struct {
			Message string `json:"message"`
		}{"Cannot deserialize request body"})
	}

	success, err := s.game.ChooseFirst(body.RoomId, body.Name)

	if !success {
		return c.JSON(http.StatusConflict, &struct {
			Message string `json:"message"`
		}{"Someone elese was first"})
	}

	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, &struct {
	}{})
}
