package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertTableBanco_20260520_172947 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertTableBanco_20260520_172947{}
							m.Created = "20260520_172947"
							
							migration.Register("InsertTableBanco_20260520_172947", m)
						}
					   
				// Run the migrations
				func (m *InsertTableBanco_20260520_172947) Up() {
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertTableBanco_20260520_172947) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					
				}
				