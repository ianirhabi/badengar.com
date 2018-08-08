package main

import (
	"badengar.com/badengar/engine"
	orm "github.com/alfatih/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:ian!@#$%^@tcp(localhost:3306)/badengar_app?charset=utf8")
}

func main() {
	engine.Router()
}
