package types

import (
	"math/rand/v2"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	Number   int64  `json:"number`
	Password string `json:"password`
}

type TransferRequest struct {
	ToAccount int `json :"toAccount"`
	Amount    int `json :"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName`
	LastName  string `json:"lastName`
	Password  string `json:"password"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt`
}

func (a *Account) ValidatePassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

func NewAccount(firstName, lastName string, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:                int(rand.Int64N(1000)),
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encpw),
		Number:            rand.Int64N(100000),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
