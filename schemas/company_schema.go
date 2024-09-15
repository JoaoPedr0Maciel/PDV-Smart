package schemas

type Company struct {
	Id    int
	Name  string
	Phone string
	Email string
	Password string
	Users   []User `gorm:"foreignKey:CompanyId"`
}
