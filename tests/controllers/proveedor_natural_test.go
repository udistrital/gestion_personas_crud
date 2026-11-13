// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controllers

import (
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

// Pruebas del controlador ProveedorNaturalController (/v1/proveedor_natural).

const (
	rutaProveedorNatural  = "/v1/proveedor_natural"
	tablaProveedorNatural = "proveedor_natural"
)

// cuerpoPostProveedorNatural retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente.
func cuerpoPostProveedorNatural(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoProveedorNatural(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	return cuerpo
}

func TestProveedorNaturalPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id es obligatorio: es el Id de Proveedor (relación 1:1)
	sinId := cuerpoPostProveedorNatural(t)
	delete(sinId, "Id")
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaProveedorNatural+"/", helpers.JSON(t, sinId)), http.StatusBadRequest)
	idInexistente := cuerpoPostProveedorNatural(t)
	idInexistente["Id"] = helpers.IdInexistente
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaProveedorNatural+"/", helpers.JSON(t, idInexistente)), http.StatusBadRequest)

	cuerpo := cuerpoPostProveedorNatural(t)
	id := helpers.CrearPorPost(t, rutaProveedorNatural, tablaProveedorNatural, cuerpo)
	helpers.VerificarIgual(t, "Id igual al de Proveedor", float64(id), cuerpo["Id"])

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaProveedorNatural, id)

	helpers.ProbarPostInvalido(t, rutaProveedorNatural)
	t.Log("TestProveedorNaturalPost Finalizado Correctamente (OK)")
}

func TestProveedorNaturalGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedorNatural(t)
	helpers.ProbarGetOne(t, rutaProveedorNatural, creado.Id)

	helpers.ProbarGetOneInexistente(t, rutaProveedorNatural)
	t.Log("TestProveedorNaturalGetOne Finalizado Correctamente (OK)")
}

func TestProveedorNaturalGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedorNatural(t)
	helpers.ProbarGetAll(t, rutaProveedorNatural, creado.Id)
	t.Log("TestProveedorNaturalGetAll Finalizado Correctamente (OK)")
}

func TestProveedorNaturalPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedorNatural(t)
	helpers.ProbarPut(t, rutaProveedorNatural, tablaProveedorNatural, creado.Id, "Cargo", "Prueba-2", "cargo")
	t.Log("TestProveedorNaturalPut Finalizado Correctamente (OK)")
}

func TestProveedorNaturalDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedorNatural(t)
	helpers.ProbarDelete(t, rutaProveedorNatural, tablaProveedorNatural, creado.Id)
	t.Log("TestProveedorNaturalDelete Finalizado Correctamente (OK)")
}
