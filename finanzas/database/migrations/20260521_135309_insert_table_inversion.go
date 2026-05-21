package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableInversion_20260521_135309 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableInversion_20260521_135309{}
	m.Created = "20260521_135309"

	migration.Register("InsertTableInversion_20260521_135309", m)
}

// Run the migrations
func (m *InsertTableInversion_20260521_135309) Up() {
	m.SQL(`INSERT INTO "finanzas"."inversion"
		("id_usuario", "id_tipo_inversion", "id_nivel_riesgo", "id_movimiento_dinero",
		 "nombre", "monto", "rentabilidad", "fecha_inicio", "fecha_fin", "activo")
	VALUES
		(1, 1, 2, 1,    'Acciones Tecnologicas', 1000000, 12.50, CURRENT_DATE - INTERVAL '180 days', CURRENT_DATE + INTERVAL '180 days', true),
		(1, 2, 1, NULL, 'Fondo Moderado',         500000,  8.00, CURRENT_DATE - INTERVAL '365 days', NULL,                               true),
		(2, 3, 1, 3,    'Bonos Estatales',       2000000,  6.75, CURRENT_DATE - INTERVAL '90 days',  CURRENT_DATE + INTERVAL '270 days', true),
		(3, 1, 3, NULL, 'Acciones de Riesgo',     500000, 25.00, CURRENT_DATE - INTERVAL '30 days',  NULL,                               true)`)
}

// Reverse the migrations
func (m *InsertTableInversion_20260521_135309) Down() {
	m.SQL(`DELETE FROM "finanzas"."inversion" WHERE "nombre" IN ('Acciones Tecnologicas', 'Fondo Moderado', 'Bonos Estatales', 'Acciones de Riesgo')`)
}
