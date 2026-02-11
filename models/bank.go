package models

type Bank struct {
	ID        uint     `gorm:"primaryKey" json:"id"`
	Name      string   `json:"name" binding:"required"`
	CreatedAt int64    `json:"created_at"`
	Branches  []Branch `gorm:"foreignKey:BankID" json:"branches,omitempty"`
}

func (Bank) TableName() string {
	return "bank"
}
