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

// Pruebas del controlador DocumentoController (/v1/documento).

const (
	rutaDocumento  = "/v1/documento"
	tablaDocumento = "documento"
)

// cuerpoPostDocumento retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostDocumento(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoDocumento(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestDocumentoPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaDocumento, tablaDocumento, cuerpoPostDocumento(t))
	id2 := helpers.CrearPorPost(t, rutaDocumento, tablaDocumento, cuerpoPostDocumento(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaDocumento, id)

	// La llave foránea ProveedorId es obligatoria
	sinProveedorId := cuerpoPostDocumento(t)
	sinProveedorId["ProveedorId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaDocumento+"/", helpers.JSON(t, sinProveedorId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaDocumento)
	t.Log("TestDocumentoPost Finalizado Correctamente (OK)")
}

func TestDocumentoGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearDocumento(t)
	dato := helpers.ProbarGetOne(t, rutaDocumento, creado.Id)

	// La llave foránea ProveedorId viene expandida
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaDocumento)
	t.Log("TestDocumentoGetOne Finalizado Correctamente (OK)")
}

func TestDocumentoGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearDocumento(t)
	dato := helpers.ProbarGetAll(t, rutaDocumento, creado.Id)

	// La llave foránea ProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaDocumento, fmt.Sprintf("query=ProveedorId.Id:%d", creado.ProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por ProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestDocumentoGetAll Finalizado Correctamente (OK)")
}

func TestDocumentoPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearDocumento(t)
	helpers.ProbarPut(t, rutaDocumento, tablaDocumento, creado.Id, "TipoDocumentoId", 2, "tipo_documento_id")
	t.Log("TestDocumentoPut Finalizado Correctamente (OK)")
}

func TestDocumentoDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearDocumento(t)
	helpers.ProbarDelete(t, rutaDocumento, tablaDocumento, creado.Id)
	t.Log("TestDocumentoDelete Finalizado Correctamente (OK)")
}
