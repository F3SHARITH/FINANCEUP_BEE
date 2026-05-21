package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertarTableCredencial_20260521_132208 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTableCredencial_20260521_132208{}
	m.Created = "20260521_132208"

	migration.Register("InsertarTableCredencial_20260521_132208", m)
}

// Run the migrations
func (m *InsertarTableCredencial_20260521_132208) Up() {
	m.SQL("INSERT INTO auth.credencial (id_usuario, contrasena_hash, salt, algoritmo, intentos_fallidos, requiere_cambio, activo) VALUES (1, '$2b$10$abcdefghijklmnopqrstuvwxyz123456789', 'salt123', 'bcrypt', 0, false, true)")

}

// Reverse the migrations
func (m *InsertarTableCredencial_20260521_132208) Down() {
	m.SQL("DELETE FROM auth.credencial WHERE id_usuario = 1")

}
