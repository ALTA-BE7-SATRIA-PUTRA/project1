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
func Update(db *gorm.DB) {
	//tampilkan data user untuk memilih
	var users []_tables.User
	tx := db.Find(&users)
	if tx.Error != nil {
		// panic(tx.Error)
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
	fmt.Println("Masukkan ID User yang akan di Update:")
	fmt.Scanln(&user.ID)
	fmt.Print("Insert New Name:")
	fmt.Scanln(&newName)
	fmt.Print("Insert New Email:")
	fmt.Scanln(&newEmail)
	fmt.Print("Insert New Phone:")
	fmt.Scanln(&newPhone)
	db.Model(&user).Where("ID = ?", &user.ID).Updates(_tables.User{Name: newName, Email: newEmail, Phone: newPhone})

}

// func (db *DB) Select(query interface{}, args ...interface{}) (tx *DB)
func Delete(db *gorm.DB) []_tables.User {
	var users []_tables.User
	tx := db.Where("ID = ?", 3).Delete(&users)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("Deleted Failed ", tx.Error)
	}
	return users
}
func Transfer(db *gorm.DB) {

	//tampilkan data user untuk memilih
	var users []_tables.User
	tx := db.Find(&users)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}

	fmt.Println("ID\t Name ", "\t", "\t Email \t", "\t \t Phone Number", "\t \t Balance")
	for _, value := range users {

		fmt.Println(value.ID, "\t", value.Name, "\t", value.Email, "\t", value.Phone, "\t", value.Balance)
	}
	fmt.Println()
	//tf := _tables.Transfer{}
	var nominal, sendPhone, ReceiverPhone string
	fmt.Println("Include Phone Number sender :")
	fmt.Scanln(&sendPhone)
	fmt.Println("Include Receiver Phone Number :")
	fmt.Scanln(&ReceiverPhone)
	fmt.Print("Insert Nominal:")
	fmt.Scanln(&nominal)
	//db.Model(&tf).Where("Phone = ?", &tf.Phone).Updates("balance", &tf.Balance - nominal)
	//db.Model(&tf).Where("Phone = ?", &tf.Phone).Updates("balance", &tf.Balance + nominal)
	db.Exec("UPDATE users SET balance = ? WHERE phone = ?", gorm.Expr("balance - ?", nominal), sendPhone)
	db.Exec("UPDATE users SET balance = ? WHERE phone = ?", gorm.Expr("balance + ?", nominal), ReceiverPhone)
}
<<<<<<< HEAD
=======
func HistoryTopUp(db *gorm.DB) {

}
>>>>>>> f7925227318c25da731822c41d856bde8c8ed9c9

func TopUp(db *gorm.DB) {
	//tampilkan data user yang akan di transfer
	var users []_tables.User
<<<<<<< HEAD
	top_ups := _tables.TopUp{}
=======
>>>>>>> f7925227318c25da731822c41d856bde8c8ed9c9
	tx := db.Find(&users)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}

	fmt.Println("ID\t Name ", "\t", "\t", "\t Phone Number \t", "\t Balance")
	for _, value := range users {

		fmt.Println(value.ID, "\t", value.Name, "\t", "\t", value.Phone, "\t", value.Balance)
	}
	fmt.Println()

	//fungsi topup
<<<<<<< HEAD
	// var Phone, nominal string
	fmt.Print("Insert Phone Number to be TopUp: ")
	fmt.Scanln(&top_ups.Phone)
	fmt.Print("Input Nominal TopUp: ")
	fmt.Scanln(&top_ups.Nominal)
	db.Exec("UPDATE users SET balance = ? WHERE phone = ?", gorm.Expr("balance + ?", top_ups.Nominal), top_ups.Phone)

	//INSERT INTO TOPUP
	db.Select("phone", "nominal").Create(&top_ups)
}
func HistoryTopUp(db *gorm.DB) {

=======
	var Phone, nominal string
	fmt.Print("Insert Phone Number to be TopUp: ")
	fmt.Scanln(&Phone)
	fmt.Print("Input Nominal TopUp: ")
	fmt.Scanln(&nominal)
	db.Exec("UPDATE users SET balance = ? WHERE phone = ?", gorm.Expr("balance + ?", nominal), Phone)
	db.Exec("UPDATE topups SET balance = ? WHERE phone = ?", gorm.Expr("balance + ?", nominal), Phone)
>>>>>>> f7925227318c25da731822c41d856bde8c8ed9c9
}
