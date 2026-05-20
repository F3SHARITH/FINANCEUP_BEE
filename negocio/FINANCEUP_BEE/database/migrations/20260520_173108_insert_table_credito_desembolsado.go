package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertTableCreditoDesembolsado_20260520_173108 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertTableCreditoDesembolsado_20260520_173108{}
							m.Created = "20260520_173108"
							
							migration.Register("InsertTableCreditoDesembolsado_20260520_173108", m)
						}
					   
				// Run the migrations
				func (m *InsertTableCreditoDesembolsado_20260520_173108) Up() {
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertTableCreditoDesembolsado_20260520_173108) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					
				}
				