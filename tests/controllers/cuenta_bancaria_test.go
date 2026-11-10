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

// Pruebas del controlador CuentaBancariaController (/v1/cuenta_bancaria).

const (
	rutaCuentaBancaria  = "/v1/cuenta_bancaria"
	tablaCuentaBancaria = "cuenta_bancaria"
)

// cuerpoPostCuentaBancaria retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostCuentaBancaria(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoCuentaBancaria(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestCuentaBancariaPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaCuentaBancaria, tablaCuentaBancaria, cuerpoPostCuentaBancaria(t))
	id2 := helpers.CrearPorPost(t, rutaCuentaBancaria, tablaCuentaBancaria, cuerpoPostCuentaBancaria(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaCuentaBancaria, id)

	// La llave foránea ProveedorId es obligatoria
	sinProveedorId := cuerpoPostCuentaBancaria(t)
	sinProveedorId["ProveedorId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaCuentaBancaria+"/", helpers.JSON(t, sinProveedorId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaCuentaBancaria)
	t.Log("TestCuentaBancariaPost Finalizado Correctamente (OK)")
}

func TestCuentaBancariaGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCuentaBancaria(t)
	dato := helpers.ProbarGetOne(t, rutaCuentaBancaria, creado.Id)

	// La llave foránea ContactoId viene expandida
	if relacionado, err := models.GetContactoById(creado.ContactoId.Id); err != nil {
		t.Fatal("No se pudo consultar Contacto:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ContactoId", relacionado)
	}

	// La llave foránea ProveedorId viene expandida
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaCuentaBancaria)
	t.Log("TestCuentaBancariaGetOne Finalizado Correctamente (OK)")
}

func TestCuentaBancariaGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCuentaBancaria(t)
	dato := helpers.ProbarGetAll(t, rutaCuentaBancaria, creado.Id)

	// La llave foránea ContactoId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetContactoById(creado.ContactoId.Id); err != nil {
		t.Fatal("No se pudo consultar Contacto:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ContactoId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaCuentaBancaria, fmt.Sprintf("query=ContactoId.Id:%d", creado.ContactoId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por ContactoId.Id y se obtuvo", l)
	}

	// La llave foránea ProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetProveedorById(creado.ProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar Proveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaCuentaBancaria, fmt.Sprintf("query=ProveedorId.Id:%d", creado.ProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por ProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestCuentaBancariaGetAll Finalizado Correctamente (OK)")
}

func TestCuentaBancariaPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCuentaBancaria(t)
	helpers.ProbarPut(t, rutaCuentaBancaria, tablaCuentaBancaria, creado.Id, "NumeroCuenta", "Prueba-2", "numero_cuenta")
	t.Log("TestCuentaBancariaPut Finalizado Correctamente (OK)")
}

func TestCuentaBancariaDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCuentaBancaria(t)
	helpers.ProbarDelete(t, rutaCuentaBancaria, tablaCuentaBancaria, creado.Id)
	t.Log("TestCuentaBancariaDelete Finalizado Correctamente (OK)")
}

// La llave foránea ContactoId es opcional (null): el registro sin ella se
// puede crear, consultar, actualizar y borrar.
func TestCuentaBancariaContactoIdOpcional(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCuentaBancaria(t)
	cuerpo["ContactoId"] = nil
	id := helpers.CrearPorPost(t, rutaCuentaBancaria, tablaCuentaBancaria, cuerpo)

	if dato := helpers.ProbarGetOne(t, rutaCuentaBancaria, id); dato["ContactoId"] != nil {
		t.Error("Se esperaba ContactoId nulo y se obtuvo", dato["ContactoId"])
	}
	if l := helpers.ConsultarGetAll(t, rutaCuentaBancaria, fmt.Sprintf("query=Id:%d", id)); !helpers.ContieneId(l, id) {
		t.Error("Se esperaba el registro sin ContactoId en GetAll y se obtuvo", l)
	}
	helpers.ProbarPut(t, rutaCuentaBancaria, tablaCuentaBancaria, id, "NumeroCuenta", "Prueba-2", "numero_cuenta")
	helpers.ProbarDelete(t, rutaCuentaBancaria, tablaCuentaBancaria, id)
	t.Log("TestCuentaBancariaContactoIdOpcional Finalizado Correctamente (OK)")
}
