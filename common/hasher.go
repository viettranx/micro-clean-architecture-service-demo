package common

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct{}

func (r *Hasher) RandomStr(length int) (string, error) {
	var b = make([]byte, length)

	_, err := rand.Read(b)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return hex.EncodeToString(b), nil
}

func (r *Hasher) HashPassword(salt, password string) (string, error) {
	spStr := fmt.Sprintf("%s.%s", salt, password)

	h, err := bcrypt.GenerateFromPassword([]byte(spStr), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(h), nil
}

func (r *Hasher) CompareHashPassword(hashedPassword, salt, password string) bool {
	spStr := fmt.Sprintf("%s.%s", salt, password)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(spStr)) == nil
}
