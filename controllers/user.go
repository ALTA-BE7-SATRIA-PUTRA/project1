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
		fmt.Println("error ", tx.Error)
	}
	return users
}

func SelectUser(db *gorm.DB) []_tables.User {
	var users []_tables.User
	tx := db.Select(&users)
	if tx.Error != nil {
		fmt.Println("error ", tx.Error)
	}
	return users
}
func Update(db *gorm.DB) {
	//tampilkan data user untuk memilih
	var users []_tables.User
	tx := db.Find(&users)
	if tx.Error != nil {
		fmt.Println("error ", tx.Error)
	}

	fmt.Println("ID\t Name ", "\t", "\t Email \t", "\t \t Phone Number")
	for _, value := range users {

		fmt.Println(value.ID, "\t", value.Name, "\t", value.Email, "\t", value.Phone)
	}
	fmt.Println()

	//fungsi untuk update
	user := _tables.User{}
	var newName, newEmail, newPhone string
	fmt.Print("Select the id to be updated: ")
	fmt.Scanln(&user.ID)
	fmt.Print("Insert New Name:")
	fmt.Scanln(&newName)
	fmt.Print("Insert New Email:")
	fmt.Scanln(&newEmail)
	fmt.Print("Insert New Phone:")
	fmt.Scanln(&newPhone)
	db.Model(&user).Where("ID = ?", &user.ID).Updates(_tables.User{Name: newName, Email: newEmail, Phone: newPhone})
}

func Delete(db *gorm.DB) []_tables.User {
	var users []_tables.User
	var delete string
	fmt.Println(users)

	fmt.Print("select the id to be delete : ")
	fmt.Scanln(&delete)
	tx := db.Where("ID = ?", delete).Delete(&users)
	if tx.Error != nil {
		fmt.Println("Deleted Failed ", tx.Error)
	}
	return users
}
func Transfer(db *gorm.DB) {

	//tampilkan data user untuk memilih
	var users []_tables.User
	tx := db.Find(&users)
	transfers := _tables.Transfer{}
	if tx.Error != nil {
		fmt.Println("error ", tx.Error)
	}

	fmt.Println("ID\t Name ", "\t", "\t Email \t", "\t \t Phone Number", "\t Balance")
	for _, value := range users {

		fmt.Println(value.ID, "\t", value.Name, "\t", value.Email, "\t", value.Phone, "\t", value.Balance)
	}
	fmt.Println()
	fmt.Println("Input Phone Number sender :")
	fmt.Scanln(&transfers.Phone)
	fmt.Println("Input Receiver Phone Number :")
	fmt.Scanln(&transfers.Receiver)
	fmt.Print("Input Nominal:")
	fmt.Scanln(&transfers.Nominal)
	//query untuk Database
	db.Exec("UPDATE users SET balance = ? WHERE phone = ?", gorm.Expr("balance - ?", transfers.Nominal), transfers.Phone)
	db.Exec("UPDATE users SET balance = ? WHERE phone = ?", gorm.Expr("balance + ?", transfers.Nominal), transfers.Receiver)
	db.Select("phone", "nominal", "receiver").Create(&transfers)
}

func TopUp(db *gorm.DB) {
	//tampilkan data user yang akan di transfer
	var users []_tables.User
	top_ups := _tables.TopUp{}
	tx := db.Find(&users)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}

	fmt.Println("ID\t Name ", "\t", "\t", "\t Phone Number \t", "Balance")
	for _, value := range users {

		fmt.Println(value.ID, "\t", value.Name, "\t", "\t", value.Phone, "\t", value.Balance)
	}
	fmt.Println()

	//fungsi topup
	fmt.Print("Insert Phone Number to be TopUp: ")
	fmt.Scanln(&top_ups.Phone)
	fmt.Print("Input Nominal TopUp: ")
	fmt.Scanln(&top_ups.Nominal)
	db.Exec("UPDATE users SET balance = ? WHERE phone = ?", gorm.Expr("balance + ?", top_ups.Nominal), top_ups.Phone)

	//INSERT INTO TOPUP
	db.Select("phone", "nominal").Create(&top_ups)
}
func HistoryTopUp(db *gorm.DB) {
	//tampilkan data user yang akan di transfer
	var topup []_tables.TopUp
	tx := db.Find(&topup)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}

	fmt.Println("ID\t ", "Nominal", "\tPhone Number \t", "\t time")
	for _, value := range topup {

		fmt.Println(value.ID, "\t", value.Nominal, "\t", "\t", value.Phone, "\t", value.CreatedAt)
	}
	fmt.Println()
}

func HistoryTransfer(db *gorm.DB) {
	var transfer []_tables.Transfer
	tx := db.Find(&transfer)
	if tx.Error != nil {
		fmt.Println("error ", tx.Error)
	}

	fmt.Println("ID\t ", "Nominal", "\tPhone Sender \t Phone Reciver \t", "\t time")

	for _, value := range transfer {
		fmt.Println(value.ID, "\t", value.Nominal, "\t", "\t", value.Phone, "\t", value.Receiver, "\t", value.CreatedAt)
	}

}
