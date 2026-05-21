package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTransaccionComision_20260521_100112 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTransaccionComision_20260521_100112{}
	m.Created = "20260521_100112"

	migration.Register("InsertTableTransaccionComision_20260521_100112", m)
}

// Run the migrations
func (m *InsertTableTransaccionComision_20260521_100112) Up() {
	m.SQL(`INSERT INTO "negocio"."transaccion_comision"
		("id_credito", "id_banco", "monto_comision", "porcentaje_aplicado", "fecha_transaccion", "estado", "referencia_pago", "activo")
	VALUES
		(1, 1, 125000, 2.50, CURRENT_TIMESTAMP - INTERVAL '48 days', 'pagada', 'TRANS-001-2024', true),
		(2, 1, 5000000, 2.50, CURRENT_TIMESTAMP - INTERVAL '18 days', 'pendiente', 'TRANS-002-2024', true),
		(3, 2, 125000, 2.50, CURRENT_TIMESTAMP - INTERVAL '38 days', 'pagada', 'TRANS-003-2024', true),
		(4, 3, 1125000, 2.25, CURRENT_TIMESTAMP - INTERVAL '5 days', 'pendiente', 'TRANS-004-2024', true)`)
}

// Reverse the migrations
func (m *InsertTableTransaccionComision_20260521_100112) Down() {
	m.SQL(`DELETE FROM "negocio"."transaccion_comision"
	WHERE "referencia_pago" IN (
		'TRANS-001-2024',
		'TRANS-002-2024',
		'TRANS-003-2024',
		'TRANS-004-2024'
	)`)
}
