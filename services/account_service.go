package services

import (
	"banking-system/config"
	"banking-system/models"
	"errors"
)

type AccountService struct{}

func (s *AccountService) OpenSavingsAccount(customerID uint) (*models.SavingsAccount, error) {
	account := &models.SavingsAccount{
		Balance:   0,
		CreatedAt: 0,
	}

	var customer models.Customer
	if err := config.DB.First(&customer, customerID).Error; err != nil {
		return nil, errors.New("customer not found")
	}

	tx := config.DB.Begin()
	if err := tx.Create(account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	custAcc := &models.CustomerAccount{
		CustomerID: customerID,
		AccountID:  account.ID,
		HolderRole: models.PRIMARY_HOLDER,
		CreatedAt:  0,
	}

	if err := tx.Create(custAcc).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
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
		Type:      models.CREDIT,
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
		Type:      models.DEBIT,
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

func (s *AccountService) AddAccountHolder(accountID uint, customerID uint, holderRole string) (*models.CustomerAccount, error) {

	var account models.SavingsAccount
	if err := config.DB.First(&account, accountID).Error; err != nil {
		return nil, errors.New("account not found")
	}

	var customer models.Customer
	if err := config.DB.First(&customer, customerID).Error; err != nil {
		return nil, errors.New("customer not found")
	}

	var existing models.CustomerAccount
	if err := config.DB.Where("customer_id = ? AND account_id = ?", customerID, accountID).First(&existing).Error; err == nil {
		return nil, errors.New("customer is already a holder of this account")
	}

	custAcc := &models.CustomerAccount{
		CustomerID: customerID,
		AccountID:  accountID,
		HolderRole: holderRole,
		CreatedAt:  0,
	}

	if err := config.DB.Create(custAcc).Error; err != nil {
		return nil, err
	}

	return custAcc, nil
}

func (s *AccountService) RemoveAccountHolder(accountID uint, customerID uint) error {

	var account models.SavingsAccount
	if err := config.DB.First(&account, accountID).Error; err != nil {
		return errors.New("account not found")
	}

	var custAcc models.CustomerAccount
	if err := config.DB.Where("customer_id = ? AND account_id = ?", customerID, accountID).First(&custAcc).Error; err != nil {
		return errors.New("customer is not a holder of this account")
	}

	var count int64
	if err := config.DB.Model(&models.CustomerAccount{}).Where("account_id = ?", accountID).Count(&count).Error; err != nil {
		return err
	}

	if count <= 1 {
		return errors.New("cannot remove the last holder from an account")
	}

	if err := config.DB.Delete(&custAcc).Error; err != nil {
		return err
	}

	return nil
}

func (s *AccountService) GetAccountHolders(accountID uint) ([]models.CustomerAccount, error) {

	var account models.SavingsAccount
	if err := config.DB.First(&account, accountID).Error; err != nil {
		return nil, errors.New("account not found")
	}

	var holders []models.CustomerAccount
	if err := config.DB.Where("account_id = ?", accountID).Preload("Customer").Find(&holders).Error; err != nil {
		return nil, err
	}

	return holders, nil
}
