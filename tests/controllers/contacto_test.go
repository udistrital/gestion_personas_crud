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

// Pruebas del controlador ContactoController (/v1/contacto).

const (
	rutaContacto  = "/v1/contacto"
	tablaContacto = "contacto"
)

// cuerpoPostContacto retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostContacto(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoContacto(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestContactoPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaContacto, tablaContacto, cuerpoPostContacto(t))
	id2 := helpers.CrearPorPost(t, rutaContacto, tablaContacto, cuerpoPostContacto(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaContacto, id)

	// La llave foránea ProveedorId es obligatoria
	sinProveedorId := cuerpoPostContacto(t)
	sinProveedorId["ProveedorId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaContacto+"/", helpers.JSON(t, sinProveedorId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaContacto)
	t.Log("TestContactoPost Finalizado Correctamente (OK)")
}

func TestContactoGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearContacto(t)
	dato := helpers.ProbarGetOne(t, rutaContacto, creado.Id)

	// La llave foránea ProveedorId viene expandida
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaContacto)
	t.Log("TestContactoGetOne Finalizado Correctamente (OK)")
}

func TestContactoGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearContacto(t)
	dato := helpers.ProbarGetAll(t, rutaContacto, creado.Id)

	// La llave foránea ProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaContacto, fmt.Sprintf("query=ProveedorId.Id:%d", creado.ProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por ProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestContactoGetAll Finalizado Correctamente (OK)")
}

func TestContactoPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearContacto(t)
	helpers.ProbarPut(t, rutaContacto, tablaContacto, creado.Id, "NombreContacto", "Prueba-2", "nombre_contacto")
	t.Log("TestContactoPut Finalizado Correctamente (OK)")
}

func TestContactoDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearContacto(t)
	helpers.ProbarDelete(t, rutaContacto, tablaContacto, creado.Id)
	t.Log("TestContactoDelete Finalizado Correctamente (OK)")
}
