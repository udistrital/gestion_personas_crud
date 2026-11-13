// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controllers

import (
	"testing"

	"github.com/udistrital/gestion_personas_crud/tests/helpers"
)

// Pruebas del controlador ProveedorController (/v1/proveedor).

const (
	rutaProveedor  = "/v1/proveedor"
	tablaProveedor = "proveedor"
)

// cuerpoPostProveedor retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostProveedor(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoProveedor(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestProveedorPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaProveedor, tablaProveedor, cuerpoPostProveedor(t))
	id2 := helpers.CrearPorPost(t, rutaProveedor, tablaProveedor, cuerpoPostProveedor(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaProveedor, id)

	helpers.ProbarPostInvalido(t, rutaProveedor)
	t.Log("TestProveedorPost Finalizado Correctamente (OK)")
}

func TestProveedorGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedor(t)
	helpers.ProbarGetOne(t, rutaProveedor, creado.Id)

	helpers.ProbarGetOneInexistente(t, rutaProveedor)
	t.Log("TestProveedorGetOne Finalizado Correctamente (OK)")
}

func TestProveedorGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedor(t)
	helpers.ProbarGetAll(t, rutaProveedor, creado.Id)
	t.Log("TestProveedorGetAll Finalizado Correctamente (OK)")
}

func TestProveedorPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedor(t)
	helpers.ProbarPut(t, rutaProveedor, tablaProveedor, creado.Id, "DescripcionPortafolio", "Prueba-2", "descripcion_portafolio")
	t.Log("TestProveedorPut Finalizado Correctamente (OK)")
}

func TestProveedorDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearProveedor(t)
	helpers.ProbarDelete(t, rutaProveedor, tablaProveedor, creado.Id)
	t.Log("TestProveedorDelete Finalizado Correctamente (OK)")
}
