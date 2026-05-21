package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertTableProductoCrediticio_20260520_173146 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertTableProductoCrediticio_20260520_173146{}
							m.Created = "20260520_173146"
							
							migration.Register("InsertTableProductoCrediticio_20260520_173146", m)
						}
					   
				// Run the migrations
				func (m *InsertTableProductoCrediticio_20260520_173146) Up() {
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertTableProductoCrediticio_20260520_173146) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					
				}
				