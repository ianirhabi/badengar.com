package user

import (
	"fmt"

	"badengar.com/badengar/model"
	orm "github.com/alfatih/beego/orm"
)

func PostLogin(a Requestlogin) (res interface{}, e error) {

	return
}

func PostRegis(a RequestRegis) (res interface{}, e error) {
	o := orm.NewOrm()
	var data Respons

	user := model.BadengarUser{Nama: a.Nama,
		Username: a.Username,
		Password: a.Password}
	if df, e := o.Insert(&user); e == nil {
		data.Status = "berhasil"
		data.Data = &user
	} else {
		fmt.Println("debug kirim regis gagal === ", df, e)
		data.Status = "gagal"
	}

	return data, e
}
