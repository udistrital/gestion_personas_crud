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

// Pruebas del controlador RepresentacionController (/v1/representacion).

const (
	rutaRepresentacion  = "/v1/representacion"
	tablaRepresentacion = "representacion"
)

// cuerpoPostRepresentacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostRepresentacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoRepresentacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestRepresentacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaRepresentacion, tablaRepresentacion, cuerpoPostRepresentacion(t))
	id2 := helpers.CrearPorPost(t, rutaRepresentacion, tablaRepresentacion, cuerpoPostRepresentacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaRepresentacion, id)

	// La llave foránea ProveedorId es obligatoria
	sinProveedorId := cuerpoPostRepresentacion(t)
	sinProveedorId["ProveedorId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaRepresentacion+"/", helpers.JSON(t, sinProveedorId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaRepresentacion)
	t.Log("TestRepresentacionPost Finalizado Correctamente (OK)")
}

func TestRepresentacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRepresentacion(t)
	dato := helpers.ProbarGetOne(t, rutaRepresentacion, creado.Id)

	// La llave foránea ProveedorId viene expandida
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaRepresentacion)
	t.Log("TestRepresentacionGetOne Finalizado Correctamente (OK)")
}

func TestRepresentacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRepresentacion(t)
	dato := helpers.ProbarGetAll(t, rutaRepresentacion, creado.Id)

	// La llave foránea ProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaRepresentacion, fmt.Sprintf("query=ProveedorId.Id:%d", creado.ProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por ProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestRepresentacionGetAll Finalizado Correctamente (OK)")
}

func TestRepresentacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRepresentacion(t)
	helpers.ProbarPut(t, rutaRepresentacion, tablaRepresentacion, creado.Id, "RepresentanteId", 2, "representante_id")
	t.Log("TestRepresentacionPut Finalizado Correctamente (OK)")
}

func TestRepresentacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRepresentacion(t)
	helpers.ProbarDelete(t, rutaRepresentacion, tablaRepresentacion, creado.Id)
	t.Log("TestRepresentacionDelete Finalizado Correctamente (OK)")
}
