package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertarTableUsuario_20260521_133250 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTableUsuario_20260521_133250{}
	m.Created = "20260521_133250"

	migration.Register("InsertarTableUsuario_20260521_133250", m)
}

// Run the migrations
func (m *InsertarTableUsuario_20260521_133250) Up() {
	m.SQL("INSERT INTO auth.usuario (tipo_documento, nombre, apellido, email, telefono, cedula, ciudad, estado, activo) VALUES(1, 'harold', 'arciniegas', 'harold.arciniegas@email.com', '3001234567', '1234567890', 'Bogota', 'activo', true)")

}

// Reverse the migrations
func (m *InsertarTableUsuario_20260521_133250) Down() {
	m.SQL("DELETE FROM auth.usuario WHERE cedula = '1234567890'")

}
