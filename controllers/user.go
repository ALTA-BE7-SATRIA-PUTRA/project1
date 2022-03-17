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

func SelectUser(db *gorm.DB) []_tables.User {
	var users []_tables.User
	tx := db.Select(&users)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return users
}

// func (db *DB) Select(query interface{}, args ...interface{}) (tx *DB)
