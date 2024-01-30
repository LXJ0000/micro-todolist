package main

import (
	"micro-todolist/app/user/repository/dao"
	"micro-todolist/conf"
)

func main() {
	conf.Init()
	dao.Init()
}
