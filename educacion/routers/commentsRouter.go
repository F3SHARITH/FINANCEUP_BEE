package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ContenidoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ContenidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ContenidoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ContenidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ContenidoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ContenidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ContenidoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ContenidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ContenidoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ContenidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:LeccionController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:LeccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:LeccionController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:LeccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:LeccionController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:LeccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:LeccionController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:LeccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:LeccionController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:LeccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloContenidoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloContenidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloContenidoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloContenidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloContenidoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloContenidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloContenidoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloContenidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloContenidoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloContenidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloEducativoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloEducativoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloEducativoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloEducativoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloEducativoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloEducativoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloEducativoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloEducativoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloEducativoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ModuloEducativoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ProgresoEducativoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ProgresoEducativoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ProgresoEducativoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ProgresoEducativoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ProgresoEducativoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ProgresoEducativoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ProgresoEducativoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ProgresoEducativoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ProgresoEducativoController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:ProgresoEducativoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:QuizController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:QuizController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:QuizController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:QuizController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:QuizController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:QuizController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:QuizController"] = append(beego.GlobalControllerRouter["FINANCEUP_BEE/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
