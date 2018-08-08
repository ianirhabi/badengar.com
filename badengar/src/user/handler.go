package user

import (
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) (e error) {

	var r Requestlogin
	if err := c.Bind(&r); err == nil {
		if cc, m := PostLogin(r); m == nil {
			return c.JSON(http.StatusOK, &cc)
		}
	}
	return e
}

func Regis(c echo.Context) (e error) {
	var r RequestRegis
	if err := c.Bind(&r); err == nil {
		if cc, m := PostRegis(r); m == nil {
			return c.JSON(http.StatusOK, &cc)
		}
	}
	return e
}
