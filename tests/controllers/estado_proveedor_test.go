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

// Pruebas del controlador EstadoProveedorController (/v1/estado_proveedor).

const (
	rutaEstadoProveedor  = "/v1/estado_proveedor"
	tablaEstadoProveedor = "estado_proveedor"
)

// cuerpoPostEstadoProveedor retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostEstadoProveedor(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoEstadoProveedor(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestEstadoProveedorPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaEstadoProveedor, tablaEstadoProveedor, cuerpoPostEstadoProveedor(t))
	id2 := helpers.CrearPorPost(t, rutaEstadoProveedor, tablaEstadoProveedor, cuerpoPostEstadoProveedor(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaEstadoProveedor, id)

	// La llave foránea ProveedorId es obligatoria
	sinProveedorId := cuerpoPostEstadoProveedor(t)
	sinProveedorId["ProveedorId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaEstadoProveedor+"/", helpers.JSON(t, sinProveedorId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaEstadoProveedor)
	t.Log("TestEstadoProveedorPost Finalizado Correctamente (OK)")
}

func TestEstadoProveedorGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearEstadoProveedor(t)
	dato := helpers.ProbarGetOne(t, rutaEstadoProveedor, creado.Id)

	// La llave foránea ProveedorId viene expandida
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaEstadoProveedor)
	t.Log("TestEstadoProveedorGetOne Finalizado Correctamente (OK)")
}

func TestEstadoProveedorGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearEstadoProveedor(t)
	dato := helpers.ProbarGetAll(t, rutaEstadoProveedor, creado.Id)

	// La llave foránea ProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaEstadoProveedor, fmt.Sprintf("query=ProveedorId.Id:%d", creado.ProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por ProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestEstadoProveedorGetAll Finalizado Correctamente (OK)")
}

func TestEstadoProveedorPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearEstadoProveedor(t)
	helpers.ProbarPut(t, rutaEstadoProveedor, tablaEstadoProveedor, creado.Id, "Observacion", "Prueba-2", "observacion")
	t.Log("TestEstadoProveedorPut Finalizado Correctamente (OK)")
}

func TestEstadoProveedorDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearEstadoProveedor(t)
	helpers.ProbarDelete(t, rutaEstadoProveedor, tablaEstadoProveedor, creado.Id)
	t.Log("TestEstadoProveedorDelete Finalizado Correctamente (OK)")
}
