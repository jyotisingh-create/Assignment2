package services

import (
	"banking-system/config"
	"banking-system/models"
	"errors"
	"math"
	"time"
)

type LoanService struct{}

func (s *LoanService) TakeLoan(customerID uint, principalAmount float64, loanType string, durationYears float64) (*models.Loan, error) {
	if principalAmount <= 0 {
		return nil, errors.New("principal amount must be positive")
	}
	if durationYears <= 0 {
		return nil, errors.New("duration_years must be positive")
	}
	if loanType == "" {
		return nil, errors.New("loan_type is required")
	}

	var customer models.Customer
	if err := config.DB.First(&customer, customerID).Error; err != nil {
		return nil, errors.New("customer not found")
	}

	const yearlyRate = 0.12
	totalAmount := principalAmount + (principalAmount * yearlyRate * durationYears)

	totalAmount = math.Round(totalAmount*100) / 100

	endDate := time.Now().AddDate(int(durationYears), 0, 0)

	loan := &models.Loan{
		CustomerID:      customerID,
		PrincipalAmount: principalAmount,
		LoanType:        loanType,
		InterestRate:    12.0,
		TotalAmount:     totalAmount,
		PendingAmount:   totalAmount,
		StartDate:       time.Now().Unix(),
		EndDate:         endDate,
		Status:          models.ACTIVE,
		CreatedAt:       time.Now().Unix(),
	}

	if err := config.DB.Create(loan).Error; err != nil {
		return nil, err
	}

	return loan, nil
}

func (s *LoanService) GetLoan(loanID uint) (*models.Loan, error) {
	var loan models.Loan
	if err := config.DB.Preload("Payments").First(&loan, loanID).Error; err != nil {
		return nil, errors.New("loan not found")
	}
	return &loan, nil
}

func (s *LoanService) RepayLoan(loanID uint, amount float64) (*models.Loan, error) {
	if amount <= 0 {
		return nil, errors.New("repayment amount must be positive")
	}

	var loan models.Loan
	if err := config.DB.First(&loan, loanID).Error; err != nil {
		return nil, errors.New("loan not found")
	}

	if loan.Status == models.CLOSED {
		return nil, errors.New("loan is already closed")
	}

	if amount > loan.PendingAmount {
		return nil, errors.New("repayment amount exceeds pending amount")
	}

	tx := config.DB.Begin()

	loan.PendingAmount -= amount

	if loan.PendingAmount <= 0 {
		loan.Status = models.CLOSED
		loan.PendingAmount = 0
	}

	if err := tx.Save(&loan).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	payment := &models.LoanPayment{
		LoanID:      loanID,
		Amount:      amount,
		PaymentDate: time.Now().Unix(),
		CreatedAt:   time.Now().Unix(),
	}
	if err := tx.Create(payment).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &loan, nil
}

func (s *LoanService) CalculateInterest(loanID uint) (float64, error) {
	var loan models.Loan
	if err := config.DB.First(&loan, loanID).Error; err != nil {
		return 0, errors.New("loan not found")
	}

	yearlyInterest := loan.PendingAmount * loan.InterestRate / 100.0

	return math.Round(yearlyInterest*100) / 100, nil
}

func (s *LoanService) GetLoanDetails(loanID uint) (map[string]interface{}, error) {
	loan, err := s.GetLoan(loanID)
	if err != nil {
		return nil, err
	}

	interest, err := s.CalculateInterest(loanID)
	if err != nil {
		return nil, err
	}

	details := map[string]interface{}{
		"id":               loan.ID,
		"customer_id":      loan.CustomerID,
		"principal_amount": loan.PrincipalAmount,
		"loan_type":        loan.LoanType,
		"total_amount":     loan.TotalAmount,
		"interest_rate":    loan.InterestRate,
		"pending_amount":   loan.PendingAmount,
		"start_date":       loan.StartDate,
		"end_date":         loan.EndDate,
		"status":           loan.Status,
		"yearly_interest":  interest,
	}

	return details, nil
}
