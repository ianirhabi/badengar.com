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
	Sim1     string `orm:"column(sim1);size(100)"`
	Sim2     string `orm:"column(sim2);size(100)"`
	Number   string `orm:"column(number);size(100)"`
	Tanggal  string `orm:"column(tanggal);size(100)"`
	Jam      string `orm:"column(jam);size(100)"`
}
