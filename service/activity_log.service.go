package service

import (
	"co-studio-e-commerce/repo"
)

type activityLog struct {
	repo repo.IRepo
}

type IActivityLog interface {
}

func NewActivityLog(repo repo.IRepo) *activityLog {
	return &activityLog{repo: repo}
}

// 1. Ghi log sự kiện
// 2. Lọc và truy vấn
// 3. Bảo mật và quản lý quyền truy cập
// 4. Báo cáo và theo dõi
// 5. Dự phòng và sao lưu log
// 6. xử lý log trực tiếp hoặc thông qua dịch vụ bên ngoài
// 7. Giới thiệu người dùng về sự cố
// 8. Quản lý kích thước và thời gian lưu trữ
// 9. Tích hợp công cụ giám sát
