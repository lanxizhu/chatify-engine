package model

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	input := "hello"
	expectedHash := "5d41402abc4b2a76b9719d911017c592" // MD5 hash of "hello"

	var user = &User{
		Username: "admin",
		Password: input,
	}

	err := user.HashPassword()
	if err != nil {
		t.Errorf("HashPassword() error = %v", err)
	}

	if user.Password == expectedHash {
		t.Errorf("HashPassword() got = %v, want %v", user.Password, expectedHash)
	}
	t.Log("User HashPassword() passed")
}
