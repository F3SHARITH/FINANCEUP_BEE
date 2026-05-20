package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertarTablaRegistroActividad_20260520_172008 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTablaRegistroActividad_20260520_172008{}
	m.Created = "20260520_172008"

	migration.Register("InsertarTablaRegistroActividad_20260520_172008", m)
}

// Run the migrations
func (m *InsertarTablaRegistroActividad_20260520_172008) Up() {
	m.SQL("INSERT INTO soporte.registro_actividad (id_usuario, tipo_actividad, descripcion, entidad_afectada, fecha_actividad) VALUES (1, 'LOGIN',            'Usuario inicio sesion',                   'auth.usuario',       CURRENT_TIMESTAMP - INTERVAL '2 hours'),") 

}

// Reverse the migrations
func (m *InsertarTablaRegistroActividad_20260520_172008) Down() {
	m.SQL("DELETE FROM soporte.registro_actividad WHERE id_usuario =1; ") 

}
