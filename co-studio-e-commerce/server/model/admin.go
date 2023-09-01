package model

type Admin struct {
	Admin_ID   int    `gorm:"primary_key;AUTO_INCREMENT" json:"admin_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

// Đặt tên trong database là admin
func (Admin) TableName() string {
	return "admin"
}

// Kiểm tra xem admin đã được set chưa
// Nếu admin_id khác 0 thì đã được set
// Nếu admin_id bằng 0 thì chưa được set
func (a *Admin) IsSet() bool {
	return a.Admin_ID != 0
}
