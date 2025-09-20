package handlers

import (
	"context"
	"database/sql"
)

type Transaction struct {
	UserID    string
	Amount    float64
	Type      string
	Direction string
	City      string
	Country   string
	Latitude  float64
	Longitude float64
	DeviceID  string
	Timestamp string
}

type User struct {
	ID        string
	Name      string
	Profile   string
	Email     string
	Phone     string
	CreatedAt string
	Status    string
}

type TransactionRepository interface {
	Insert(ctx context.Context, tx *Transaction) error
	GetUserByID(ctx context.Context, userID string) (*User, error)
	GetTransactionsByUserID(ctx context.Context, userID string) ([]Transaction, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) GetAllUsers(ctx context.Context) ([]*User, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, profile, email, phone, created_at, status FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Profile, &u.Email, &u.Phone, &u.CreatedAt, &u.Status); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *transactionRepository) Insert(ctx context.Context, t *Transaction) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO transactions (user_id, amount, type, direction, city, country, latitude, longitude, device_id, timestamp) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
		t.UserID, t.Amount, t.Type, t.Direction, t.City, t.Country, t.Latitude, t.Longitude, t.DeviceID, t.Timestamp)
	return err
}

func (r *transactionRepository) GetUserByID(ctx context.Context, userID string) (*User, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, name, profile, email, phone, created_at, status FROM users WHERE id = $1`, userID)
	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Profile, &u.Email, &u.Phone, &u.CreatedAt, &u.Status)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *transactionRepository) GetTransactionsByUserID(ctx context.Context, userID string) ([]Transaction, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT user_id, amount, type, direction, city, country, latitude, longitude, device_id, timestamp FROM transactions WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var txs []Transaction
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.UserID, &t.Amount, &t.Type, &t.Direction, &t.City, &t.Country, &t.Latitude, &t.Longitude, &t.DeviceID, &t.Timestamp); err != nil {
			return nil, err
		}
		txs = append(txs, t)
	}
	return txs, nil
}
