package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertarTablaAdjunto_20260520_171927 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTablaAdjunto_20260520_171927{}
	m.Created = "20260520_171927"

	migration.Register("InsertarTablaAdjunto_20260520_171927", m)
}

// Run the migrations
func (m *InsertarTablaAdjunto_20260520_171927) Up() {
	m.SQL("INSERT INTO soporte.adjunto (id_pqr, nombre_archivo, ruta_archivo, tipo_mime, tamano_bytes, activo) VALUES (1, 'error_screenshot.png', '/uploads/pqr/error_screenshot.png', 'image/png', 512000, true)") 

}

// Reverse the migrations
func (m *InsertarTablaAdjunto_20260520_171927) Down() {
	m.SQL("DELETE FROM soporte.adjunto WHERE id_pqr = 1 AND nombre_archivo = 'error_screenshot.png'")

}
