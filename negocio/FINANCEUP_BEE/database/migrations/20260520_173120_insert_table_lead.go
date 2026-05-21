package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertTableLead_20260520_173120 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertTableLead_20260520_173120{}
							m.Created = "20260520_173120"
							
							migration.Register("InsertTableLead_20260520_173120", m)
						}
					   
				// Run the migrations
				func (m *InsertTableLead_20260520_173120) Up() {
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertTableLead_20260520_173120) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					
				}
				