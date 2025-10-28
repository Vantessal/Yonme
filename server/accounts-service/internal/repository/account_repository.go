package repository

import (
	"accounts-service/internal/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) CreateAccount(account *models.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) GetAll() ([]models.Account, error) {
	var accounts []models.Account
	err := r.db.Find(&accounts).Error
	return accounts, err
}

func (r *AccountRepository) GetAccountByID(id uuid.UUID) (*models.Account, error) {
	var account models.Account
	err := r.db.First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) UpdateAccount(account *models.Account) error {
	ctx := context.Background()
	_, err := gorm.G[models.Account](r.db).Where("id = ?", account.ID).Updates(ctx,*account)
	if err != nil {
		return err
	}
	err = r.db.First(&account, "id = ?", account.ID).Error; 
    return err
}

func (r *AccountRepository) DeleteAccount(id uuid.UUID) error {
	return r.db.Delete(&models.Account{}, id).Error
}