package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionEsquemaNegocio_20260520_161322 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionEsquemaNegocio_20260520_161322{}
	m.Created = "20260520_161322"

	migration.Register("CreacionEsquemaNegocio_20260520_161322", m)
}

// Run the migrations
func (m *CreacionEsquemaNegocio_20260520_161322) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreacionEsquemaNegocio_20260520_161322) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
