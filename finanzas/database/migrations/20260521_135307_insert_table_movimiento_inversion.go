package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableMovimientoInversion_20260521_135307 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableMovimientoInversion_20260521_135307{}
	m.Created = "20260521_135307"

	migration.Register("InsertTableMovimientoInversion_20260521_135307", m)
}

// Run the migrations
func (m *InsertTableMovimientoInversion_20260521_135307) Up() {
	m.SQL(`INSERT INTO "finanzas"."movimiento_inversion"
		("nombre", "monto", "es_ingreso", "activo")
	VALUES
		('Inversion Acciones', 1000000, false, true),
		('Ganancia Fondos',      50000, true,  true),
		('Inversion Bonos',    2000000, false, true)`)
}

// Reverse the migrations
func (m *InsertTableMovimientoInversion_20260521_135307) Down() {
	m.SQL(`DELETE FROM "finanzas"."movimiento_inversion" WHERE "nombre" IN ('Inversion Acciones', 'Ganancia Fondos', 'Inversion Bonos')`)
}
