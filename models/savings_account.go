package models

type SavingsAccount struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	Balance          float64           `json:"balance"`
	CreatedAt        int64             `json:"created_at"`
	CustomerAccounts []CustomerAccount `gorm:"foreignKey:AccountID" json:"customer_accounts,omitempty"`
	Transactions     []Transaction     `gorm:"foreignKey:AccountID" json:"transactions,omitempty"`
}

func (SavingsAccount) TableName() string {
	return "savings_account"
}
