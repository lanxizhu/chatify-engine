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

func TestVerifyPassword(t *testing.T) {
	input := "hello"

	var user = &User{
		Username: "admin",
		Password: input,
	}

	_ = user.HashPassword()
	if !user.VerifyPassword(input) {
		t.Errorf("CheckPassword() got = %v, want %v", user.VerifyPassword(input), true)
	}
	t.Log("User CheckPassword() passed")
}
