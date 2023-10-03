package middleware

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/util"
)

func CurrentUserIsAdmin(user model.User) bool {
	// Kiểm tra nếu trường quyền của người dùng là "admin"
	return user.Role == util.ADMIN_ROLE
}

func CurrenUserIsUser(user model.User) bool {
	// Kiểm tra nếu trường quyền của người dùng là "user"
	return user.Role == util.CUSTOMER_ROLE
}

func CurrenUserIsGuest(user model.User) bool {
	// Kiểm tra nếu trường quyền của người dùng là "guest"
	return user.Role == util.GUEST_ROLE
}

func currenUserIsStaff(user model.User) bool {
	// Kiểm tra nếu trường quyền của người dùng là "staff"
	return user.Role == util.STAFF_ROLE
}
