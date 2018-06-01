package routers

import (
	"caidao/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/filelist", &controllers.MainController{}, "get:Filelist;post:Filelist")
	beego.Router("/fileupload", &controllers.MainController{}, "post:Fileupload")
	beego.Router("/rename", &controllers.MainController{}, "post:Rename")
	beego.Router("/delete", &controllers.MainController{}, "post:Delete")
	beego.Router("/download", &controllers.MainController{}, "post:Download")
	beego.Router("/newFolder", &controllers.MainController{}, "post:NewFolder")
	beego.Router("/move", &controllers.MainController{}, "post:Move")
	beego.Router("/copy", &controllers.MainController{}, "post:Copy")
	beego.Router("/search", &controllers.MainController{}, "post:Search")
	beego.Router("/wenjian", &controllers.MainController{}, "get:Wenjian")
	beego.Router("/listjson", &controllers.MainController{}, "get:Listjson")
	beego.Router("/listadd", &controllers.MainController{}, "get:Addurl")
	beego.Router("/Add", &controllers.MainController{}, "post:Add")
	beego.Router("/deleteurl", &controllers.MainController{}, "post:DeleteUrl")

}
