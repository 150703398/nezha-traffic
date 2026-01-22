package main

import (
	"github.com/你的用户名/nezha-traffic/dao"
	"github.com/你的用户名/nezha-traffic/router"
)

func main() {
	dao.InitDB()
	router.Run()
}
