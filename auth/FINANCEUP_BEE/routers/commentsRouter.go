package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:AuditoriaLoginController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:AuditoriaLoginController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:AuditoriaLoginController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:AuditoriaLoginController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:AuditoriaLoginController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:AuditoriaLoginController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:AuditoriaLoginController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:AuditoriaLoginController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:AuditoriaLoginController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:AuditoriaLoginController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:CredencialController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:CredencialController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:CredencialController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:CredencialController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:CredencialController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:CredencialController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:CredencialController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:CredencialController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:CredencialController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:CredencialController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:RolController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:RolController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:RolController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:RolController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:RolController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:RolController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:RolController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:RolController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioRolController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioRolController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioRolController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioRolController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioRolController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioRolController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioRolController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioRolController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioRolController"] = append(beego.GlobalControllerRouter["auth/FINANCEUP_BEE/controllers:UsuarioRolController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
