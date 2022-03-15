package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

/* membuat tipe data struct disini
sesuaikan dengan kebutuhan tabel mysql*/

func InitDB() {
	// Membuat koneksi ke db
	// Isi sesuai db yang digunakan (jangan ada spasi dalam "")
	connectionString := "root: <Pasword> @tcp(127.0.0.1:3306)/ <nama db> ?charset=utf8&parseTime=True&loc=Local"

	var err error
	// gorm.Open mereturn 2 value dan di tampung ke DB, err
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

// func InitialMigration() {
// 	/*
// 	ketika tabel belum ada di mysql
// 	maka dibuat otomatis oleh gorm sesuai data pada struct
// 	*/
// 	DB.AutoMigrate(&/*Sesuaikan dengan nama struct*/{})

// }

// func init() { // Menjalankan sebelum
// 	InitDB()
// 	InitialMigration()
// }

func main() {

	fmt.Println("Pilih Menu \n1 = Add User \n2 = Read User\n3 = Update")
	fmt.Println("4 = Delete User \n5 = Top-Up \n6 = Transfer \n7 = History Top-Up \n8 = History Transfer")
	fmt.Print("Masukan pilihan anda : ")
	var pilihan string
	fmt.Scanln(&pilihan)

	// switch pilihan {
	// case "1": // Add User

	// case "2": // Read User

	// case "3": // Update User

	// case "4": // Delete User

	// case "5": // Top Up

	// case "6": // Transfer

	// case "7": // History Top Up

	// case "8": // History Transfer

	// }
}
