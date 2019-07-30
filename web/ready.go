package web

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func (s *Server) ready(c echo.Context) error {

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

	err = s.game.Ready(body.RoomId, body.Name)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, &struct {
	}{})
}
