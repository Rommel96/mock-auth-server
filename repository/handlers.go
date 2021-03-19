package repository

import "log"

func (user *User) Insert() {
	encrypted, _ := hash(user.Password)
	user.Password = string(encrypted)
	err := db.Create(user).Error
	if err != nil {
		log.Println(err)
	}
}

func FindUser(email, password string) *User {
	var user User
	err := db.First(&user, "email=?", email).Error
	if err != nil {
		return nil
	}
	err = validPassword(user.Password, password)
	if err != nil {
		return nil
	}
	return &user
}
