package models

type LoanPayment struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	LoanID      uint    `gorm:"index" json:"loan_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	PaymentDate int64   `json:"payment_date"`
	CreatedAt   int64   `json:"created_at"`
	Loan        *Loan   `gorm:"foreignKey:LoanID" json:"loan,omitempty"`
}

func (LoanPayment) TableName() string {
	return "loan_payment"
}
