package main

import (
	"fmt"
	"io/ioutil"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Finanzas_20260521_135251 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Finanzas_20260521_135251{}
	m.Created = "20260521_135251"

	migration.Register("Finanzas_20260521_135251", m)
}

// Run the migrations
func (m *Finanzas_20260521_135251) Up() {
	file, err := ioutil.ReadFile("../scripts/20260521_135251_finanzas_up.sql")

	if err != nil {
		fmt.Println("Error al momento de leer el archivo SQL:", err)
		return
	}

	m.SQL(string(file))
}

// Reverse the migrations
func (m *Finanzas_20260521_135251) Down() {
	file, err := ioutil.ReadFile("../scripts/20260521_135251_finanzas_down.sql")

	if err != nil {
		fmt.Println("Error al momento de leer el archivo SQL:", err)
		return
	}

	m.SQL(string(file))
}
