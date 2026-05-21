package main

import (
	"fmt"
	"io/ioutil"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Negocio_20260521_095957 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Negocio_20260521_095957{}
	m.Created = "20260521_095957"

	migration.Register("Negocio_20260521_095957", m)
}

// Run the migrations
func (m *Negocio_20260521_095957) Up() {
	file, err := ioutil.ReadFile("../scripts/20260521_095957_negocio_up.sql")

	if err != nil {
		fmt.Println("Error al momento de leer el archivo SQL:", err)
		return
	}

	m.SQL(string(file))
}

// Reverse the migrations
func (m *Negocio_20260521_095957) Down() {
	file, err := ioutil.ReadFile("../scripts/20260521_095957_negocio_down.sql")

	if err != nil {
		fmt.Println("Error al momento de leer el archivo SQL:", err)
		return
	}

	m.SQL(string(file))
}
