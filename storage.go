package main

import (
	"database/sql"
	"fmt"
)

type Storage interface {
	CreateAccount(account *Account) error
	DeleteAccount(id int) error
	UpdateAccount(account *Account) error
	GetAccountById(id int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Init() error {
	return s.createAccountTable()

}

func (s *PostgresStorage) createAccountTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS accounts (
			id SERIAL PRIMARY KEY,
			first_name TEXT,
			last_name TEXT,
			number BIGINT,
			balance BIGINT,
			created_at TIMESTAMP DEFAULT NOW()
		)
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateAccount(account *Account) error {
	r, err := s.db.Exec(`
		INSERT INTO accounts (first_name, last_name, number, balance)
		VALUES ($1, $2, $3, $4)
	`, account.FirstName, account.LastName, account.Number, account.Balance)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", r)

	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	_, err := s.db.Query(`
		DELETE FROM accounts 
		WHERE id = $1
	`, id)
	return err
}

func (s *PostgresStorage) UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgresStorage) GetAccountById(id int) (*Account, error) {
	rows, err := s.db.Query(`
		SELECT *
		FROM accounts
		WHERE id = $1
	`, id)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account not found")
}

func (s *PostgresStorage) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query(`
		SELECT *
		FROM accounts
	`)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt)

	return account, err
}
