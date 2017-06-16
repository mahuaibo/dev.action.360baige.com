package main

import (
	_ "dev.action.360baige.com/database"
	_ "dev.action.360baige.com/rpc/server"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
