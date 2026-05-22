package main
import (
	"github.com/beego/beego/v2/client/orm/migration"
	)

// DO NOT MODIFY
type InsertarTableRoles_20260520_172305 struct {
	migration.Migration
						}

// DO NOT MODIFY
func init() {
	m := &InsertarTableRoles_20260520_172305{}
	m.Created = "20260520_172305"
							
	migration.Register("InsertarTableRoles_20260520_172305", m)
						}
					   
// Run the migrations
func (m *InsertarTableRoles_20260520_172305) Up() {
	m.SQL("INSERT INTO auth.usuario_rol (id_usuario, id_rol, activo )VALUES (1, 1, true)")
					
}
// Reverse the migrations
func (m *InsertarTableRoles_20260520_172305) Down() {
	m.SQL("DELETE FROM auth.usuario_rol WHERE id_usuario = 1")
					
}
				