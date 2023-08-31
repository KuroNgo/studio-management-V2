package model

type Admin struct {
	admin_id   int    `gorm:"primary_key;AUTO_INCREMENT" json:"admin_id"`
	username   string `json:"username"`
	password   string `json:"password"`
	first_name string `json:"first_name"`
	last_name  string `json:"last_name"`
	email      string `json:"email"`
	phone      string `json:"phone"`
}

// Đặt tên trong database là admin
func (Admin) TableName() string {
	return "admin"
}

// Kiểm tra xem admin đã được set chưa
// Nếu admin_id khác 0 thì đã được set
// Nếu admin_id bằng 0 thì chưa được set
func (a *Admin) IsSet() bool {
	return a.admin_id != 0
}
