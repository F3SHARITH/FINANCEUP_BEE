package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionSquemaSoporte_20260520_170627 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionSquemaSoporte_20260520_170627{}
	m.Created = "20260520_170627"

	migration.Register("CreacionSquemaSoporte_20260520_170627", m)
}

// Run the migrations
func (m *CreacionSquemaSoporte_20260520_170627) Up() {
	file, err:= ioutil.ReadFile("../scripts/20260520_170627_creacion_squema_soporte_up.sql")


	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *CreacionSquemaSoporte_20260520_170627) Down() {
	file, err :=ioutil.ReadFile("../scripts/20260520_170627_creacion_squema_soporte_down.sql")


	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}