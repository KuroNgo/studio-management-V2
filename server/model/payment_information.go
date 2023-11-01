package model

import "github.com/google/uuid"

type PaymentInformation struct {
	PaymentID      int       `gorm:"primary_key;AUTO_INCREMENT" json:"payment_id"`
	UserID         uuid.UUID `json:"user_id"`
	CardNumber     string    `json:"card_number"`
	CardHolderName string    `json:"card_holder_name"` // This is the name of the card holder
	ExpireDate     string    `json:"expire_date"`
	CVV            string    `json:"cvv"`
	BillingAddress string    `json:"billing_address"`
	User           User      `gorm:"foreignKey:UserID"`
}

func (PaymentInformation) TableName() string {
	return "payment_information"
}
