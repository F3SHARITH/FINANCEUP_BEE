package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableCategoria_20260521_135301 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableCategoria_20260521_135301{}
	m.Created = "20260521_135301"

	migration.Register("InsertTableCategoria_20260521_135301", m)
}

// Run the migrations
func (m *InsertTableCategoria_20260521_135301) Up() {
	m.SQL(`INSERT INTO "finanzas"."categoria" ("nombre", "descripcion", "activo") VALUES
		('Arriendo',      'Gastos de vivienda',    true),
		('Servicios',     'Servicios publicos',    true),
		('Alimentacion',  'Gastos en alimentos',  true),
		('Transporte',    'Gastos en transporte',  true),
		('Salud',         'Gastos medicos',        true)`)
}

// Reverse the migrations
func (m *InsertTableCategoria_20260521_135301) Down() {
	m.SQL(`DELETE FROM "finanzas"."categoria" WHERE "nombre" IN ('Arriendo', 'Servicios', 'Alimentacion', 'Transporte', 'Salud')`)
}
