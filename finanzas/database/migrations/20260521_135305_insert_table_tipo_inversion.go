package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTipoInversion_20260521_135305 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTipoInversion_20260521_135305{}
	m.Created = "20260521_135305"

	migration.Register("InsertTableTipoInversion_20260521_135305", m)
}

// Run the migrations
func (m *InsertTableTipoInversion_20260521_135305) Up() {
	m.SQL(`INSERT INTO "finanzas"."tipo_inversion" ("nombre", "descripcion", "activo") VALUES
		('Acciones',        'Inversion en acciones de bolsa',         true),
		('Fondos Mutuales', 'Inversion en fondos mutuales',           true),
		('Bonos',           'Inversion en bonos gubernamentales',     true),
		('Criptomonedas',   'Inversion en criptomonedas',             true)`)
}

// Reverse the migrations
func (m *InsertTableTipoInversion_20260521_135305) Down() {
	m.SQL(`DELETE FROM "finanzas"."tipo_inversion" WHERE "nombre" IN ('Acciones', 'Fondos Mutuales', 'Bonos', 'Criptomonedas')`)
}
