package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["TrueCond"] = true

	type u struct {
		Name string
		Age  int
	}

	user := &u{
		Name: "leighj",
		Age:  10,
	}
	c.Data["user"] = user

	nums := []int{1, 2, 5, 6, 7, 9}
	c.Data["nums"] = nums

	c.Data["TplVar"] = "hey guys"

	c.Data["Html"] = "<div>html</div>"

}
