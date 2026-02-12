package models

const (
	PRIMARY_HOLDER string = "primary_holder"
	JOINT_HOLDER   string = "joint_holder"
)

type CustomerAccount struct {
	ID         uint            `gorm:"primaryKey" json:"id"`
	CustomerID uint            `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"customer_id" binding:"required"`
	AccountID  uint            `gorm:"index;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"account_id" binding:"required"`
	HolderRole string          `gorm:"type:varchar(20)" json:"holder_role" binding:"required"`
	CreatedAt  int64           `json:"created_at"`
	Customer   *Customer       `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Account    *SavingsAccount `gorm:"foreignKey:AccountID" json:"account,omitempty"`
}

func (CustomerAccount) TableName() string {
	return "customer_account"
}
