package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Userlogindetails struct {
	gorm.Model
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`

	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (user *Userlogindetails) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *Userlogindetails) CheckPassword(Password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if err != nil {
		return err
	}
	return nil
}
