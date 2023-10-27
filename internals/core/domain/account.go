package domain

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"password"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (a *Account) ValidPassword(pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw))

	return err == nil
}

func NewAccount(firstName, lastName, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encpw),
		Number:            int64(rand.Intn(10000)),
		Balance:           0,
		CreatedAt:         time.Now().UTC(),
	}, nil
}
