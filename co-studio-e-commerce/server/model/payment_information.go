package model

type PaymentInformation struct {
	PaymentID       int    `gorm:"primary_key;AUTO_INCREMENT" json:"payment_id"`
	UserID          int    `json:"user_id"`
	CardNumber      string `json:"card_number"`
	CardHolderName  string `json:"card_holder_name"` // This is the name of the card holder
	ExpireDate      string `json:"expire_date"`
	CVV             string `json:"cvv"`
	Billing_Address string `json:"billing_address"`
}

func (PaymentInformation) TableName() string {
	return "payment_information"
}

func (a *PaymentInformation) IsSet() bool {
	return a.PaymentID != 0
}
