package services

import (
	"github.com/natanaelrusli/hexagonal-arch/internals/core/domain"
	"github.com/natanaelrusli/hexagonal-arch/internals/core/ports"
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
	s.accountRepository.CreateAccount(&domain.Account{
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: password,
		Balance:           0,
		Number:            280198,
	})

	return nil
}
