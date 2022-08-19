package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashPwd, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPwd)

	err = CheckPassword(password, hashPwd)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashPwd)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
