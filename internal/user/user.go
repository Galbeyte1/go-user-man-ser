package comment

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	Username string
	Firstname string
	Password string
	Email string
}

type UserService interface{
	GetUser(ID uint) (User, error)
	PostUser(user User) (User, error)
	GetAllUsers() ([]User, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetUser(ID uint) (User, error) {
	var user User
	if result := s.DB.First(&user, ID); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (s *Service) PostUser(user User) (User, error) {
	if result := s.DB.Save(&user); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (s *Service) GetAllUsers() ([]User, error) {
	var users []User
	if result := s.DB.Find(&users); result.Error != nil {
		return users, result.Error
	}
	return users, nil
}