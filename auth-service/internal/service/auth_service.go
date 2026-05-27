package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"errors"
	sharedjwt "venuex/shared/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id string) (models.User, error)
	GetAllUsers() ([]models.User, error)
	DeleteUser(userID string) error
	CreateAuditLog(log models.AuditLog) error
}

type authService struct {
	repo UserRepository
}

func NewAuthService(
	repo UserRepository,
) *authService {

	return &authService{
		repo: repo,
	}
}

var AuthService = NewAuthService(
	repositoryWrapper{},
)

func (s *authService) Register(
	req models.RegisterRequest,
) error {

	allowedRoles := map[string]bool{
		"user":   true,
		"artist": true,
		"venue":  true,
	}

	if !allowedRoles[req.Role] {
		return errors.New("invalid role")
	}

	existingUser, err := s.repo.GetUserByEmail(req.Email)

	if err == nil && existingUser.ID != "" {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	err = s.repo.CreateUser(user)

	if err != nil {
		return err
	}

	PublishNotification(
		"new user registered: " + user.Email,
	)

	CreateAuditLog(models.AuditLog{
		Action: "User registered: " + user.Email,
	})

	return nil
}

	func (s *authService) Login(
		req models.LoginRequest,
		) (string, error) {

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return "", err
	}

	token, err := sharedjwt.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

type repositoryWrapper struct{}

func (repositoryWrapper) CreateUser(
	user models.User,
) error {

	return repository.CreateUser(user)
}

func (repositoryWrapper) GetUserByEmail(
	email string,
) (models.User, error) {

	return repository.GetUserByEmail(email)
}

func (repositoryWrapper) GetUserByID(
	id string,
) (models.User, error) {

	return repository.GetUserByID(id)
}

func (repositoryWrapper) GetAllUsers() (
	[]models.User,
	error,
) {

	return repository.GetAllUsers()
}

func (repositoryWrapper) DeleteUser(
	userID string,
) error {

	return repository.DeleteUser(userID)
}

func (repositoryWrapper) CreateAuditLog(
	log models.AuditLog,
) error {

	return repository.CreateAuditLog(log)
}