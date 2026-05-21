package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertTableAsesorBancario_20260520_172658 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertTableAsesorBancario_20260520_172658{}
							m.Created = "20260520_172658"
							
							migration.Register("InsertTableAsesorBancario_20260520_172658", m)
						}
					   
				// Run the migrations
				func (m *InsertTableAsesorBancario_20260520_172658) Up() {
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertTableAsesorBancario_20260520_172658) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					
				}
				