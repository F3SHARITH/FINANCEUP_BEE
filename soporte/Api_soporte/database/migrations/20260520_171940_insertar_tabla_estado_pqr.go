package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertarTablaEstadoPqr_20260520_171940 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTablaEstadoPqr_20260520_171940{}
	m.Created = "20260520_171940"

	migration.Register("InsertarTablaEstadoPqr_20260520_171940", m)
}

// Run the migrations
func (m *InsertarTablaEstadoPqr_20260520_171940) Up() {
	m.SQL("INSERT INTO soporte.estado_pqr (nombre, descripcion, activo) VALUES ('Abierta', 'Peticion recientemente registrada', true)") 

}

// Reverse the migrations
func (m *InsertarTablaEstadoPqr_20260520_171940) Down() {
	m.SQL("DELETE FROM soporte.estado_pqr WHERE nombre = 'Abierta'")

}
