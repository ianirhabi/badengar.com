package lihatsms

import (
	"net/http"

	"github.com/labstack/echo"
)

func Postsms(c echo.Context) (e error) {
	var r Requestsms
	id := c.Param("id")
	if err := c.Bind(&r); err == nil {
		if cc, m := Kirimsms(r, id); m == nil {
			return c.JSON(http.StatusOK, &cc)
		}
	}
	return e
	return
}
