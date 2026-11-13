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

// Pruebas del controlador TituloProveedorNaturalController (/v1/titulo_proveedor_natural).

const (
	rutaTituloProveedorNatural  = "/v1/titulo_proveedor_natural"
	tablaTituloProveedorNatural = "titulo_proveedor_natural"
)

// cuerpoPostTituloProveedorNatural retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostTituloProveedorNatural(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoTituloProveedorNatural(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestTituloProveedorNaturalPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaTituloProveedorNatural, tablaTituloProveedorNatural, cuerpoPostTituloProveedorNatural(t))
	id2 := helpers.CrearPorPost(t, rutaTituloProveedorNatural, tablaTituloProveedorNatural, cuerpoPostTituloProveedorNatural(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaTituloProveedorNatural, id)

	// La llave foránea ProveedorId es obligatoria
	sinProveedorId := cuerpoPostTituloProveedorNatural(t)
	sinProveedorId["ProveedorId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaTituloProveedorNatural+"/", helpers.JSON(t, sinProveedorId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaTituloProveedorNatural)
	t.Log("TestTituloProveedorNaturalPost Finalizado Correctamente (OK)")
}

func TestTituloProveedorNaturalGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearTituloProveedorNatural(t)
	dato := helpers.ProbarGetOne(t, rutaTituloProveedorNatural, creado.Id)

	// La llave foránea ProveedorId viene expandida
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaTituloProveedorNatural)
	t.Log("TestTituloProveedorNaturalGetOne Finalizado Correctamente (OK)")
}

func TestTituloProveedorNaturalGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearTituloProveedorNatural(t)
	dato := helpers.ProbarGetAll(t, rutaTituloProveedorNatural, creado.Id)

	// La llave foránea ProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaTituloProveedorNatural, fmt.Sprintf("query=ProveedorId.Id:%d", creado.ProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por ProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestTituloProveedorNaturalGetAll Finalizado Correctamente (OK)")
}

func TestTituloProveedorNaturalPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearTituloProveedorNatural(t)
	helpers.ProbarPut(t, rutaTituloProveedorNatural, tablaTituloProveedorNatural, creado.Id, "TituloObtenido", "Prueba-2", "titulo_obtenido")
	t.Log("TestTituloProveedorNaturalPut Finalizado Correctamente (OK)")
}

func TestTituloProveedorNaturalDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearTituloProveedorNatural(t)
	helpers.ProbarDelete(t, rutaTituloProveedorNatural, tablaTituloProveedorNatural, creado.Id)
	t.Log("TestTituloProveedorNaturalDelete Finalizado Correctamente (OK)")
}
