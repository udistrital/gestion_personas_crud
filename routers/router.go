// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/gestion_personas_crud/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/contacto",
			beego.NSInclude(
				&controllers.ContactoController{},
			),
		),

		beego.NSNamespace("/cuenta_bancaria",
			beego.NSInclude(
				&controllers.CuentaBancariaController{},
			),
		),

		beego.NSNamespace("/tipo_conformacion",
			beego.NSInclude(
				&controllers.TipoConformacionController{},
			),
		),

		beego.NSNamespace("/proveedor_juridico",
			beego.NSInclude(
				&controllers.ProveedorJuridicoController{},
			),
		),

		beego.NSNamespace("/proveedor",
			beego.NSInclude(
				&controllers.ProveedorController{},
			),
		),

		beego.NSNamespace("/actividad_economica_proveedor",
			beego.NSInclude(
				&controllers.ActividadEconomicaProveedorController{},
			),
		),

		beego.NSNamespace("/representacion",
			beego.NSInclude(
				&controllers.RepresentacionController{},
			),
		),

		beego.NSNamespace("/persona_juridica_tipo_entidad",
			beego.NSInclude(
				&controllers.PersonaJuridicaTipoEntidadController{},
			),
		),

		beego.NSNamespace("/proveedor_natural",
			beego.NSInclude(
				&controllers.ProveedorNaturalController{},
			),
		),

		beego.NSNamespace("/perfil_proveedor",
			beego.NSInclude(
				&controllers.PerfilProveedorController{},
			),
		),

		beego.NSNamespace("/documento",
			beego.NSInclude(
				&controllers.DocumentoController{},
			),
		),

		beego.NSNamespace("/titulo_proveedor_natural",
			beego.NSInclude(
				&controllers.TituloProveedorNaturalController{},
			),
		),

		beego.NSNamespace("/estado_proveedor",
			beego.NSInclude(
				&controllers.EstadoProveedorController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
