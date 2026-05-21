// @APIVersion 1.0.0
// @Title API FINANCEUP Negocio
// @Description API para gestionar bancos, productos crediticios, asesores, leads, creditos y comisiones.
// @Contact Alejandro Zorro
// @TermsOfServiceUrl http://localhost:8080/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"negocio/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/banco",
			beego.NSRouter("/", &controllers.BancoController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.BancoController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.BancoController{},
			),
		),

		beego.NSNamespace("/producto_crediticio",
			beego.NSRouter("/", &controllers.ProductoCrediticioController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.ProductoCrediticioController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.ProductoCrediticioController{},
			),
		),

		beego.NSNamespace("/asesor_bancario",
			beego.NSRouter("/", &controllers.AsesorBancarioController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.AsesorBancarioController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.AsesorBancarioController{},
			),
		),

		beego.NSNamespace("/contacto_asesor",
			beego.NSRouter("/", &controllers.ContactoAsesorController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.ContactoAsesorController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.ContactoAsesorController{},
			),
		),

		beego.NSNamespace("/lead",
			beego.NSRouter("/", &controllers.LeadController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.LeadController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.LeadController{},
			),
		),

		beego.NSNamespace("/conversacion_usuario_asesor",
			beego.NSRouter("/", &controllers.ConversacionUsuarioAsesorController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.ConversacionUsuarioAsesorController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.ConversacionUsuarioAsesorController{},
			),
		),

		beego.NSNamespace("/credito_desembolsado",
			beego.NSRouter("/", &controllers.CreditoDesembolsadoController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.CreditoDesembolsadoController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.CreditoDesembolsadoController{},
			),
		),

		beego.NSNamespace("/transaccion_comision",
			beego.NSRouter("/", &controllers.TransaccionComisionController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/:id", &controllers.TransaccionComisionController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSInclude(
				&controllers.TransaccionComisionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
