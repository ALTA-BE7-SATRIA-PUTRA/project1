package main

import (
	"fmt"
	//Import file lain dengan _<alias> <nama gomod>/<nama folder>
	_config "project1/config"
	_controllers "project1/controllers"
	_tables "project1/tables"

	"gorm.io/gorm"
)

var (
	connect *gorm.DB
)

func InitialMigration() {
	/*
		ketika tabel belum ada di mysql
		maka dibuat otomatis oleh gorm sesuai data pada struct
	*/

	connect.AutoMigrate(&_tables.User{})

}

func init() { // Menjalankan sebelum main
	connect = _config.InitDB() //panggil variabel koneksi yang ada di _namaFolder.func
	InitialMigration()
}

func main() {

	fmt.Println("Pilih Menu \n1 = Add User \n2 = Read User\n3 = Update")
	fmt.Println("4 = Delete User \n5 = Top-Up \n6 = Transfer \n7 = History Top-Up \n8 = History Transfer")
	fmt.Print("Masukan pilihan anda : ")
	var pilihan string
	fmt.Scanln(&pilihan)

	switch pilihan {
	case "1": // Add User
		newUser := _tables.User{}
		fmt.Println("Masukkan Nama:")
		fmt.Scanln(&newUser.Name)
		fmt.Println("Masukkan Email:")
		fmt.Scanln(&newUser.Email)
		fmt.Println("Masukkan Password:")
		fmt.Scanln(&newUser.Password)
		fmt.Println("Masukkan Phone:")
		fmt.Scanln(&newUser.Phone)

		tx := connect.Save(&newUser)
		if tx.Error != nil {
			// panic(tx.Error)
			fmt.Println("error when insert data")
		}
		if tx.RowsAffected == 0 {
			fmt.Println("insert failed")
		}
		fmt.Println("Insert successfully")

	case "2": // Read User
		users := _controllers.Read(connect)
		for _, value := range users {
			fmt.Println(value.ID, "-", value.Name)
		}

	case "3": // Update User
		_controllers.Update(connect)

	case "4": // Delete User
		err := _controllers.Delete(connect)
		if err != nil {
			fmt.Println("Deleted Failed")
		} else {
			fmt.Println("Delete Succes")
		}
		// case "5": // Top Up

		// case "6": // Transfer

		// case "7": // History Top Up

		// case "8": // History Transfer

	}
}
