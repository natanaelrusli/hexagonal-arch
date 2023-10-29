package services

import (
	"time"

	"github.com/natanaelrusli/hexagonal-arch/internals/core/domain"
	"github.com/natanaelrusli/hexagonal-arch/internals/core/ports"
	"github.com/natanaelrusli/hexagonal-arch/pkg/encryption"
)

type AccountService struct {
	accountRepository ports.AccountRepository
}

var _ ports.AccountService = (*AccountService)(nil)

func NewAccountService(repository ports.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: repository,
	}
}

func (s *AccountService) CreateAccount(firstName string, lastName string, password string) error {
	hashedPassword, err := encryption.HashAndSalt(password)
	currentTime := time.Now().UTC()

	if err != nil {
		return err
	}

	s.accountRepository.CreateAccount(&domain.Account{
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: hashedPassword,
		Balance:           0,
		Number:            280198,
		CreatedAt:         currentTime,
	})

	return nil
}
