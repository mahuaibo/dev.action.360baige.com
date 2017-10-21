package main

import (
	_ "dev.action.360baige.com/database"
	_ "dev.action.360baige.com/rpc/server"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	beego.Get("/", func(ctx *context.Context) {

		ctx.WriteString("ok")
	})
	beego.Run()
}

