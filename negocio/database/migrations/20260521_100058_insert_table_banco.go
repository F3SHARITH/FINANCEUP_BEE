package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableBanco_20260521_100058 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableBanco_20260521_100058{}
	m.Created = "20260521_100058"

	migration.Register("InsertTableBanco_20260521_100058", m)
}

// Run the migrations
func (m *InsertTableBanco_20260521_100058) Up() {
	m.SQL(`INSERT INTO "negocio"."banco"
		("nombre_banco", "ciudad", "contacto", "telefono", "email", "comision_porcentaje", "url_logo", "descripcion", "sitio_web", "estado", "activo")
	VALUES
		('Banco Colombiano', 'Bogota', 'Juan Gomez', '6015551234', 'contacto@bancocol.com', 2.50, 'https://example.com/logo1.png', 'Banco con amplia cobertura nacional', 'www.bancocol.com', 'activo', true),
		('Banco Metropolitano', 'Medellin', 'Maria Lopez', '5745551234', 'contacto@bancometro.com', 2.75, 'https://example.com/logo2.png', 'Banco especializado en creditos', 'www.bancometro.com', 'activo', true),
		('Banco del Oriente', 'Bucaramanga', 'Carlos Ruiz', '7685551234', 'contacto@bancoriente.com', 2.25, 'https://example.com/logo3.png', 'Banco regional con buenos servicios', 'www.bancoriente.com', 'activo', true),
		('Banco Pacifico', 'Cali', 'Pedro Diaz', '3105551234', 'contacto@bancopac.com', 2.60, 'https://example.com/logo4.png', 'Banco del sector occidental', 'www.bancopac.com', 'activo', true)`)
}

// Reverse the migrations
func (m *InsertTableBanco_20260521_100058) Down() {
	m.SQL(`DELETE FROM "negocio"."banco"
	WHERE "email" IN (
		'contacto@bancocol.com',
		'contacto@bancometro.com',
		'contacto@bancoriente.com',
		'contacto@bancopac.com'
	)`)
}
