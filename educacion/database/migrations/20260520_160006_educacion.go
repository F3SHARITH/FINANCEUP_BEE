package main

import (
	"fmt"
	"io/ioutil"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Educacion_20260520_160006 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Educacion_20260520_160006{}
	m.Created = "20260520_160006"

	migration.Register("Educacion_20260520_160006", m)
}

// Run the migrations
func (m *Educacion_20260520_160006) Up() {
	file, err := ioutil.ReadFile("../scripts/20260520_160006_educacion_up.sql")

	if err != nil {
		fmt.Println("Error al momento de leer el archivo SQL:", err)
		return
	}

	m.SQL(string(file))
}

// Reverse the migrations
func (m *Educacion_20260520_160006) Down() {
	file, err := ioutil.ReadFile("../scripts/20260520_160006_educacion_down.sql")

	if err != nil {
		fmt.Println("Error al momento de leer el archivo SQL:", err)
		return
	}

	m.SQL(string(file))
}
