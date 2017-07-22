package routers

import (
	"github.com/goagcode/gofunda/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
