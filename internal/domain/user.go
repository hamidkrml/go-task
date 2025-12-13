package domain

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"` // JSON'a çevrilirken şifre gitmesin
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
}

type UserUseCase interface {
	Register(name, email, password string) error
	Login(email, password string) (string, error) // Token dönecek
}
