package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CrecionEsquemaAuth_20260520_165657 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrecionEsquemaAuth_20260520_165657{}
	m.Created = "20260520_165657"

	migration.Register("CrecionEsquemaAuth_20260520_165657", m)
}

// Run the migrations
func (m *CrecionEsquemaAuth_20260520_165657) Up() {
	file, err:=ioutil.ReadFile(".../scrips/20260520_165657_crecion_esquema_auth_Up.sql")
	if err !=nil{
		fmt.Println(err)
	}
	requests:=strings.Split(string(file),";")
	for _,request:=range requests{
		fmt.Println(request)
		m.SQL(request)
	}
	// use m.SQL("CREATE TABLE ...") to make schema update

	}


// Reverse the migrations
func (m *CrecionEsquemaAuth_20260520_165657) Down() {
	file, err:= ioutil.ReadFile("../scripts/20260520_165657_crecion_esquema_auth_Down.sq")
	
		if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}


}
