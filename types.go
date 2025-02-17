package main

type Account struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(id, firstName, lastName string, number, balance int64) *Account {
	return &Account{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Number:    number,
		Balance:   balance,
	}
}
