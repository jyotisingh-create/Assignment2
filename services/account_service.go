package services

import (
	"banking-system/config"
	"banking-system/models"
	"errors"
)

type AccountService struct{}

func (s *AccountService) OpenSavingsAccount(customerID uint) (*models.SavingsAccount, error) {
	account := &models.SavingsAccount{
		CustomerID: customerID,
		Balance:    0,
		CreatedAt:  0,
	}

	var customer models.Customer
	if err := config.DB.First(&customer, customerID).Error; err != nil {
		return nil, errors.New("customer not found")
	}

	var existingAccount models.SavingsAccount
	if err := config.DB.Where("customer_id = ?", customerID).First(&existingAccount).Error; err == nil {
		return nil, errors.New("savings account already exists for this customer")
	}

	if err := config.DB.Create(account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountService) GetAccount(accountID uint) (*models.SavingsAccount, error) {
	var account models.SavingsAccount
	if err := config.DB.Preload("Transactions").First(&account, accountID).Error; err != nil {
		return nil, errors.New("account not found")
	}
	return &account, nil
}

func (s *AccountService) Deposit(accountID uint, amount float64) (*models.SavingsAccount, error) {
	if amount <= 0 {
		return nil, errors.New("deposit amount must be positive")
	}

	var account models.SavingsAccount
	if err := config.DB.First(&account, accountID).Error; err != nil {
		return nil, errors.New("account not found")
	}

	tx := config.DB.Begin()

	account.Balance += amount
	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	transaction := &models.Transaction{
		AccountID: accountID,
		Type:      models.DEPOSIT,
		Amount:    amount,
		CreatedAt: 0,
	}
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *AccountService) Withdraw(accountID uint, amount float64) (*models.SavingsAccount, error) {
	if amount <= 0 {
		return nil, errors.New("withdrawal amount must be positive")
	}

	var account models.SavingsAccount
	if err := config.DB.First(&account, accountID).Error; err != nil {
		return nil, errors.New("account not found")
	}

	if account.Balance < amount {
		return nil, errors.New("insufficient balance")
	}

	tx := config.DB.Begin()

	account.Balance -= amount
	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	transaction := &models.Transaction{
		AccountID: accountID,
		Type:      models.WITHDRAW,
		Amount:    amount,
		CreatedAt: 0,
	}
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *AccountService) GetTransactions(accountID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := config.DB.Where("account_id = ?", accountID).Order("created_at DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *AccountService) GetBalance(accountID uint) (float64, error) {
	var account models.SavingsAccount
	if err := config.DB.First(&account, accountID).Error; err != nil {
		return 0, errors.New("account not found")
	}
	return account.Balance, nil
}
