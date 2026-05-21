package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableProgresoLeccion_20260521_100305 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableProgresoLeccion_20260521_100305{}
	m.Created = "20260521_100305"

	migration.Register("InsertTableProgresoLeccion_20260521_100305", m)
}

// Run the migrations
func (m *InsertTableProgresoLeccion_20260521_100305) Up() {
	m.SQL(`INSERT INTO "educacion"."progreso_leccion"
		("id_usuario", "id_leccion", "completado", "fecha_inicio", "fecha_completado", "activo")
	VALUES
		(1, 1, true, CURRENT_TIMESTAMP - INTERVAL '30 days', CURRENT_TIMESTAMP - INTERVAL '28 days', true),
		(1, 2, true, CURRENT_TIMESTAMP - INTERVAL '27 days', CURRENT_TIMESTAMP - INTERVAL '25 days', true),
		(2, 1, true, CURRENT_TIMESTAMP - INTERVAL '20 days', CURRENT_TIMESTAMP - INTERVAL '19 days', true),
		(3, 3, true, CURRENT_TIMESTAMP - INTERVAL '15 days', CURRENT_TIMESTAMP - INTERVAL '14 days', true)`)
}

// Reverse the migrations
func (m *InsertTableProgresoLeccion_20260521_100305) Down() {
	m.SQL(`DELETE FROM "educacion"."progreso_leccion"
	WHERE ("id_usuario", "id_leccion") IN ((1, 1), (1, 2), (2, 1), (3, 3))`)
}
