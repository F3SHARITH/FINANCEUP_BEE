package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableNivelRiesgo_20260521_135306 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableNivelRiesgo_20260521_135306{}
	m.Created = "20260521_135306"

	migration.Register("InsertTableNivelRiesgo_20260521_135306", m)
}

// Run the migrations
func (m *InsertTableNivelRiesgo_20260521_135306) Up() {
	m.SQL(`INSERT INTO "finanzas"."nivel_riesgo" ("nombre", "activo") VALUES
		('Bajo',  true),
		('Medio', true),
		('Alto',  true)`)
}

// Reverse the migrations
func (m *InsertTableNivelRiesgo_20260521_135306) Down() {
	m.SQL(`DELETE FROM "finanzas"."nivel_riesgo" WHERE "nombre" IN ('Bajo', 'Medio', 'Alto')`)
}
