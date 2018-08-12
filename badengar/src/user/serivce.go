package user

import (
	"fmt"

	"badengar.com/badengar/model"
	orm "github.com/alfatih/beego/orm"
)

func PostLogin(a Requestlogin) (res interface{}, e error) {
	return
}

type df struct {
	nama string
}

func PostRegis(a RequestRegis, imei string) (res interface{}, e error) {
	o := orm.NewOrm()
	var data Respons
	var validasi []*model.BadengarUser

	if f, d := o.Raw("SELECT * FROM badengar_user where username =?", imei).QueryRows(&validasi); d == nil {
		fmt.Println("imeinya ", f, d, imei, validasi[0].Id)
		//Conversi Interface
		// var token = map[string]interface{}{
		// 	"nama": *validasi[0],
		// }

		//fmt.Println(validasi.Id)
		// s := strings.Split(*validasi[0], "{")
		// id := s[0]
		// fmt.Println("id nya ", id)
		if f == 0 {
			fmt.Println("masuk sini")
			user := model.BadengarUser{Nama: a.Nama,
				Username: a.Username,
				Password: a.Password,
				Tanggal:  a.Tanggal,
				Jam:      a.Jam}
			if df, e := o.Insert(&user); e == nil {
				data.Status = "berhasil insert"
				data.Data = &user
			} else {
				fmt.Println("debug kirim regis gagal === ", df, e)
				data.Status = "gagal"
			}
		} else {
			imeinya := model.BadengarUser{Id: validasi[0].Id}
			if o.Read(&imeinya) == nil {
				imeinya.Password = a.Password
				imeinya.Tanggal = a.Tanggal
				imeinya.Jam = a.Jam
				if _, err := o.Update(&imeinya); err == nil {
					data.Status = "berhasil update"
					data.Data = &imeinya
				} else {
					data.Status = "gagal"
				}
			} else {
				fmt.Println("gagal baca data")
			}
		}

	} else {
		fmt.Println("kesalahan pada mysql")
	}
	return data, e
}
