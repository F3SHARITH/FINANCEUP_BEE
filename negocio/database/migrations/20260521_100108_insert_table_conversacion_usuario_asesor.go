package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableConversacionUsuarioAsesor_20260521_100108 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableConversacionUsuarioAsesor_20260521_100108{}
	m.Created = "20260521_100108"

	migration.Register("InsertTableConversacionUsuarioAsesor_20260521_100108", m)
}

// Run the migrations
func (m *InsertTableConversacionUsuarioAsesor_20260521_100108) Up() {
	m.SQL(`INSERT INTO "negocio"."conversacion_usuario_asesor"
		("id_lead", "id_usuario", "id_asesor", "tipo_contacto", "asunto", "contenido", "fecha_mensaje", "activo")
	VALUES
		(1, 1, 1, 'email', 'Solicitud de Credito Aprobada', 'Le informamos que su solicitud fue aprobada.', CURRENT_TIMESTAMP - INTERVAL '50 days', true),
		(1, 1, 1, 'whatsapp', 'Confirmacion de Desembolso', 'El dinero sera depositado en 2 dias habiles.', CURRENT_TIMESTAMP - INTERVAL '45 days', true),
		(2, 2, 2, 'telefono', 'Solicitud de Informacion', 'Llamada para recabar informacion sobre su propiedad.', CURRENT_TIMESTAMP - INTERVAL '25 days', true),
		(3, 3, 3, 'email', 'Plan de Acompanamiento', 'Iniciamos plan de acompanamiento para su negocio.', CURRENT_TIMESTAMP - INTERVAL '40 days', true),
		(4, 4, 4, 'whatsapp', 'Consulta sobre Vehiculo', 'Informacion sobre financiamiento de su vehiculo.', CURRENT_TIMESTAMP - INTERVAL '10 days', true)`)
}

// Reverse the migrations
func (m *InsertTableConversacionUsuarioAsesor_20260521_100108) Down() {
	m.SQL(`DELETE FROM "negocio"."conversacion_usuario_asesor"
	WHERE "asunto" IN (
		'Solicitud de Credito Aprobada',
		'Confirmacion de Desembolso',
		'Solicitud de Informacion',
		'Plan de Acompanamiento',
		'Consulta sobre Vehiculo'
	)`)
}
