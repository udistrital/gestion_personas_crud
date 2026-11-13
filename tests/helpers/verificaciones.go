package helpers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

// FechaCliente es una fecha que el cliente envía en los campos administrados
// por el servicio (FechaCreacion/FechaModificacion) y que debe ignorarse.
const FechaCliente = "2000-01-01T00:00:00Z"

// CrearPorPost crea un registro mediante el servicio POST, verifica la
// respuesta 201 y programa su eliminación al terminar la prueba.
func CrearPorPost(t *testing.T, ruta, tabla string, cuerpo interface{}) int {
	t.Helper()
	w := Peticion(http.MethodPost, ruta+"/", JSON(t, cuerpo))
	id := VerificarExito(t, w, http.StatusCreated).IdDato(t)
	EliminarAlFinal(t, tabla, id)
	return id
}

// VerificarCamposAdministrados comprueba que, al crear, el servicio asigna
// Activo=true y las fechas de creación y modificación, ignorando los valores
// enviados por el cliente.
func VerificarCamposAdministrados(t *testing.T, tabla string, id int) {
	t.Helper()
	fila := Fila(t, tabla, id)
	if fila["activo"] != true {
		t.Error("Se esperaba activo=true al crear (aunque el cliente envíe false) y se obtuvo", fila["activo"])
	}
	for _, columna := range []string{"fecha_creacion", "fecha_modificacion"} {
		valor, _ := fila[columna].(string)
		if valor == "" || strings.HasPrefix(valor, "2000-") {
			t.Errorf("Se esperaba %s asignada por el servicio y se obtuvo %v", columna, fila[columna])
		}
	}
}

// ProbarPostInvalido comprueba que POST responda 400 ante un JSON mal formado
// o un cuerpo vacío.
func ProbarPostInvalido(t *testing.T, ruta string) {
	t.Helper()
	VerificarError(t, Peticion(http.MethodPost, ruta+"/", `{"Id":`), http.StatusBadRequest)
	VerificarError(t, Peticion(http.MethodPost, ruta+"/", ""), http.StatusBadRequest)
}

// ProbarGetOne consulta el registro por Id, verifica la respuesta 200 y
// retorna el objeto obtenido.
func ProbarGetOne(t *testing.T, ruta string, id int) map[string]interface{} {
	t.Helper()
	r := VerificarExito(t, Peticion(http.MethodGet, ruta+"/"+strconv.Itoa(id), ""), http.StatusOK)
	VerificarIgual(t, "Id del registro consultado", r.IdDato(t), id)
	return r.Dato(t)
}

// ProbarGetOneInexistente comprueba que GetOne responda 404 ante un Id que no
// existe o que no es numérico.
func ProbarGetOneInexistente(t *testing.T, ruta string) {
	t.Helper()
	VerificarError(t, Peticion(http.MethodGet, ruta+"/"+strconv.Itoa(IdInexistente), ""), http.StatusNotFound)
	VerificarError(t, Peticion(http.MethodGet, ruta+"/abc", ""), http.StatusNotFound)
}

// ConsultarGetAll ejecuta GetAll con los parámetros indicados y retorna la
// lista obtenida.
func ConsultarGetAll(t *testing.T, ruta, parametros string) []map[string]interface{} {
	t.Helper()
	return VerificarExito(t, Peticion(http.MethodGet, ruta+"/?"+parametros, ""), http.StatusOK).Lista(t)
}

// ContieneId indica si la lista contiene un objeto con el Id indicado.
func ContieneId(lista []map[string]interface{}, id int) bool {
	for _, m := range lista {
		if m["Id"] == float64(id) {
			return true
		}
	}
	return false
}

// ProbarGetAll verifica el servicio GetAll: filtro por Id, lista vacía sin
// resultados, limit/sortby/order/fields y HTTP 400 ante filtros inválidos.
// Retorna el registro obtenido con el filtro por Id.
func ProbarGetAll(t *testing.T, ruta string, id int) map[string]interface{} {
	t.Helper()
	lista := ConsultarGetAll(t, ruta, fmt.Sprintf("query=Id:%d", id))
	if len(lista) != 1 || !ContieneId(lista, id) {
		t.Fatalf("Se esperaba solo el registro %d y se obtuvo %v", id, lista)
	}

	if l := ConsultarGetAll(t, ruta, fmt.Sprintf("query=Id:%d", IdInexistente)); len(l) != 0 {
		t.Error("Se esperaba una lista vacía (Data: []) y se obtuvo", l)
	}
	if l := ConsultarGetAll(t, ruta, "sortby=Id&order=desc&limit=1"); len(l) != 1 {
		t.Error("Se esperaba 1 registro con limit=1 y se obtuvo", len(l))
	}
	l := ConsultarGetAll(t, ruta, fmt.Sprintf("query=Id:%d&fields=Id", id))
	if len(l) != 1 || len(l[0]) != 1 || l[0]["Id"] != float64(id) {
		t.Error("Se esperaba solo el campo Id con fields=Id y se obtuvo", l)
	}

	// Filtros inválidos: HTTP 400 con la estructura institucional
	VerificarFiltroInvalido(t, ruta+"/?query=Id")
	VerificarFiltroInvalido(t, fmt.Sprintf("%s/?query=Id:%d,Activo", ruta, id))
	return lista[0]
}

