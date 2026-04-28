package service

import (
	"auth-service/internal/domain"
	"auth-service/internal/kafka"
	"auth-service/internal/repository"
	"auth-service/pkg/jwt"
	"auth-service/pkg/util"
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type Service struct {
	Repo *repository.Repository
	Auth *jwt.AuthService
	producer *kafka.Producer
}

func NewService(repo *repository.Repository, jwt *jwt.AuthService, producer *kafka.Producer) *Service {
	return &Service{
		Repo: repo,
		Auth: jwt,
		producer: producer,
	}
}

func (s *Service) Register(ctx context.Context, usrRequest *domain.UserRequest) (*domain.UserResponse, error) {
	hashedPassword, err := util.HashPassword(usrRequest.Password)
	if err != nil {
		return nil, err
	}
	usr := &domain.UserDB{
		ID:           uuid.New().String(),
		Email:        usrRequest.Email,
		PasswordHASH: hashedPassword,
	}

	usrResponse, err := s.Repo.Register(ctx, usr)
	if err != nil {
		return nil, err
	}

	s.sendProfileCreatedEvent(usrResponse.ID)
	return usrResponse, nil
}

func (s *Service) Login(ctx context.Context, usrRequest *domain.UserRequest) (string, error) {
	usrDB, err := s.Repo.Login(ctx, usrRequest.Email)
	if err != nil {
		return "", err
	}

	err = util.ValidatePassword(usrRequest.Password, usrDB.PasswordHASH)
	if err != nil {
		return "", domain.ErrWrongPassword
	}

	token, err := s.Auth.GenerateToken(usrDB.ID, usrDB.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GetUserById(ctx context.Context, id string) (*domain.UserResponse, error) {
	return s.Repo.GetUserById(ctx, id)
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.Repo.DeleteUser(ctx, id)
}

// kafka events
func (s *Service) sendProfileCreatedEvent(userID string) {
	event := map[string]interface{}{
		"user_id":  userID,

	}
	value, _ := json.Marshal(event)

	err := s.producer.Send("user.registered", userID, value)
	if err != nil {
		log.Printf("kafka send error:%v\n", err)
	}
}

