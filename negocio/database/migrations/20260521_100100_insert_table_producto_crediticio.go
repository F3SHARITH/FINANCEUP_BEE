package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableProductoCrediticio_20260521_100100 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableProductoCrediticio_20260521_100100{}
	m.Created = "20260521_100100"

	migration.Register("InsertTableProductoCrediticio_20260521_100100", m)
}

// Run the migrations
func (m *InsertTableProductoCrediticio_20260521_100100) Up() {
	m.SQL(`INSERT INTO "negocio"."producto_crediticio"
		("id_banco", "nombre_producto", "descripcion", "monto_minimo", "monto_maximo", "tasa_minima", "tasa_maxima", "plazo_minimo", "plazo_maximo", "requisitos", "activo")
	VALUES
		(1, 'Credito Personal', 'Credito para gastos personales', 1000000, 50000000, 10.00, 20.00, 12, 84, 'Cedula, comprobante ingresos', true),
		(1, 'Credito Hipotecario', 'Credito para compra de vivienda', 100000000, 1000000000, 8.00, 12.00, 120, 360, 'Cedula, avaluo, comprobante ingresos', true),
		(2, 'Microcredito', 'Credito para pequenos negocios', 500000, 10000000, 15.00, 25.00, 6, 60, 'Cedula, plan de negocio', true),
		(3, 'Credito de Vehiculo', 'Financiamiento de vehiculos', 10000000, 150000000, 9.00, 18.00, 24, 84, 'Cedula, documento vehiculo', true)`)
}

// Reverse the migrations
func (m *InsertTableProductoCrediticio_20260521_100100) Down() {
	m.SQL(`DELETE FROM "negocio"."producto_crediticio"
	WHERE "nombre_producto" IN (
		'Credito Personal',
		'Credito Hipotecario',
		'Microcredito',
		'Credito de Vehiculo'
	)`)
}
