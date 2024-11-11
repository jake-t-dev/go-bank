package main

import "math/rand"

type Account struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Number    int64   `json:"number"`
	Balance   float64 `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        0,
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63(),
	}
}
