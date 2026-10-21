package services

import (
	"errors"
	"go-shortener/models"
	"go-shortener/repositories"
	"go-shortener/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) SignUp(name, email, password string) (*models.User, error) {
	name = strings.TrimSpace(name)
	email = strings.ToLower(strings.TrimSpace(email))

	if name == "" || email == "" || password == "" {
		return nil, errors.New("All fields are required.")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		if errors.Is(err, utils.ErrEmailExists) {
			return nil, utils.ErrEmailExists
		}
		return nil, err
	}

	return user, nil

}

func (s *AuthService) SignIn(email, password string) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("Invalid email or password")
	}

	token, err := utils.CreateToken(user.ID)

	if err != nil {
		return "", err
	}
	// fmt.Println(user)

	return token, nil
}
