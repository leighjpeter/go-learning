package routers

import (
	"github.com/astaxie/beego"
	"github.com/leighjpeter/go-learning/myapp/controllers"
	// "github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.MainController{})
	// beego.Get("/", func(ctx *context.Context) {
	// 	ctx.Output.Body([]byte("hello world"))
	// })·
	return
}
