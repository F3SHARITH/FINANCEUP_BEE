package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableLeccion_20260521_100303 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableLeccion_20260521_100303{}
	m.Created = "20260521_100303"

	migration.Register("InsertTableLeccion_20260521_100303", m)
}

// Run the migrations
func (m *InsertTableLeccion_20260521_100303) Up() {
	m.SQL(`INSERT INTO "educacion"."leccion"
		("id_modulo", "id_contenido", "titulo", "descripcion", "duracion_minutos", "url_video", "numero_leccion", "activo")
	VALUES
		(1, 1, 'Leccion 1: Conceptos Basicos', 'Primera leccion del modulo', 30, 'https://example.com/leccion1.mp4', 1, true),
		(1, NULL, 'Leccion 2: Presupuesto Personal', 'Segunda leccion del modulo', 35, 'https://example.com/leccion2.mp4', 2, true),
		(2, 2, 'Leccion 1: Tipos de Credito', 'Primer tema de creditos', 40, 'https://example.com/leccion3.mp4', 1, true),
		(3, 3, 'Leccion 1: Analisis de Riesgo', 'Analisis avanzado', 45, 'https://example.com/leccion4.mp4', 1, true)`)
}

// Reverse the migrations
func (m *InsertTableLeccion_20260521_100303) Down() {
	m.SQL(`DELETE FROM "educacion"."leccion"
	WHERE "titulo" IN (
		'Leccion 1: Conceptos Basicos',
		'Leccion 2: Presupuesto Personal',
		'Leccion 1: Tipos de Credito',
		'Leccion 1: Analisis de Riesgo'
	)`)
}
