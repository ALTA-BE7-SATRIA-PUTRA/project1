package main

import (
	"fmt"
	//Import file lain dengan _<alias> <nama gomod>/<nama folder>
	_config "project1/config"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

/* membuat tipe data struct disini
sesuaikan dengan kebutuhan tabel mysql*/

func InitialMigration() {
	/*
		ketika tabel belum ada di mysql
		maka dibuat otomatis oleh gorm sesuai data pada struct
	*/

	// buka komentar ini <-- DB.AutoMigrate(&/*Sesuaikan dengan nama struct*/{})

}

func init() { // Menjalankan sebelum main
	DB = _config.InitDB //panggil variabel koneksi yang ada di _namaFolder.func
	InitialMigration()
}

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
