package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertTableConversacionUsuarioAsesor_20260520_173030 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertTableConversacionUsuarioAsesor_20260520_173030{}
							m.Created = "20260520_173030"
							
							migration.Register("InsertTableConversacionUsuarioAsesor_20260520_173030", m)
						}
					   
				// Run the migrations
				func (m *InsertTableConversacionUsuarioAsesor_20260520_173030) Up() {
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertTableConversacionUsuarioAsesor_20260520_173030) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					
				}
				