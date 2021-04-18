package usermanager

import (
	"encoding/json"
	"errors"

	simplehash "github.com/fulgurant/simplehash"
)

// Errors
var (
	ErrNoUser     = errors.New("no user")
	ErrNoEmail    = errors.New("no email address")
	ErrNoPassword = errors.New("no password")
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Approved bool   `json:"approved"`
}

func (u *User) Check() error {
	if u == nil {
		return ErrNoUser
	}
	if u.Email == "" {
		return ErrNoEmail
	}
	if u.Password == "" {
		return ErrNoPassword
	}
	return nil
}

func (u *User) Hash(hs simplehash.Hasher) error {
	u.Email = hs.Hash(u.Email)
	u.Password = hs.Hash(u.Password)
	return nil
}

func (u *User) ToBytes() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) FromBytes(value []byte) error {
	return json.Unmarshal(value, u)
}
