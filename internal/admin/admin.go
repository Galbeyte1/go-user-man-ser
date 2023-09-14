package admin

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

type Admin struct {
	gorm.Model
	Username string
	Firstname string
	Passowrd string
	Email string
}

type AdminService interface{
	GetAdmin(ID uint) (Admin, error)
	PostUser(admin Admin) (Admin, error)
	GetAllUsers() ([]Admin, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

