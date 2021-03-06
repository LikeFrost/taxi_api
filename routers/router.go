// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"taxi_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/car",
			beego.NSInclude(
				&controllers.CarController{},
			),
		),
		beego.NSNamespace("/question",
			beego.NSInclude(
				&controllers.QuestionController{},
			),
		),
		beego.NSNamespace("/traffic",
			beego.NSInclude(
				&controllers.TrafficController{},
			),
		),
		beego.NSNamespace("/money",
			beego.NSInclude(
				&controllers.MoneyController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
