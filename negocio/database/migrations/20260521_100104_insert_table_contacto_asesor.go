package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableContactoAsesor_20260521_100104 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableContactoAsesor_20260521_100104{}
	m.Created = "20260521_100104"

	migration.Register("InsertTableContactoAsesor_20260521_100104", m)
}

// Run the migrations
func (m *InsertTableContactoAsesor_20260521_100104) Up() {
	m.SQL(`INSERT INTO "negocio"."contacto_asesor"
		("id_asesor", "whatsapp", "email", "telefono", "disponible_desde", "disponible_hasta", "dias_disponibles", "activo")
	VALUES
		(1, '3001112222', 'roberto.sanchez@bancocol.com', '3001112222', '08:00', '18:00', 'Lunes a Viernes', true),
		(2, '3009998888', 'diana.valenzuela@bancocol.com', '3009998888', '09:00', '17:00', 'Lunes a Viernes', true),
		(3, '3107776666', 'fernando.castillo@bancometro.com', '3107776666', '08:00', '20:00', 'Lunes a Sabado', true),
		(4, '3104445555', 'claudia.morales@bancoriente.com', '3104445555', '07:00', '19:00', 'Lunes a Viernes', true)`)
}

// Reverse the migrations
func (m *InsertTableContactoAsesor_20260521_100104) Down() {
	m.SQL(`DELETE FROM "negocio"."contacto_asesor"
	WHERE "email" IN (
		'roberto.sanchez@bancocol.com',
		'diana.valenzuela@bancocol.com',
		'fernando.castillo@bancometro.com',
		'claudia.morales@bancoriente.com'
	)`)
}
