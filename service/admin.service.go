package service

import "co-studio-e-commerce/repo"

type Admin struct {
	repo repo.IRepo
}

type IAdmin interface {
}

func NewAdmin(repo repo.IRepo) *Admin {
	return &Admin{repo: repo}
}

// 1. Thêm/ quản ly nhân viên
// 2. Xác thực và đăng nhập
// 3. Quản l quyền hạn và vai trò
// 4. Đặt lại mặt khẩu
// 5. Theo dõi và bảo mật tài khoản
// 6. Báo cáo va ghi log
