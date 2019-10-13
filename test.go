package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

//restful controller 路由
type RESTfulController struct {
	beego.Controller
}
type RegExpController struct {
	beego.Controller
}

func (this *RESTfulController) Get() {
	this.Ctx.WriteString("get hello world")
}

func (this *RESTfulController) Post() {
	this.Ctx.WriteString("post hello world")
}

func (this *RegExpController) Get() {

	splat := this.Ctx.Input.Param(":splat")
	this.Ctx.WriteString(fmt.Sprintln("splat: %s\n", splat))
	path := this.Ctx.Input.Param(":path")
	this.Ctx.WriteString(fmt.Sprintln("path: ", path))
	ext := this.Ctx.Input.Param(":ext")
	this.Ctx.WriteString(fmt.Sprintln("ext: ", ext))

}

func main() {
	//restful controller 路由
	beego.Router("/RESTful", &RESTfulController{})
	beego.Router("/RegExp/*", &RegExpController{})

	beego.Run("127.0.0.1:8081")

}
