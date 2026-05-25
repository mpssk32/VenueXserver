package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Register(req models.RegisterRequest) error {

	allowedRoles := map[string]bool{
		"user":   true,
		"artist": true,
		"venue":  true,
	}

	if !allowedRoles[req.Role] {
		return errors.New("invalid role")
	}

	existingUser, err := repository.GetUserByEmail(req.Email)

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

	err = repository.CreateUser(user)

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

func Login(req models.LoginRequest) (string, error) {
	user, err := repository.GetUserByEmail(req.Email)
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

	token, err := GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}