package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["finanzas/controllers:EditarMetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:EditarMetaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:EditarMetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:EditarMetaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:EditarMetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:EditarMetaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:EditarMetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:EditarMetaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:EditarMetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:EditarMetaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:InversionController"] = append(beego.GlobalControllerRouter["finanzas/controllers:InversionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:InversionController"] = append(beego.GlobalControllerRouter["finanzas/controllers:InversionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:InversionController"] = append(beego.GlobalControllerRouter["finanzas/controllers:InversionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:InversionController"] = append(beego.GlobalControllerRouter["finanzas/controllers:InversionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:InversionController"] = append(beego.GlobalControllerRouter["finanzas/controllers:InversionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:MetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:MetaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:MetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:MetaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:MetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:MetaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:MetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:MetaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:MetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:MetaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:MovimientoDineroController"] = append(beego.GlobalControllerRouter["finanzas/controllers:MovimientoDineroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:MovimientoDineroController"] = append(beego.GlobalControllerRouter["finanzas/controllers:MovimientoDineroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:MovimientoDineroController"] = append(beego.GlobalControllerRouter["finanzas/controllers:MovimientoDineroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:MovimientoDineroController"] = append(beego.GlobalControllerRouter["finanzas/controllers:MovimientoDineroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:MovimientoDineroController"] = append(beego.GlobalControllerRouter["finanzas/controllers:MovimientoDineroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:TipoIngresoMetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:TipoIngresoMetaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:TipoIngresoMetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:TipoIngresoMetaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:TipoIngresoMetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:TipoIngresoMetaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:TipoIngresoMetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:TipoIngresoMetaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:TipoIngresoMetaController"] = append(beego.GlobalControllerRouter["finanzas/controllers:TipoIngresoMetaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:TipoInversionController"] = append(beego.GlobalControllerRouter["finanzas/controllers:TipoInversionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:TipoInversionController"] = append(beego.GlobalControllerRouter["finanzas/controllers:TipoInversionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:TipoInversionController"] = append(beego.GlobalControllerRouter["finanzas/controllers:TipoInversionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:TipoInversionController"] = append(beego.GlobalControllerRouter["finanzas/controllers:TipoInversionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["finanzas/controllers:TipoInversionController"] = append(beego.GlobalControllerRouter["finanzas/controllers:TipoInversionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
