package usermanager

import (
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
	Name     string
	Email    string
	Password string
	Approved bool
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
