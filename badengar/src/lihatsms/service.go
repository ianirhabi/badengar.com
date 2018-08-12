package lihatsms

import (
	"fmt"

	"badengar.com/badengar/model"
	orm "github.com/alfatih/beego/orm"
)

func Kirimsms(a Requestsms, id string) (res interface{}, e error) {
	o := orm.NewOrm()
	var data Respons

	sms := model.Smsuser{Userid: id,
		Number:  a.Number,
		Message: a.Message}
	if df, e := o.Insert(&sms); e == nil {
		data.Status = "berhasil"
		data.Data = &sms
	} else {
		fmt.Println("debug kirim sms gagal === ", df, e)
		data.Status = "gagal"
	}
	return data, e
}
