package service

import (
	"github.com/analopesdev/go-gateway/internal/domain"
	"github.com/analopesdev/go-gateway/internal/dto"
)

type AccountService struct {
	accountRepository domain.AccountRepository
}

func NewAccountService(accountRepository domain.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (s *AccountService) CreateAccount(input *dto.CreateAccount) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.accountRepository.FindByApiKey(account.ApiKey)

	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrDuplicatedApiKey
	}

	newAccount := domain.NewAccount(account.Name, account.Email)

	err = s.accountRepository.Save(newAccount)

	if err != nil {
		return nil, err
	}

	return dto.FromAccount(newAccount), nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.accountRepository.FindByApiKey(apiKey)

	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)

	err = s.accountRepository.UpdateBalance(account)

	if err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}

func (s *AccountService) FindByApiKey(apiKey string) (*domain.Account, error) {
	account, err := s.accountRepository.FindByApiKey(apiKey)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountService) FindById(id string) (*domain.Account, error) {
	account, err := s.accountRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	return account, nil
}
