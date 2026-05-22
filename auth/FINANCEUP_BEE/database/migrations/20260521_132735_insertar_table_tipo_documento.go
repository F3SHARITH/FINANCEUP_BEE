package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertarTableTipoDocumento_20260521_132735 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTableTipoDocumento_20260521_132735{}
	m.Created = "20260521_132735"

	migration.Register("InsertarTableTipoDocumento_20260521_132735", m)
}

// Run the migrations
func (m *InsertarTableTipoDocumento_20260521_132735) Up() {
	m.SQL("INSERT INTO auth.tipo_documento (nombre, codigo, activo) VALUES ('Cedula de Ciudadania', 'CC', true)")

}

// Reverse the migrations
func (m *InsertarTableTipoDocumento_20260521_132735) Down() {
	m.SQL("DELETE FROM auth.tipo_documento WHERE codigo = 'CC'")
}
