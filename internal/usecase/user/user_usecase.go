package usecase

import (
	"errors"
	"go-task-manager/internal/domain"
	"go-task-manager/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) Register(name, email, password string) error {
	// 1. Email kontrolü (Var mı?)
	if _, err := u.repo.GetByEmail(email); err == nil {
		return errors.New("email already exists")
	}

	// 2. Şifre hashleme
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. Kullanıcı oluşturma
	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return u.repo.Create(user)
}

func (u *userUseCase) Login(email, password string) (string, error) {
	// 1. Kullanıcıyı bul
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// 2. Şifreyi doğrula
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}


	// 3. Token oluştur
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
