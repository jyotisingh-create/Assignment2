package models

type TransactionType string

const (
	CREDIT TransactionType = "credit"
	DEBIT  TransactionType = "debit"
)

type Transaction struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	AccountID uint            `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"account_id" binding:"required"`
	Type      TransactionType `json:"type" binding:"required"`
	Amount    float64         `json:"amount" binding:"required"`
	CreatedAt int64           `json:"created_at"`
	Account   *SavingsAccount `gorm:"foreignKey:AccountID" json:"account,omitempty"`
}

func (Transaction) TableName() string {
	return "transaction"
}
