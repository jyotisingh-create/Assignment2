package models

type Branch struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	BankID    uint       `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"bank_id" binding:"required"`
	Name      string     `json:"name" binding:"required"`
	Address   string     `json:"address" binding:"required"`
	CreatedAt int64      `json:"created_at"`
	Bank      *Bank      `gorm:"foreignKey:BankID" json:"bank,omitempty"`
	Customers []Customer `gorm:"foreignKey:BranchID" json:"customers,omitempty"`
}

func (Branch) TableName() string {
	return "branch"
}
