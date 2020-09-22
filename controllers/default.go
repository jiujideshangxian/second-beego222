package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "原神"
	c.Data["Email"] = "yuanshen.com"
	c.TplName = "index.tpl"
}
