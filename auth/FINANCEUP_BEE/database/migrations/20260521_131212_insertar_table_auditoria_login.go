package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertarTableAuditoriaLogin_20260521_131212 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTableAuditoriaLogin_20260521_131212{}
	m.Created = "20260521_131212"

	migration.Register("InsertarTableAuditoriaLogin_20260521_131212", m)
}

// Run the migrations
func (m *InsertarTableAuditoriaLogin_20260521_131212) Up() {
	m.SQL("INSERT INTO auth.auditoria_login (id_usuario, tipo_evento, ip_address, navegador, estado_evento) VALUES (1, 'login', '192.168.1.1', 'Chrome 120', 'exitoso')")

}

// Reverse the migrations
func (m *InsertarTableAuditoriaLogin_20260521_131212) Down() {
	m.SQL("DELETE FROM auth.auditoria_login WHERE id_usuario = 1")
}
