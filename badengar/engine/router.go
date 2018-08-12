package engine

import (
	"badengar.com/badengar/src/lihatsms"
	"badengar.com/badengar/src/list_situs"
	"badengar.com/badengar/src/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Router() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	endpoin_badengar := e.Group("/badengar/user")
	endpoin_badengar.POST("", user.Login)
	endpoin_badengar.POST("/register/:imei", user.Regis)

	endpoin_lis := e.Group("/badengar/list")
	endpoin_lis.GET("/:id", list_situs.Show)
	endpoin_lis.GET("/getfoto/:namafoto", list_situs.Getphoto)

	endpoin_lihatsms := e.Group("/badengar/sms")
	endpoin_lihatsms.POST("/:id", lihatsms.Postsms)

	e.Logger.Fatal(e.Start(":4000"))
}
