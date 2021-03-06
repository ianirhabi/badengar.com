package model

import orm "github.com/alfatih/beego/orm"

func init() {
	orm.RegisterModel(new(Smsuser))
}

type Smsuser struct {
	Id      int64  `orm:"column(id);auto"json:"id"`
	Userid  string `orm:"column(user_id);size(100)"json:"userid"`
	Number  string `orm:"column(number);size(100)"json:"number"`
	Message string `orm:"column(message);size(100)"json:"pesan"`
}
