package controllers

import (
	"github.com/astaxie/beego"
	"projetoBeego/beegoProject/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	m := models.Pessoa{}
	c.Data["Pessoas"] = m.GetAll()
	c.Data["Website"] = "about.me/diegofernando"
	c.Data["Email"] = "diegofernando5672@gmail.com"
	c.TplName = "index.tpl"
}
