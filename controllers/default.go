package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
	c.Data["Website"] = "原神"
	c.Data["Email"] = "yuanshen.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	fmt.Println("post服务器")
	user:=c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为",user)
	psd:=c.Ctx.Request.FormValue("psd")
	fmt.Println("密码为",psd)

	if user!="yuanshen"||psd!="123456" {
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据不值钱"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据值钱"))

	c.Data["Website"] = "原神"
	c.Data["Email"] = "yuanshen.com"
	c.TplName = "index.tpl"
}