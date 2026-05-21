package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableContenido_20260521_100302 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableContenido_20260521_100302{}
	m.Created = "20260521_100302"

	migration.Register("InsertTableContenido_20260521_100302", m)
}

// Run the migrations
func (m *InsertTableContenido_20260521_100302) Up() {
	m.SQL(`INSERT INTO "educacion"."contenido"
		("titulo", "descripcion", "duracion_minutos", "url_video", "activo")
	VALUES
		('Video Introduccion Finanzas', 'Video introductorio', 15, 'https://example.com/videos/intro1.mp4', true),
		('Video Creditos Explicado', 'Explicacion creditos', 20, 'https://example.com/videos/creditos1.mp4', true),
		('Video Inversiones', 'Guia de inversiones', 25, 'https://example.com/videos/inversiones1.mp4', true),
		('Video Ahorro Basico', 'Conceptos de ahorro', 18, 'https://example.com/videos/ahorro1.mp4', true)`)
}

// Reverse the migrations
func (m *InsertTableContenido_20260521_100302) Down() {
	m.SQL(`DELETE FROM "educacion"."contenido"
	WHERE "titulo" IN (
		'Video Introduccion Finanzas',
		'Video Creditos Explicado',
		'Video Inversiones',
		'Video Ahorro Basico'
	)`)
}
