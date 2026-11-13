// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package helpers

import (
	"testing"
	"time"

	"github.com/udistrital/gestion_personas_crud/models"
)

// Las funciones Nuevo<Modelo> retornan un registro de ejemplo válido sin
// guardar (creando en la base de datos sus llaves foráneas). Las funciones
// Crear<Modelo> lo guardan directamente con el modelo y programan su
// eliminación al terminar la prueba.

func NuevoActividadEconomicaProveedor(t *testing.T) *models.ActividadEconomicaProveedor {
	t.Helper()
	return &models.ActividadEconomicaProveedor{
		ProveedorId:           &models.Proveedor{Id: CrearProveedor(t).Id},
		ActividadEconomicaId:  1,
		UsuarioModificacionId: 1,
	}
}

func CrearActividadEconomicaProveedor(t *testing.T) *models.ActividadEconomicaProveedor {
	t.Helper()
	v := NuevoActividadEconomicaProveedor(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddActividadEconomicaProveedor(v); err != nil {
		t.Fatal("No se pudo crear ActividadEconomicaProveedor:", err)
	}
	EliminarAlFinal(t, "actividad_economica_proveedor", v.Id)
	return v
}

func NuevoContacto(t *testing.T) *models.Contacto {
	t.Helper()
	return &models.Contacto{
		NombreContacto:        "Prueba",
		ValorEspecifico:       "Prueba",
		ProveedorId:           &models.Proveedor{Id: CrearProveedor(t).Id},
		TipoCanalContactoId:   1,
		Proposito:             "Prueba",
		Extension:             "Prueba",
		UsuarioModificacionId: 1,
	}
}

func CrearContacto(t *testing.T) *models.Contacto {
	t.Helper()
	v := NuevoContacto(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddContacto(v); err != nil {
		t.Fatal("No se pudo crear Contacto:", err)
	}
	EliminarAlFinal(t, "contacto", v.Id)
	return v
}

func NuevoCuentaBancaria(t *testing.T) *models.CuentaBancaria {
	t.Helper()
	return &models.CuentaBancaria{
		TipoCuentaId:           1,
		NumeroCuenta:           "Prueba",
		EntidadBancariaId:      1,
		ContactoId:             &models.Contacto{Id: CrearContacto(t).Id},
		TipoCuentaBancariaId:   true,
		MontoCapitalAutorizado: 1.5,
		ProveedorId:            &models.Proveedor{Id: CrearProveedor(t).Id},
		UsuarioModificacionId:  1,
	}
}

func CrearCuentaBancaria(t *testing.T) *models.CuentaBancaria {
	t.Helper()
	v := NuevoCuentaBancaria(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddCuentaBancaria(v); err != nil {
		t.Fatal("No se pudo crear CuentaBancaria:", err)
	}
	EliminarAlFinal(t, "cuenta_bancaria", v.Id)
	return v
}

func NuevoDocumento(t *testing.T) *models.Documento {
	t.Helper()
	return &models.Documento{
		TipoDocumentoId:       1,
		Enlace:                "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
		ProveedorId:           &models.Proveedor{Id: CrearProveedor(t).Id},
		UsuarioModificacionId: 1,
	}
}

func CrearDocumento(t *testing.T) *models.Documento {
	t.Helper()
	v := NuevoDocumento(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddDocumento(v); err != nil {
		t.Fatal("No se pudo crear Documento:", err)
	}
	EliminarAlFinal(t, "documento", v.Id)
	return v
}

func NuevoEstadoProveedor(t *testing.T) *models.EstadoProveedor {
	t.Helper()
	return &models.EstadoProveedor{
		EstadoId:              1,
		ProveedorId:           &models.Proveedor{Id: CrearProveedor(t).Id},
		Observacion:           "Prueba",
		UsuarioModificacionId: 1,
	}
}

func CrearEstadoProveedor(t *testing.T) *models.EstadoProveedor {
	t.Helper()
	v := NuevoEstadoProveedor(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddEstadoProveedor(v); err != nil {
		t.Fatal("No se pudo crear EstadoProveedor:", err)
	}
	EliminarAlFinal(t, "estado_proveedor", v.Id)
	return v
}

func NuevoPerfilProveedor(t *testing.T) *models.PerfilProveedor {
	t.Helper()
	return &models.PerfilProveedor{
		PerfilId:              1,
		ProveedorId:           &models.Proveedor{Id: CrearProveedor(t).Id},
		UsuarioModificacionId: 1,
	}
}

func CrearPerfilProveedor(t *testing.T) *models.PerfilProveedor {
	t.Helper()
	v := NuevoPerfilProveedor(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddPerfilProveedor(v); err != nil {
		t.Fatal("No se pudo crear PerfilProveedor:", err)
	}
	EliminarAlFinal(t, "perfil_proveedor", v.Id)
	return v
}

func NuevoPersonaJuridicaTipoEntidad(t *testing.T) *models.PersonaJuridicaTipoEntidad {
	t.Helper()
	return &models.PersonaJuridicaTipoEntidad{
		NombrePerfilEntidad:   "Prueba",
		TipoEntidadId:         1,
		PersonaJuridicaId:     &models.ProveedorJuridico{Id: CrearProveedorJuridico(t).Id},
		UsuarioModificacionId: 1,
	}
}

func CrearPersonaJuridicaTipoEntidad(t *testing.T) *models.PersonaJuridicaTipoEntidad {
	t.Helper()
	v := NuevoPersonaJuridicaTipoEntidad(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddPersonaJuridicaTipoEntidad(v); err != nil {
		t.Fatal("No se pudo crear PersonaJuridicaTipoEntidad:", err)
	}
	EliminarAlFinal(t, "persona_juridica_tipo_entidad", v.Id)
	return v
}

func NuevoProveedor(t *testing.T) *models.Proveedor {
	t.Helper()
	return &models.Proveedor{
		TipoProveedor:         1,
		DescripcionPortafolio: "Prueba",
		Consentimiento:        true,
		TerceroId:             1,
		UsuarioModificacionId: 1,
	}
}

func CrearProveedor(t *testing.T) *models.Proveedor {
	t.Helper()
	v := NuevoProveedor(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddProveedor(v); err != nil {
		t.Fatal("No se pudo crear Proveedor:", err)
	}
	EliminarAlFinal(t, "proveedor", v.Id)
	return v
}

func NuevoProveedorJuridico(t *testing.T) *models.ProveedorJuridico {
	t.Helper()
	return &models.ProveedorJuridico{
		Id:                    CrearProveedor(t).Id, // relación 1:1: comparte el Id de Proveedor
		ProcedenciaEmpresaId:  1,
		CiudadOrigenId:        1,
		PaisOrigenId:          1,
		CodigoPostal:          "Prueba",
		TipoConformacionId:    1,
		ExclusividadProducto:  true,
		RegimenContributivo:   1,
		Pyme:                  true,
		RegistroMercantil:     true,
		SujetoRetencion:       true,
		AgenteRetenedor:       true,
		ResponsableIca:        true,
		ResponsableIva:        true,
		UsuarioModificacionId: 1,
	}
}

func CrearProveedorJuridico(t *testing.T) *models.ProveedorJuridico {
	t.Helper()
	v := NuevoProveedorJuridico(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddProveedorJuridico(v); err != nil {
		t.Fatal("No se pudo crear ProveedorJuridico:", err)
	}
	EliminarAlFinal(t, "proveedor_juridico", v.Id)
	return v
}

func NuevoProveedorNatural(t *testing.T) *models.ProveedorNatural {
	t.Helper()
	return &models.ProveedorNatural{
		Id:                                CrearProveedor(t).Id, // relación 1:1: comparte el Id de Proveedor
		PaisNacimientoId:                  1,
		Cargo:                             "Prueba",
		MedicinaPrepagada:                 true,
		ValorUvtPrepagada:                 1.5,
		InteresViviendaAfc:                1.5,
		DependienteHijoMejorEdad:          true,
		DependienteHijoMenos23Estudiando:  true,
		DependienteHijoMas23Discapacitado: true,
		DependienteConyuge:                true,
		DependientePadreOHermano:          true,
		Pensionado:                        true,
		ValorAfc:                          1.5,
		ResponsableIva:                    true,
		ResideColombia:                    true,
		PensionVoluntaria:                 true,
		ValorPensionVoluntaria:            1.5,
		ExperienciaLaboral:                1,
		ExperienciaProfesional:            1,
		UsuarioModificacionId:             1,
	}
}

func CrearProveedorNatural(t *testing.T) *models.ProveedorNatural {
	t.Helper()
	v := NuevoProveedorNatural(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddProveedorNatural(v); err != nil {
		t.Fatal("No se pudo crear ProveedorNatural:", err)
	}
	EliminarAlFinal(t, "proveedor_natural", v.Id)
	return v
}

func NuevoRepresentacion(t *testing.T) *models.Representacion {
	t.Helper()
	return &models.Representacion{
		ProveedorId:           &models.Proveedor{Id: CrearProveedor(t).Id},
		RepresentanteId:       1,
		UsuarioModificacionId: 1,
	}
}

func CrearRepresentacion(t *testing.T) *models.Representacion {
	t.Helper()
	v := NuevoRepresentacion(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddRepresentacion(v); err != nil {
		t.Fatal("No se pudo crear Representacion:", err)
	}
	EliminarAlFinal(t, "representacion", v.Id)
	return v
}

func NuevoTituloProveedorNatural(t *testing.T) *models.TituloProveedorNatural {
	t.Helper()
	return &models.TituloProveedorNatural{
		TituloObtenido:                  "Prueba",
		NucleoBasicoConocimientoSniesId: 1,
		ProveedorId:                     &models.Proveedor{Id: CrearProveedor(t).Id},
		PerfilId:                        1,
		UsuarioModificacionId:           1,
	}
}

func CrearTituloProveedorNatural(t *testing.T) *models.TituloProveedorNatural {
	t.Helper()
	v := NuevoTituloProveedorNatural(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddTituloProveedorNatural(v); err != nil {
		t.Fatal("No se pudo crear TituloProveedorNatural:", err)
	}
	EliminarAlFinal(t, "titulo_proveedor_natural", v.Id)
	return v
}
