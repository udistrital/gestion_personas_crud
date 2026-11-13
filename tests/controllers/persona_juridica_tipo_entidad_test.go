// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/models"
	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

// Pruebas del controlador PersonaJuridicaTipoEntidadController (/v1/persona_juridica_tipo_entidad).

const (
	rutaPersonaJuridicaTipoEntidad  = "/v1/persona_juridica_tipo_entidad"
	tablaPersonaJuridicaTipoEntidad = "persona_juridica_tipo_entidad"
)

// cuerpoPostPersonaJuridicaTipoEntidad retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostPersonaJuridicaTipoEntidad(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoPersonaJuridicaTipoEntidad(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestPersonaJuridicaTipoEntidadPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaPersonaJuridicaTipoEntidad, tablaPersonaJuridicaTipoEntidad, cuerpoPostPersonaJuridicaTipoEntidad(t))
	id2 := helpers.CrearPorPost(t, rutaPersonaJuridicaTipoEntidad, tablaPersonaJuridicaTipoEntidad, cuerpoPostPersonaJuridicaTipoEntidad(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaPersonaJuridicaTipoEntidad, id)

	// La llave foránea PersonaJuridicaId es obligatoria
	sinPersonaJuridicaId := cuerpoPostPersonaJuridicaTipoEntidad(t)
	sinPersonaJuridicaId["PersonaJuridicaId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaPersonaJuridicaTipoEntidad+"/", helpers.JSON(t, sinPersonaJuridicaId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaPersonaJuridicaTipoEntidad)
	t.Log("TestPersonaJuridicaTipoEntidadPost Finalizado Correctamente (OK)")
}

func TestPersonaJuridicaTipoEntidadGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearPersonaJuridicaTipoEntidad(t)
	dato := helpers.ProbarGetOne(t, rutaPersonaJuridicaTipoEntidad, creado.Id)

	// La llave foránea PersonaJuridicaId viene expandida
	if relacionado, err := models.GetProveedorJuridicoById(creado.PersonaJuridicaId.Id); err != nil {
		t.Fatal("No se pudo consultar ProveedorJuridico:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "PersonaJuridicaId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaPersonaJuridicaTipoEntidad)
	t.Log("TestPersonaJuridicaTipoEntidadGetOne Finalizado Correctamente (OK)")
}

func TestPersonaJuridicaTipoEntidadGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearPersonaJuridicaTipoEntidad(t)
	dato := helpers.ProbarGetAll(t, rutaPersonaJuridicaTipoEntidad, creado.Id)

	// La llave foránea PersonaJuridicaId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetProveedorJuridicoById(creado.PersonaJuridicaId.Id); err != nil {
		t.Fatal("No se pudo consultar ProveedorJuridico:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "PersonaJuridicaId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaPersonaJuridicaTipoEntidad, fmt.Sprintf("query=PersonaJuridicaId.Id:%d", creado.PersonaJuridicaId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por PersonaJuridicaId.Id y se obtuvo", l)
	}
	t.Log("TestPersonaJuridicaTipoEntidadGetAll Finalizado Correctamente (OK)")
}

func TestPersonaJuridicaTipoEntidadPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearPersonaJuridicaTipoEntidad(t)
	helpers.ProbarPut(t, rutaPersonaJuridicaTipoEntidad, tablaPersonaJuridicaTipoEntidad, creado.Id, "NombrePerfilEntidad", "Prueba-2", "nombre_perfil_entidad")
	t.Log("TestPersonaJuridicaTipoEntidadPut Finalizado Correctamente (OK)")
}

func TestPersonaJuridicaTipoEntidadDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearPersonaJuridicaTipoEntidad(t)
	helpers.ProbarDelete(t, rutaPersonaJuridicaTipoEntidad, tablaPersonaJuridicaTipoEntidad, creado.Id)
	t.Log("TestPersonaJuridicaTipoEntidadDelete Finalizado Correctamente (OK)")
}
