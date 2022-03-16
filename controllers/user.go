package controllers

import (
	"fmt"
	_tables "project1/tables"

	"gorm.io/gorm"
)

func Read(db *gorm.DB) []_tables.User {
	var users []_tables.User
	tx := db.Find(&users)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return users
}
func Delete(db *gorm.DB) []_tables.User {
	var users []_tables.User
	tx := db.Where("ID = ?", 3).Delete(&users)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("Deleted Failed ", tx.Error)
	}
	return users
}
