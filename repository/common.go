package repository

import "golang.org/x/crypto/bcrypt"

func hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func validPassword(hashedPassword, planePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(planePassword))
}
