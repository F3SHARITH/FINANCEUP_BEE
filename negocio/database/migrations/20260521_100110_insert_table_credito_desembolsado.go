package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableCreditoDesembolsado_20260521_100110 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableCreditoDesembolsado_20260521_100110{}
	m.Created = "20260521_100110"

	migration.Register("InsertTableCreditoDesembolsado_20260521_100110", m)
}

// Run the migrations
func (m *InsertTableCreditoDesembolsado_20260521_100110) Up() {
	m.SQL(`INSERT INTO "negocio"."credito_desembolsado"
		("id_lead", "id_usuario", "id_producto", "id_banco", "numero_credito", "monto_aprobado", "tasa_interes_final", "plazo_meses", "fecha_aprobacion", "fecha_desembolso", "estado_credito", "saldo_actual", "activo")
	VALUES
		(1, 1, 1, 1, 'CRED-2024-001', 5000000, 15.50, 36, CURRENT_DATE - INTERVAL '50 days', CURRENT_DATE - INTERVAL '48 days', 'activo', 4800000, true),
		(2, 2, 2, 1, 'CRED-2024-002', 200000000, 10.00, 240, CURRENT_DATE - INTERVAL '20 days', CURRENT_DATE - INTERVAL '18 days', 'activo', 199800000, true),
		(3, 3, 3, 2, 'CRED-2024-003', 5000000, 18.75, 36, CURRENT_DATE - INTERVAL '40 days', CURRENT_DATE - INTERVAL '38 days', 'activo', 4700000, true),
		(4, 4, 4, 3, 'CRED-2024-004', 50000000, 12.50, 60, CURRENT_DATE - INTERVAL '5 days', NULL, 'activo', 50000000, true)`)
}

// Reverse the migrations
func (m *InsertTableCreditoDesembolsado_20260521_100110) Down() {
	m.SQL(`DELETE FROM "negocio"."credito_desembolsado"
	WHERE "numero_credito" IN (
		'CRED-2024-001',
		'CRED-2024-002',
		'CRED-2024-003',
		'CRED-2024-004'
	)`)
}
