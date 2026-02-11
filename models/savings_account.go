package models

type SavingsAccount struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	CustomerID   uint          `gorm:"index" json:"customer_id" binding:"required"`
	Balance      float64       `json:"balance"`
	CreatedAt    int64         `json:"created_at"`
	Customer     *Customer     `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Transactions []Transaction `gorm:"foreignKey:AccountID" json:"transactions,omitempty"`
}

func (SavingsAccount) TableName() string {
	return "savings_account"
}
