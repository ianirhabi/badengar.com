package list_situs

import (
	"fmt"

	"badengar.com/badengar/model"
	"github.com/alfatih/beego/orm"
	"github.com/labstack/echo"
)

func Showlist(id int) (d interface{}, e error) {
	o := orm.NewOrm()
	var cc []*model.SitusList
	var b Respons
	if _, d := o.Raw("SELECT * FROM situs_list where id_user = ?", id).QueryRows(&cc); d == nil {
		fmt.Println("debug error === ", d)
		b.Status = "berhasil"
		b.Data = cc
	} else {
		fmt.Println("debug error === ", d)
		b.Status = "gagal"
	}
	return b, e
}

func Getphoto(c echo.Context) (e error) {
	ifile := c.Param("namafoto")
	fmt.Println("masuk")
	return c.File("/home/ubuntu/go/src/badengar.com/badengar/foto/" + ifile)
}
