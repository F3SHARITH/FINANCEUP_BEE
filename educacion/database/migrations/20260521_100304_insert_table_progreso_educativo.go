package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableProgresoEducativo_20260521_100304 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableProgresoEducativo_20260521_100304{}
	m.Created = "20260521_100304"

	migration.Register("InsertTableProgresoEducativo_20260521_100304", m)
}

// Run the migrations
func (m *InsertTableProgresoEducativo_20260521_100304) Up() {
	m.SQL(`INSERT INTO "educacion"."progreso_educativo"
		("id_usuario", "id_modulo", "porcentaje_completado", "fecha_inicio", "fecha_completado", "calificacion", "activo")
	VALUES
		(1, 1, 100, CURRENT_TIMESTAMP - INTERVAL '30 days', CURRENT_TIMESTAMP - INTERVAL '5 days', 95, true),
		(2, 1, 50, CURRENT_TIMESTAMP - INTERVAL '20 days', NULL, NULL, true),
		(3, 2, 75, CURRENT_TIMESTAMP - INTERVAL '15 days', NULL, NULL, true),
		(4, 1, 100, CURRENT_TIMESTAMP - INTERVAL '10 days', CURRENT_TIMESTAMP - INTERVAL '2 days', 88, true)`)
}

// Reverse the migrations
func (m *InsertTableProgresoEducativo_20260521_100304) Down() {
	m.SQL(`DELETE FROM "educacion"."progreso_educativo"
	WHERE ("id_usuario", "id_modulo") IN ((1, 1), (2, 1), (3, 2), (4, 1))`)
}
