package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableLead_20260521_100106 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableLead_20260521_100106{}
	m.Created = "20260521_100106"

	migration.Register("InsertTableLead_20260521_100106", m)
}

// Run the migrations
func (m *InsertTableLead_20260521_100106) Up() {
	m.SQL(`INSERT INTO "negocio"."lead"
		("id_usuario", "id_producto", "id_asesor", "tipo_credito", "monto_interes", "plazo_interes", "estado_lead", "fecha_generacion", "fecha_contacto", "observaciones", "activo")
	VALUES
		(1, 1, 1, 'Personal', 5000000, 36, 'aprobado', CURRENT_TIMESTAMP - INTERVAL '60 days', CURRENT_TIMESTAMP - INTERVAL '50 days', 'Cliente solvente, aprobado sin inconvenientes', true),
		(2, 2, 2, 'Hipotecario', 200000000, 240, 'en_proceso', CURRENT_TIMESTAMP - INTERVAL '30 days', CURRENT_TIMESTAMP - INTERVAL '25 days', 'Pendiente evaluacion de inmueble', true),
		(3, 3, 3, 'Microcredito', 5000000, 36, 'aprobado', CURRENT_TIMESTAMP - INTERVAL '45 days', CURRENT_TIMESTAMP - INTERVAL '40 days', 'Negocio informal, requiere seguimiento', true),
		(4, 4, 4, 'Vehiculo', 50000000, 60, 'contactado', CURRENT_TIMESTAMP - INTERVAL '15 days', CURRENT_TIMESTAMP - INTERVAL '10 days', 'Interesado en financiamiento de camioneta', true),
		(5, 1, 1, 'Personal', 3000000, 24, 'nuevo', CURRENT_TIMESTAMP - INTERVAL '5 days', NULL, 'Lead recien generado', true)`)
}

// Reverse the migrations
func (m *InsertTableLead_20260521_100106) Down() {
	m.SQL(`DELETE FROM "negocio"."lead"
	WHERE ("id_usuario", "id_producto", "tipo_credito") IN (
		(1, 1, 'Personal'),
		(2, 2, 'Hipotecario'),
		(3, 3, 'Microcredito'),
		(4, 4, 'Vehiculo'),
		(5, 1, 'Personal')
	)`)
}
