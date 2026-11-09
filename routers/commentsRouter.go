package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ActividadEconomicaProveedorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:EstadoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:EstadoProveedorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:EstadoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:EstadoProveedorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:EstadoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:EstadoProveedorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:EstadoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:EstadoProveedorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:EstadoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:EstadoProveedorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilProveedorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilProveedorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilProveedorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilProveedorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PerfilProveedorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PersonaJuridicaTipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PersonaJuridicaTipoEntidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PersonaJuridicaTipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PersonaJuridicaTipoEntidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PersonaJuridicaTipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PersonaJuridicaTipoEntidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PersonaJuridicaTipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PersonaJuridicaTipoEntidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PersonaJuridicaTipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:PersonaJuridicaTipoEntidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorJuridicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:ProveedorNaturalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:RepresentacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TituloProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TituloProveedorNaturalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TituloProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TituloProveedorNaturalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TituloProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TituloProveedorNaturalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TituloProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TituloProveedorNaturalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TituloProveedorNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/gestion_personas_crud/controllers:TituloProveedorNaturalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
