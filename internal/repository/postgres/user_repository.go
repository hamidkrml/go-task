package repository

import (
	"database/sql"
	"go-task-manager/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}
