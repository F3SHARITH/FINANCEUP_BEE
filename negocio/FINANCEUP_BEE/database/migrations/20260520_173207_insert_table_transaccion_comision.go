package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertTableTransaccionComision_20260520_173207 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertTableTransaccionComision_20260520_173207{}
							m.Created = "20260520_173207"
							
							migration.Register("InsertTableTransaccionComision_20260520_173207", m)
						}
					   
				// Run the migrations
				func (m *InsertTableTransaccionComision_20260520_173207) Up() {
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertTableTransaccionComision_20260520_173207) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					
				}
				