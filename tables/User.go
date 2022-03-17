package tables

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `gorm:"unique" json:"phone" form:"phone"`
	Balance  string `json:"balance" form:"balance"`
}
type TopUp struct {
	gorm.Model
	Phone   string `json:"phone" form:"phone"`
	Nominal string `json:"nominal" form:"nominal"`
	Balance string `json:"balance" form:"balance"`
	ID_User string `gorm:"foreignkey" json:"name" form:"name"`
}
type Transfer struct {
	gorm.Model
	Phone   string `json:"phone" form:"phone"`
	Nominal string `json:"nominal" form:"nominal"`
	Balance string `json:"balance" form:"balance"`
}
