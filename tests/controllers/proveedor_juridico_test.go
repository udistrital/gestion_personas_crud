// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controllers

import (
	"net/http"
	"testing"

	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

// Pruebas del controlador ProveedorJuridicoController (/v1/proveedor_juridico).

const (
	rutaProveedorJuridico  = "/v1/proveedor_juridico"
	tablaProveedorJuridico = "proveedor_juridico"
)

// cuerpoPostProveedorJuridico retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente.
func cuerpoPostProveedorJuridico(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoProveedorJuridico(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	return cuerpo
}

func TestProveedorJuridicoPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id es obligatorio: es el Id de Proveedor (relación 1:1)
	sinId := cuerpoPostProveedorJuridico(t)
	delete(sinId, "Id")
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaProveedorJuridico+"/", helpers.JSON(t, sinId)), http.StatusBadRequest)
	idInexistente := cuerpoPostProveedorJuridico(t)
	idInexistente["Id"] = helpers.IdInexistente
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaProveedorJuridico+"/", helpers.JSON(t, idInexistente)), http.StatusBadRequest)

	cuerpo := cuerpoPostProveedorJuridico(t)
	id := helpers.CrearPorPost(t, rutaProveedorJuridico, tablaProveedorJuridico, cuerpo)
	helpers.VerificarIgual(t, "Id igual al de Proveedor", float64(id), cuerpo["Id"])

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaProveedorJuridico, id)

	helpers.ProbarPostInvalido(t, rutaProveedorJuridico)
	t.Log("TestProveedorJuridicoPost Finalizado Correctamente (OK)")
}

func TestProveedorJuridicoGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedorJuridico(t)
	helpers.ProbarGetOne(t, rutaProveedorJuridico, creado.Id)

	helpers.ProbarGetOneInexistente(t, rutaProveedorJuridico)
	t.Log("TestProveedorJuridicoGetOne Finalizado Correctamente (OK)")
}

func TestProveedorJuridicoGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedorJuridico(t)
	helpers.ProbarGetAll(t, rutaProveedorJuridico, creado.Id)
	t.Log("TestProveedorJuridicoGetAll Finalizado Correctamente (OK)")
}

func TestProveedorJuridicoPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedorJuridico(t)
	helpers.ProbarPut(t, rutaProveedorJuridico, tablaProveedorJuridico, creado.Id, "CodigoPostal", "Prueba-2", "codigo_postal")
	t.Log("TestProveedorJuridicoPut Finalizado Correctamente (OK)")
}

func TestProveedorJuridicoDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedorJuridico(t)
	helpers.ProbarDelete(t, rutaProveedorJuridico, tablaProveedorJuridico, creado.Id)
	t.Log("TestProveedorJuridicoDelete Finalizado Correctamente (OK)")
}
