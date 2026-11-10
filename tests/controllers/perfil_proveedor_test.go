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

// Pruebas del controlador PerfilProveedorController (/v1/perfil_proveedor).

const (
	rutaPerfilProveedor  = "/v1/perfil_proveedor"
	tablaPerfilProveedor = "perfil_proveedor"
)

// cuerpoPostPerfilProveedor retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostPerfilProveedor(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoPerfilProveedor(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestPerfilProveedorPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaPerfilProveedor, tablaPerfilProveedor, cuerpoPostPerfilProveedor(t))
	id2 := helpers.CrearPorPost(t, rutaPerfilProveedor, tablaPerfilProveedor, cuerpoPostPerfilProveedor(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaPerfilProveedor, id)

	// La llave foránea ProveedorId es obligatoria
	sinProveedorId := cuerpoPostPerfilProveedor(t)
	sinProveedorId["ProveedorId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaPerfilProveedor+"/", helpers.JSON(t, sinProveedorId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaPerfilProveedor)
	t.Log("TestPerfilProveedorPost Finalizado Correctamente (OK)")
}

func TestPerfilProveedorGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearPerfilProveedor(t)
	dato := helpers.ProbarGetOne(t, rutaPerfilProveedor, creado.Id)

	// La llave foránea ProveedorId viene expandida
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaPerfilProveedor)
	t.Log("TestPerfilProveedorGetOne Finalizado Correctamente (OK)")
}

func TestPerfilProveedorGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearPerfilProveedor(t)
	dato := helpers.ProbarGetAll(t, rutaPerfilProveedor, creado.Id)

	// La llave foránea ProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaPerfilProveedor, fmt.Sprintf("query=ProveedorId.Id:%d", creado.ProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por ProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestPerfilProveedorGetAll Finalizado Correctamente (OK)")
}

func TestPerfilProveedorPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearPerfilProveedor(t)
	helpers.ProbarPut(t, rutaPerfilProveedor, tablaPerfilProveedor, creado.Id, "PerfilId", 2, "perfil_id")
	t.Log("TestPerfilProveedorPut Finalizado Correctamente (OK)")
}

func TestPerfilProveedorDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearPerfilProveedor(t)
	helpers.ProbarDelete(t, rutaPerfilProveedor, tablaPerfilProveedor, creado.Id)
	t.Log("TestPerfilProveedorDelete Finalizado Correctamente (OK)")
}
