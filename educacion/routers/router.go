// @APIVersion 1.0.0
// @Title API FINANCEUP Educacion
// @Description API para gestionar modulos, contenidos, lecciones y progreso educativo.
// @Contact Alejandro Zorro
// @TermsOfServiceUrl http://localhost:8080/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"educacion/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/modulo_educativo",
			beego.NSInclude(
				&controllers.ModuloEducativoController{},
			),
		),

		beego.NSNamespace("/contenido",
			beego.NSInclude(
				&controllers.ContenidoController{},
			),
		),

		beego.NSNamespace("/progreso_educativo",
			beego.NSInclude(
				&controllers.ProgresoEducativoController{},
			),
		),

		beego.NSNamespace("/leccion",
			beego.NSInclude(
				&controllers.LeccionController{},
			),
		),

		beego.NSNamespace("/progreso_leccion",
			beego.NSInclude(
				&controllers.ProgresoLeccionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
