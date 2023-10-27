package services

import (
	"errors"

	"github.com/natanaelrusli/hexagonal-arch/internals/core/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.UserService = (*UserService)(nil)

func NewUserService(repository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(email string, password string) error {
	err := s.userRepository.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Register(email string, password string, confirmPass string) error {
	if password != confirmPass {
		return errors.New("the passwords are not equal")
	}

	err := s.userRepository.Register(email, password)
	if err != nil {
		return err
	}
	return nil
}
