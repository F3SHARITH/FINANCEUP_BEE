package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableFinanzas_20260521_135304 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableFinanzas_20260521_135304{}
	m.Created = "20260521_135304"

	migration.Register("InsertTableFinanzas_20260521_135304", m)
}

// Run the migrations
func (m *InsertTableFinanzas_20260521_135304) Up() {
	m.SQL(`INSERT INTO "finanzas"."finanzas"
		("id_usuario", "id_movimiento_dinero", "id_categoria", "monto_presupuesto", "gasto", "disponible", "fecha", "activo")
	VALUES
		(1, 1,    1, 5000000, 3200000, 1800000, CURRENT_DATE, true),
		(2, NULL, 2, 3500000, 2100000, 1400000, CURRENT_DATE, true),
		(3, 3,    3, 4200000, 2800000, 1400000, CURRENT_DATE, true),
		(4, 1,    4, 6000000, 3500000, 2500000, CURRENT_DATE, true)`)
}

// Reverse the migrations
func (m *InsertTableFinanzas_20260521_135304) Down() {
	m.SQL(`DELETE FROM "finanzas"."finanzas" WHERE "id_usuario" IN (1, 2, 3, 4)`)
}
