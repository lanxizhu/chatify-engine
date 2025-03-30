package database

import "testing"

func TestDatabase(t *testing.T) {
	_, err := Create()
	if err != nil {
		t.Fatalf("Failed to create database: %v", err)
	}
	t.Log("Database connect successfully")
}
