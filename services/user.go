package services

import (
	"context"

	"github.com/tumeraltunbas/go-blog/models/entities"
	"github.com/tumeraltunbas/go-blog/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	return user, err
}

func (s *UserService) InsertUser(ctx context.Context, email string, password string) error {
	err := s.userRepository.InsertUser(ctx, email, password)
	return err
}
