package model

import orm "github.com/alfatih/beego/orm"

func init() {
	orm.RegisterModel(new(BadengarUser))
}

type BadengarUser struct {
	Id       int    `orm:"column(id);auto"json:"-"`
	Username string `orm:"column(username);size(100)"`
	Password string `orm:"column(password);size(100)"`
	Nama     string `orm:"column(nama);size(100)"`
}
