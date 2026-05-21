package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTipoIngresoMeta_20260521_135312 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTipoIngresoMeta_20260521_135312{}
	m.Created = "20260521_135312"

	migration.Register("InsertTableTipoIngresoMeta_20260521_135312", m)
}

// Run the migrations
func (m *InsertTableTipoIngresoMeta_20260521_135312) Up() {
	m.SQL(`INSERT INTO "finanzas"."tipo_ingreso_meta"
		("id_movimiento_dinero", "nombre_movimiento_pago", "descripcion", "activo")
	VALUES
		(1, 'Aporte Regular',  'Aporte regular a la meta',    true),
		(2, 'Aporte Especial', 'Aporte adicional a la meta',  true)`)
}

// Reverse the migrations
func (m *InsertTableTipoIngresoMeta_20260521_135312) Down() {
	m.SQL(`DELETE FROM "finanzas"."tipo_ingreso_meta" WHERE "nombre_movimiento_pago" IN ('Aporte Regular', 'Aporte Especial')`)
}
