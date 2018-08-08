package model

import orm "github.com/alfatih/beego/orm"

func init() {
	orm.RegisterModel(new(SitusList))
}

type SitusList struct {
	Id       int64  `orm:"column(id);auto"json:"id"`
	IdUser   int64  `orm:"column(id_user);size(100)"json:"id_user"`
	List     string `orm:"column(list);size(100)"json:"list_situs"`
	Namafoto string `orm:"column(namafoto);size(100)"json:"nama_foto"`
}
