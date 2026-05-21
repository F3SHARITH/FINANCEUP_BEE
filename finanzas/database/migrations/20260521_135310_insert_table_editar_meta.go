package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableEditarMeta_20260521_135310 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableEditarMeta_20260521_135310{}
	m.Created = "20260521_135310"

	migration.Register("InsertTableEditarMeta_20260521_135310", m)
}

// Run the migrations
func (m *InsertTableEditarMeta_20260521_135310) Up() {
	m.SQL(`INSERT INTO "finanzas"."editar_meta"
		("nombre", "monto_actual", "monto_objetivo", "ahorro_mensual", "fecha_objetivo", "descripcion", "activo")
	VALUES
		('Fondo Emergencia', 4500000,  10000000,  500000, CURRENT_DATE + INTERVAL '365 days', 'Ahorrar 6 meses de gastos', true),
		('Vacaciones Europa',5000000,   5000000, 1000000, CURRENT_DATE + INTERVAL '90 days',  'Viaje familiar',            true),
		('Compra Casa',     12000000,  50000000, 2000000, CURRENT_DATE + INTERVAL '730 days', 'Cuota inicial vivienda',    true),
		('Auto Nuevo',       8000000,  20000000, 1500000, CURRENT_DATE + INTERVAL '365 days', 'Compra de vehiculo',        true)`)
}

// Reverse the migrations
func (m *InsertTableEditarMeta_20260521_135310) Down() {
	m.SQL(`DELETE FROM "finanzas"."editar_meta" WHERE "nombre" IN ('Fondo Emergencia', 'Vacaciones Europa', 'Compra Casa', 'Auto Nuevo')`)
}
