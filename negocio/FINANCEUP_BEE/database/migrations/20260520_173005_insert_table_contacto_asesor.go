package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertTableContactoAsesor_20260520_173005 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertTableContactoAsesor_20260520_173005{}
							m.Created = "20260520_173005"
							
							migration.Register("InsertTableContactoAsesor_20260520_173005", m)
						}
					   
				// Run the migrations
				func (m *InsertTableContactoAsesor_20260520_173005) Up() {
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertTableContactoAsesor_20260520_173005) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					
				}
				