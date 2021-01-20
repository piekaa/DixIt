package web

import (
	"bytes"
	"dixit/web/qr"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func (s *Server) qr(c echo.Context) error {

	url := c.QueryParam("url")

	//url := s.externalize("room/join.html?id=" + roomId)

	fmt.Println(url)

	image, err := qr.Create(url)
	if err != nil {
		return handleError(c, err)
	}

	return c.Stream(http.StatusOK, "image/png", bytes.NewReader(image))
}
