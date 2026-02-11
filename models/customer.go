package models

type Customer struct {
	ID             uint            `gorm:"primaryKey" json:"id"`
	BranchID       uint            `gorm:"index" json:"branch_id" binding:"required"`
	Name           string          `json:"name" binding:"required"`
	Email          string          `json:"email" binding:"required,email"`
	Phone          string          `json:"phone" binding:"required"`
	CreatedAt      int64           `json:"created_at"`
	Branch         *Branch         `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	SavingsAccount *SavingsAccount `gorm:"foreignKey:CustomerID" json:"savings_account,omitempty"`
	Loans          []Loan          `gorm:"foreignKey:CustomerID" json:"loans,omitempty"`
}

func (Customer) TableName() string {
	return "customer"
}
