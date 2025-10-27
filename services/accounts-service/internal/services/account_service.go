package services

import (
	"accounts-service/internal/models"
	"accounts-service/internal/repository"

	"github.com/google/uuid"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Create(account *models.Account) error {
	return s.repo.CreateAccount(account)
}

func (s *AccountService) GetAll() ([]models.Account, error) {
	return s.repo.GetAll()
}

func (s *AccountService) GetByID(id uuid.UUID) (*models.Account, error) {
	return s.repo.GetAccountByID(id)
}

func (s *AccountService) Update(account *models.Account) error {
	return s.repo.UpdateAccount(account)
}

func (s *AccountService) Delete(id uuid.UUID) error {
	return s.repo.DeleteAccount(id)
}
