package models

import "time"

type LoanStatus string

const (
	ACTIVE LoanStatus = "ACTIVE"
	CLOSED LoanStatus = "CLOSED"
)

type Loan struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	CustomerID      uint          `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"customer_id" binding:"required"`
	PrincipalAmount float64       `json:"principal_amount" binding:"required"`
	InterestRate    float64       `json:"interest_rate" gorm:"default:12"`
	LoanType        string        `gorm:"type:varchar(30)" json:"loan_type" binding:"required"`
	TotalAmount     float64       `json:"total_amount"`
	PendingAmount   float64       `json:"pending_amount"`
	StartDate       int64         `json:"start_date"`
	EndDate         time.Time     `json:"end_date"`
	Status          LoanStatus    `json:"status" gorm:"default:ACTIVE"`
	CreatedAt       int64         `json:"created_at"`
	Customer        *Customer     `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Payments        []LoanPayment `gorm:"foreignKey:LoanID" json:"payments,omitempty"`
}

func (Loan) TableName() string {
	return "loan"
}
