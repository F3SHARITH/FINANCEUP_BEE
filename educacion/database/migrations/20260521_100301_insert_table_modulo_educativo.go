package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableModuloEducativo_20260521_100301 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableModuloEducativo_20260521_100301{}
	m.Created = "20260521_100301"

	migration.Register("InsertTableModuloEducativo_20260521_100301", m)
}

// Run the migrations
func (m *InsertTableModuloEducativo_20260521_100301) Up() {
	m.SQL(`INSERT INTO "educacion"."modulo_educativo"
		("titulo", "descripcion", "contenido", "nivel", "url_thumbnail", "activo")
	VALUES
		('Introduccion a Finanzas Personales', 'Conceptos basicos de finanzas personales', 'Contenido teorico basico', 'basico', 'https://example.com/thumb1.jpg', true),
		('Creditos y Deudas', 'Como gestionar creditos y deudas', 'Contenido intermedio sobre gestion', 'intermedio', 'https://example.com/thumb2.jpg', true),
		('Inversiones Avanzadas', 'Estrategias avanzadas de inversion', 'Contenido avanzado de portafolio', 'avanzado', 'https://example.com/thumb3.jpg', true),
		('Ahorro e Inversion', 'Tecnicas de ahorro e inversion', 'Contenido sobre ahorro', 'basico', 'https://example.com/thumb4.jpg', true)`)
}

// Reverse the migrations
func (m *InsertTableModuloEducativo_20260521_100301) Down() {
	m.SQL(`DELETE FROM "educacion"."modulo_educativo"
	WHERE "titulo" IN (
		'Introduccion a Finanzas Personales',
		'Creditos y Deudas',
		'Inversiones Avanzadas',
		'Ahorro e Inversion'
	)`)
}
