package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableAsesorBancario_20260521_100102 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableAsesorBancario_20260521_100102{}
	m.Created = "20260521_100102"

	migration.Register("InsertTableAsesorBancario_20260521_100102", m)
}

// Run the migrations
func (m *InsertTableAsesorBancario_20260521_100102) Up() {
	m.SQL(`INSERT INTO "negocio"."asesor_bancario"
		("id_banco", "nombre", "apellido", "email", "telefono", "especialidad", "activo")
	VALUES
		(1, 'Roberto', 'Sanchez', 'roberto.sanchez@bancocol.com', '3001112222', 'Creditos Personales', true),
		(1, 'Diana', 'Valenzuela', 'diana.valenzuela@bancocol.com', '3009998888', 'Hipotecarios', true),
		(2, 'Fernando', 'Castillo', 'fernando.castillo@bancometro.com', '3107776666', 'Microcreditos', true),
		(3, 'Claudia', 'Morales', 'claudia.morales@bancoriente.com', '3104445555', 'Vehiculos', true)`)
}

// Reverse the migrations
func (m *InsertTableAsesorBancario_20260521_100102) Down() {
	m.SQL(`DELETE FROM "negocio"."asesor_bancario"
	WHERE "email" IN (
		'roberto.sanchez@bancocol.com',
		'diana.valenzuela@bancocol.com',
		'fernando.castillo@bancometro.com',
		'claudia.morales@bancoriente.com'
	)`)
}
