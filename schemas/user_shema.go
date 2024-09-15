package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        int
	Name      string
	Email     string
	Password  string
	CompanyId int
	Company   Company `gorm:"foreignKey:CompanyId"`
}
