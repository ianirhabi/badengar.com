package list_situs

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Show(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	if cc, m := Showlist(id); m == nil {
		return c.JSON(http.StatusOK, &cc)
	}
	return e
}
