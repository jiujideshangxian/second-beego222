package controllers

import (
	"beetwo2/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
	/*name1:=c.GetString("name")
	age1,_:=c.GetInt("")*/

	c.Data["Website"] = "原神"
	c.Data["Email"] = "yuanshen.com"
	user:=c.Ctx.Input.Query("user")
	fmt.Println(user)
	psd:=c.Ctx.Input.Query("psd")
	fmt.Println(psd)

	if user!="yuanshen"||psd!="123456" {
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据不对"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据正确"))
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {

	/*dataBytes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据请求失败，请重试")

	}*/
	//body:=c.Ctx.Request.Body
	dataByes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据加载失败，请重试")
		return
	}
	//json包解析
	var person models.Person
	err=json.Unmarshal(dataByes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("用户名",person.User,"年龄",person.Age)
	c.Ctx.WriteString("用户名是："+person.User)
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