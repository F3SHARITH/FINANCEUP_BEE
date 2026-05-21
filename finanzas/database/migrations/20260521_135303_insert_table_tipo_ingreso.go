package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTipoIngreso_20260521_135303 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTipoIngreso_20260521_135303{}
	m.Created = "20260521_135303"

	migration.Register("InsertTableTipoIngreso_20260521_135303", m)
}

// Run the migrations
func (m *InsertTableTipoIngreso_20260521_135303) Up() {
	m.SQL(`INSERT INTO "finanzas"."tipo_ingreso"
		("id_movimiento_dinero", "nombre_movimiento_pago", "descripcion", "activo")
	VALUES
		(1,    'Salario',      'Ingreso por salario mensual',              true),
		(3,    'Freelance',    'Ingresos por trabajo independiente',       true),
		(NULL, 'Bonificacion', 'Bonificacion laboral',                     true)`)
}

// Reverse the migrations
func (m *InsertTableTipoIngreso_20260521_135303) Down() {
	m.SQL(`DELETE FROM "finanzas"."tipo_ingreso" WHERE "nombre_movimiento_pago" IN ('Salario', 'Freelance', 'Bonificacion')`)
}
