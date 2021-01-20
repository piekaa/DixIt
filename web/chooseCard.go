package web

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func (s *Server) chooseCard(c echo.Context) error {

	body := struct {
		Name   string `json:"playerName"`
		RoomId string `json:"roomId"`
		MyCard int    `json:"myCard"`
		MyType int    `json:"myType"`
	}{}
	err := c.Bind(&body)

	if err != nil {
		fmt.Println("Error: ", err)
		return c.JSON(http.StatusBadRequest, &struct {
			Message string `json:"message"`
		}{"Cannot deserialize request body"})
	}

	err = s.game.ChooseCards(body.RoomId, body.Name, body.MyCard, body.MyType)

	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, &struct {
	}{})
}
