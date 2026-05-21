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
			beego.NSRouter("/", &controllers.ModuloEducativoController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.ModuloEducativoController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.ModuloEducativoController{},
			),
		),

		beego.NSNamespace("/contenido",
			beego.NSRouter("/", &controllers.ContenidoController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.ContenidoController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.ContenidoController{},
			),
		),

		beego.NSNamespace("/progreso_educativo",
			beego.NSRouter("/", &controllers.ProgresoEducativoController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.ProgresoEducativoController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.ProgresoEducativoController{},
			),
		),

		beego.NSNamespace("/leccion",
			beego.NSRouter("/", &controllers.LeccionController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.LeccionController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.LeccionController{},
			),
		),

		beego.NSNamespace("/progreso_leccion",
			beego.NSRouter("/", &controllers.ProgresoLeccionController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.ProgresoLeccionController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.ProgresoLeccionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
