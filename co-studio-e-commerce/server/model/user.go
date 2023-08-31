package model

type User struct {
	BaseModel
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Role         string `json:"role"`
	Address      string `json:"address"`
	Phone_Number string `json:"phone_number"`
	Is_Admin     bool   `json:"is_admin"`
}

func (User) TableName() string {
	return "user"
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
