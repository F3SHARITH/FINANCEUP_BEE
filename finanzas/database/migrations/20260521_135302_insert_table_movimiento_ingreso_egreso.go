package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableMovimientoIngresoEgreso_20260521_135302 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableMovimientoIngresoEgreso_20260521_135302{}
	m.Created = "20260521_135302"

	migration.Register("InsertTableMovimientoIngresoEgreso_20260521_135302", m)
}

// Run the migrations
func (m *InsertTableMovimientoIngresoEgreso_20260521_135302) Up() {
	m.SQL(`INSERT INTO "finanzas"."movimiento_ingreso_egreso"
		("nombre", "monto", "es_ingreso", "activo")
	VALUES
		('Salario Enero',      5000000, true,  true),
		('Gasto Arriendo',     1500000, false, true),
		('Freelance Enero',     800000, true,  true),
		('Gasto Servicios',     300000, false, true),
		('Gasto Alimentacion',  600000, false, true)`)
}

// Reverse the migrations
func (m *InsertTableMovimientoIngresoEgreso_20260521_135302) Down() {
	m.SQL(`DELETE FROM "finanzas"."movimiento_ingreso_egreso" WHERE "nombre" IN ('Salario Enero', 'Gasto Arriendo', 'Freelance Enero', 'Gasto Servicios', 'Gasto Alimentacion')`)
}
