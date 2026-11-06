package controllers

import "github.com/udistrital/gestion_personas_crud/models"

// Tipos que documentan en Swagger la respuesta de GetAll de cada entidad: los campos
// comunes de APIResponseLista y Data como lista del modelo correspondiente. Solo se
// usan en las anotaciones @Success; la respuesta la construye writeSuccessLista.

// ActividadEconomicaProveedorLista es la respuesta de GetAll de ActividadEconomicaProveedor
type ActividadEconomicaProveedorLista struct {
	APIResponseLista
	Data []models.ActividadEconomicaProveedor `json:"Data"`
}

// ContactoLista es la respuesta de GetAll de Contacto
type ContactoLista struct {
	APIResponseLista
	Data []models.Contacto `json:"Data"`
}

// CuentaBancariaLista es la respuesta de GetAll de CuentaBancaria
type CuentaBancariaLista struct {
	APIResponseLista
	Data []models.CuentaBancaria `json:"Data"`
}

// DeclaracionLista es la respuesta de GetAll de Declaracion
type DeclaracionLista struct {
	APIResponseLista
	Data []models.Declaracion `json:"Data"`
}

// DocumentoLista es la respuesta de GetAll de Documento
type DocumentoLista struct {
	APIResponseLista
	Data []models.Documento `json:"Data"`
}

// InformacionFinancieraLista es la respuesta de GetAll de InformacionFinanciera
type InformacionFinancieraLista struct {
	APIResponseLista
	Data []models.InformacionFinanciera `json:"Data"`
}

// PerfilLista es la respuesta de GetAll de Perfil
type PerfilLista struct {
	APIResponseLista
	Data []models.Perfil `json:"Data"`
}

// PerfilFinancieroLista es la respuesta de GetAll de PerfilFinanciero
type PerfilFinancieroLista struct {
	APIResponseLista
	Data []models.PerfilFinanciero `json:"Data"`
}

// PrefijoFacturacionLista es la respuesta de GetAll de PrefijoFacturacion
type PrefijoFacturacionLista struct {
	APIResponseLista
	Data []models.PrefijoFacturacion `json:"Data"`
}

// ProveedorLista es la respuesta de GetAll de Proveedor
type ProveedorLista struct {
	APIResponseLista
	Data []models.Proveedor `json:"Data"`
}

// ProveedorJuridicoLista es la respuesta de GetAll de ProveedorJuridico
type ProveedorJuridicoLista struct {
	APIResponseLista
	Data []models.ProveedorJuridico `json:"Data"`
}

// ProveedorNaturalLista es la respuesta de GetAll de ProveedorNatural
type ProveedorNaturalLista struct {
	APIResponseLista
	Data []models.ProveedorNatural `json:"Data"`
}

// RangoFacturacionLista es la respuesta de GetAll de RangoFacturacion
type RangoFacturacionLista struct {
	APIResponseLista
	Data []models.RangoFacturacion `json:"Data"`
}

// RepresentacionLista es la respuesta de GetAll de Representacion
type RepresentacionLista struct {
	APIResponseLista
	Data []models.Representacion `json:"Data"`
}

// ResponsabilidadFiscalPerfilLista es la respuesta de GetAll de ResponsabilidadFiscalPerfil
type ResponsabilidadFiscalPerfilLista struct {
	APIResponseLista
	Data []models.ResponsabilidadFiscalPerfil `json:"Data"`
}

// TipoDeclaracionProveedorLista es la respuesta de GetAll de TipoDeclaracionProveedor
type TipoDeclaracionProveedorLista struct {
	APIResponseLista
	Data []models.TipoDeclaracionProveedor `json:"Data"`
}
