package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTipoIngresoInversion_20260521_135308 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTipoIngresoInversion_20260521_135308{}
	m.Created = "20260521_135308"

	migration.Register("InsertTableTipoIngresoInversion_20260521_135308", m)
}

// Run the migrations
func (m *InsertTableTipoIngresoInversion_20260521_135308) Up() {
	m.SQL(`INSERT INTO "finanzas"."tipo_ingreso_inversion"
		("id_movimiento_dinero", "nombre_movimiento_pago", "descripcion", "activo")
	VALUES
		(1, 'Rendimiento', 'Ganancia por rendimiento de inversiones', true),
		(2, 'Dividendos',  'Dividendos de acciones',                  true)`)
}

// Reverse the migrations
func (m *InsertTableTipoIngresoInversion_20260521_135308) Down() {
	m.SQL(`DELETE FROM "finanzas"."tipo_ingreso_inversion" WHERE "nombre_movimiento_pago" IN ('Rendimiento', 'Dividendos')`)
}
