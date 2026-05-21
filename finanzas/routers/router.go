// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"finanzas/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_inversion",
			beego.NSInclude(
				&controllers.TipoInversionController{},
			),
		),

		beego.NSNamespace("/inversion",
			beego.NSInclude(
				&controllers.InversionController{},
			),
		),

		beego.NSNamespace("/movimiento_dinero",
			beego.NSInclude(
				&controllers.MovimientoDineroController{},
			),
		),

		beego.NSNamespace("/editar_meta",
			beego.NSInclude(
				&controllers.EditarMetaController{},
			),
		),

		beego.NSNamespace("/tipo_ingreso_meta",
			beego.NSInclude(
				&controllers.TipoIngresoMetaController{},
			),
		),

		beego.NSNamespace("/meta",
			beego.NSInclude(
				&controllers.MetaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
