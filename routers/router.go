// @APIVersion 1.0.0
// @Title beego Eventos Crud
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/eventos_campus_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns1 := beego.NewNamespace("/v1",

		beego.NSNamespace("/calendario_evento",
			beego.NSInclude(
				&controllers.CalendarioEventoController{},
			),
		),

		beego.NSNamespace("/tipo_recurrencia",
			beego.NSInclude(
				&controllers.TipoRecurrenciaController{},
			),
		),

		beego.NSNamespace("/tipo_evento",
			beego.NSInclude(
				&controllers.TipoEventoController{},
			),
		),

		beego.NSNamespace("/rol_encargado_evento",
			beego.NSInclude(
				&controllers.RolEncargadoEventoController{},
			),
		),

		beego.NSNamespace("/encargado_evento",
			beego.NSInclude(
				&controllers.EncargadoEventoController{},
			),
		),

		beego.NSNamespace("/tipo_publico",
			beego.NSInclude(
				&controllers.TipoPublicoController{},
			),
		),

		beego.NSNamespace("/tr_evento",
			beego.NSInclude(
				&controllers.TrEventoController{},
			),
		),

		beego.NSNamespace("/tipo_recurrencia",
			beego.NSInclude(
				&controllers.TipoRecurrenciaController{},
			),
		),
	)

	ns2 := beego.NewNamespace("/v2",

		beego.NSNamespace("/rol_encargado_evento",
			beego.NSInclude(
				&controllers.RolEncargadoEventoController{},
			),
		),

		beego.NSNamespace("/tipo_publico",
			beego.NSInclude(
				&controllers.TipoPublicoController{},
			),
		),

		beego.NSNamespace("/tipo_recurrencia",
			beego.NSInclude(
				&controllers.TipoRecurrenciaController{},
			),
		),

		beego.NSNamespace("/tipo_evento",
			beego.NSInclude(
				&controllers.TipoEventoController{},
			),
		),

		beego.NSNamespace("/tr_evento",
			beego.NSInclude(
				&controllers.TrEventoController{},
			),
		),

		beego.NSNamespace("/calendario_evento",
			beego.NSInclude(
				&controllers.CalendarioEventoController{},
			),
		),

		beego.NSNamespace("/encargado_evento",
			beego.NSInclude(
				&controllers.EncargadoEventoController{},
			),
		),
	)

	beego.AddNamespace(ns1)
	beego.AddNamespace(ns2)
}
