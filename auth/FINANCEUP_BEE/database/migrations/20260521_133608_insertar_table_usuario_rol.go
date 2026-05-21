package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertarTableUsuarioRol_20260521_133608 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTableUsuarioRol_20260521_133608{}
	m.Created = "20260521_133608"

	migration.Register("InsertarTableUsuarioRol_20260521_133608", m)
}

// Run the migrations
func (m *InsertarTableUsuarioRol_20260521_133608) Up() {
	m.SQL("INSERT INTO auth.usuario_rol (id_usuario, id_rol, activo) VALUES (1, 1, true)")
	
}

// Reverse the migrations
func (m *InsertarTableUsuarioRol_20260521_133608) Down() {
	m.SQL("DELETE FROM auth.usuario_rol WHERE id_usuario = 1")

}