// VerificarFKExpandida comprueba que la llave foránea campo venga expandida
// con los mismos datos del registro relacionado.
func VerificarFKExpandida(t *testing.T, dato map[string]interface{}, campo string, relacionado interface{}) {
	t.Helper()
	VerificarIgual(t, "Llave foránea "+campo+" expandida", dato[campo], AMapa(t, relacionado))
}

// ProbarPut verifica el servicio PUT: actualiza el campo indicado, conserva
// fecha_creacion tras dos actualizaciones, actualiza fecha_modificacion y
// responde 400 ante Id distinto, Id inexistente o JSON inválido.
func ProbarPut(t *testing.T, ruta, tabla string, id int, campo string, valor interface{}, columna string) {
	t.Helper()
	rutaId := ruta + "/" + strconv.Itoa(id)
	antes := Fila(t, tabla, id)
	cuerpo := ProbarGetOne(t, ruta, id)
	delete(cuerpo, "FechaCreacion")
	delete(cuerpo, "FechaModificacion")
	cuerpo[campo] = valor

	for i := 0; i < 2; i++ {
		VerificarExito(t, Peticion(http.MethodPut, rutaId, JSON(t, cuerpo)), http.StatusOK)
	}
	despues := Fila(t, tabla, id)
	VerificarIgual(t, "Columna "+columna+" actualizada", fmt.Sprint(despues[columna]), fmt.Sprint(valor))
	VerificarIgual(t, "fecha_creacion se conserva tras el PUT", despues["fecha_creacion"], antes["fecha_creacion"])
	if despues["fecha_modificacion"] == nil || despues["fecha_modificacion"] == antes["fecha_modificacion"] {
		t.Error("Se esperaba fecha_modificacion actualizada y se obtuvo", despues["fecha_modificacion"])
	}

	cuerpo["Id"] = id + 1
	VerificarError(t, Peticion(http.MethodPut, rutaId, JSON(t, cuerpo)), http.StatusBadRequest)
	delete(cuerpo, "Id")
	VerificarError(t, Peticion(http.MethodPut, ruta+"/"+strconv.Itoa(IdInexistente), JSON(t, cuerpo)), http.StatusBadRequest)
	VerificarError(t, Peticion(http.MethodPut, rutaId, `{"Id":"texto"}`), http.StatusBadRequest)
}

// ProbarDelete verifica el borrado lógico: responde 200, deja activo=false,
// actualiza fecha_modificacion, no modifica las demás columnas, el registro
// sigue consultable y un Id inexistente responde 404.
func ProbarDelete(t *testing.T, ruta, tabla string, id int) {
	t.Helper()
	antes := Fila(t, tabla, id)
	VerificarExito(t, Peticion(http.MethodDelete, ruta+"/"+strconv.Itoa(id), ""), http.StatusOK)
	despues := Fila(t, tabla, id)

	if despues["activo"] != false {
		t.Error("Se esperaba activo=false (borrado lógico) y se obtuvo", despues["activo"])
	}
	if despues["fecha_modificacion"] == nil || despues["fecha_modificacion"] == antes["fecha_modificacion"] {
		t.Error("Se esperaba fecha_modificacion actualizada y se obtuvo", despues["fecha_modificacion"])
	}
	VerificarIgual(t, "Las demás columnas no cambian con el borrado lógico",
		SinColumnas(despues, "activo", "fecha_modificacion"), SinColumnas(antes, "activo", "fecha_modificacion"))

	if dato := ProbarGetOne(t, ruta, id); dato["Activo"] != false {
		t.Error("Se esperaba el registro consultable con Activo=false y se obtuvo", dato["Activo"])
	}
	VerificarError(t, Peticion(http.MethodDelete, ruta+"/"+strconv.Itoa(IdInexistente), ""), http.StatusNotFound)
}
