package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableMovimientoMeta_20260521_135311 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableMovimientoMeta_20260521_135311{}
	m.Created = "20260521_135311"

	migration.Register("InsertTableMovimientoMeta_20260521_135311", m)
}

// Run the migrations
func (m *InsertTableMovimientoMeta_20260521_135311) Up() {
	m.SQL(`INSERT INTO "finanzas"."movimiento_meta"
		("nombre", "monto", "es_ingreso", "activo")
	VALUES
		('Aporte Meta Enero',   500000, true, true),
		('Aporte Meta Febrero', 500000, true, true),
		('Aporte Meta Marzo',   500000, true, true)`)
}

// Reverse the migrations
func (m *InsertTableMovimientoMeta_20260521_135311) Down() {
	m.SQL(`DELETE FROM "finanzas"."movimiento_meta" WHERE "nombre" IN ('Aporte Meta Enero', 'Aporte Meta Febrero', 'Aporte Meta Marzo')`)
}
