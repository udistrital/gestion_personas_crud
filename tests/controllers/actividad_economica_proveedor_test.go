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

// Pruebas del controlador ActividadEconomicaProveedorController (/v1/actividad_economica_proveedor).

const (
	rutaActividadEconomicaProveedor  = "/v1/actividad_economica_proveedor"
	tablaActividadEconomicaProveedor = "actividad_economica_proveedor"
)

// cuerpoPostActividadEconomicaProveedor retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostActividadEconomicaProveedor(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoActividadEconomicaProveedor(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestActividadEconomicaProveedorPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaActividadEconomicaProveedor, tablaActividadEconomicaProveedor, cuerpoPostActividadEconomicaProveedor(t))
	id2 := helpers.CrearPorPost(t, rutaActividadEconomicaProveedor, tablaActividadEconomicaProveedor, cuerpoPostActividadEconomicaProveedor(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaActividadEconomicaProveedor, id)

	// La llave foránea ProveedorId es obligatoria
	sinProveedorId := cuerpoPostActividadEconomicaProveedor(t)
	sinProveedorId["ProveedorId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaActividadEconomicaProveedor+"/", helpers.JSON(t, sinProveedorId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaActividadEconomicaProveedor)
	t.Log("TestActividadEconomicaProveedorPost Finalizado Correctamente (OK)")
}

func TestActividadEconomicaProveedorGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearActividadEconomicaProveedor(t)
	dato := helpers.ProbarGetOne(t, rutaActividadEconomicaProveedor, creado.Id)

	// La llave foránea ProveedorId viene expandida
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaActividadEconomicaProveedor)
	t.Log("TestActividadEconomicaProveedorGetOne Finalizado Correctamente (OK)")
}

func TestActividadEconomicaProveedorGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearActividadEconomicaProveedor(t)
	dato := helpers.ProbarGetAll(t, rutaActividadEconomicaProveedor, creado.Id)

	// La llave foránea ProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaActividadEconomicaProveedor, fmt.Sprintf("query=ProveedorId.Id:%d", creado.ProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por ProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestActividadEconomicaProveedorGetAll Finalizado Correctamente (OK)")
}

func TestActividadEconomicaProveedorPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearActividadEconomicaProveedor(t)
	helpers.ProbarPut(t, rutaActividadEconomicaProveedor, tablaActividadEconomicaProveedor, creado.Id, "ActividadEconomicaId", 2, "actividad_economica_id")
	t.Log("TestActividadEconomicaProveedorPut Finalizado Correctamente (OK)")
}

func TestActividadEconomicaProveedorDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearActividadEconomicaProveedor(t)
	helpers.ProbarDelete(t, rutaActividadEconomicaProveedor, tablaActividadEconomicaProveedor, creado.Id)
	t.Log("TestActividadEconomicaProveedorDelete Finalizado Correctamente (OK)")
}
