package main

import (
	"github.com/nezhahq/nezha/dao"
	"github.com/nezhahq/nezha/router"
)

func main() {
	dao.InitDB()
	router.Run()
}
