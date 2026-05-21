package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	registerCRUD("CategoriaController")
	registerCRUD("MovimientoDineroController")
	registerCRUD("TipoIngresoController")
	registerCRUD("FinanzasController")
	registerCRUD("TipoInversionController")
	registerCRUD("NivelRiesgoController")
	registerCRUD("MovimientoInversionController")
	registerCRUD("TipoIngresoInversionController")
	registerCRUD("InversionController")
	registerCRUD("EditarMetaController")
	registerCRUD("MovimientoMetaController")
	registerCRUD("TipoIngresoMetaController")
	registerCRUD("MetaController")
}

func registerCRUD(controller string) {
	key := "finanzas/controllers:" + controller
	beego.GlobalControllerRouter[key] = append(beego.GlobalControllerRouter[key],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	)
}
