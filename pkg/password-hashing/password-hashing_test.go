package passwordHashing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "secret"

	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(hashedPassword)

	assert.NotNil(t, hashedPassword)
}

func TestVerifyPassword(t *testing.T) {
	password := "secret"

	hashed, hashErr := HashPassword(password)
	if hashErr != nil {
		t.Error(hashErr)
	}

	fmt.Println(hashed)

	isMatched := VerifyPassword(password, hashed)

	assert.Equal(t, true, isMatched)
}
