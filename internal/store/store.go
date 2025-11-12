package store

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Store struct {
	db *pgxpool.Pool
}

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(ctx context.Context, name, email, password string, role int32) (int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	var id int64
	err = s.db.QueryRow(ctx,
		`INSERT INTO users(name, email, password, role) VALUES($1, $2, $3, $4) RETURNING id`,
		name, email, string(hash), role,
	).Scan(&id)
	return id, err
}

func (s *Store) GetUser(ctx context.Context, id int64) (*User, error) {
	var u User
	err := s.db.QueryRow(ctx,
		`SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id=$1`, id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *Store) UpdateUser(ctx context.Context, id int64, name, email *string, role int32) error {
	query := `UPDATE users SET updated_at = NOW()`
	args := []interface{}{}
	argID := 1

	if name != nil {
		query += fmt.Sprintf(", name = $%d", argID)
		args = append(args, *name)
		argID++
	}
	if email != nil {
		query += fmt.Sprintf(", email = $%d", argID)
		args = append(args, *email)
		argID++
	}
	if role > 0 {
		query += fmt.Sprintf(", role = $%d", argID)
		args = append(args, role)
		argID++
	}

	query += fmt.Sprintf(" WHERE id = $%d", argID)
	args = append(args, id)

	_, err := s.db.Exec(ctx, query, args...)
	return err
}

func (s *Store) DeleteUser(ctx context.Context, id int64) error {
	_, err := s.db.Exec(ctx, `DELETE FROM users WHERE id=$1`, id)
	return err
}
