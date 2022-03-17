package tables

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `gorm:"unique" gorm:"primarykey" json:"phone" form:"phone"`
	Balance  string `gorm:"default:0" json:"balance" form:"balance"`
}
type TopUp struct {
	gorm.Model
<<<<<<< HEAD
	Phone   string `json:"phone" form:"phone"`
	Nominal string `json:"nominal" form:"nominal"`
	// Balance string `gorm:"default:0" json:"balance" form:"balance"`
	// Phone_User string `gorm:"foreignKey:Phone" json:"name" form:"name"`
=======
	Phones     string `json:"phone" form:"phone"`
	Nominal    string `json:"nominal" form:"nominal"`
	Balance    string `json:"balance" form:"balance"`
	Phone_User string `gorm:"foreignKey:Phone" json:"name" form:"name"`
>>>>>>> f7925227318c25da731822c41d856bde8c8ed9c9
}
type Transfer struct {
	gorm.Model
	Phone      string `json:"phone" form:"phone"`
	Nominal    string `json:"nominal" form:"nominal"`
	Balance    string `json:"balance" form:"balance"`
	Phone_User string `gorm:"foreignKey:Phone" json:"name" form:"name"`
}
