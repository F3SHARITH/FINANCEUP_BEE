package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertarTablaPqr_20260520_171953 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTablaPqr_20260520_171953{}
	m.Created = "20260520_171953"

	migration.Register("InsertarTablaPqr_20260520_171953", m)
}

// Run the migrations
func (m *InsertarTablaPqr_20260520_171953) Up() {
	m.SQL("INSERT INTO soporte.pqr (id_usuario, descripcion, id_estado, activo) VALUES (1, 'Problema con acceso a la plataforma',1, true),") 

}

// Reverse the migrations
func (m *InsertarTablaPqr_20260520_171953) Down() {
	m.SQL("DELETE FROM soporte.pqr WHERE descripcion = 'Problema con acceso a la plataforma'")

}
