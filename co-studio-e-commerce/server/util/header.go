package util

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUserId là hàm lấy user id từ header
// Nếu không có user id thì trả về uuid.Nil
func GetXUserId(ctx *gin.Context) (uuid.UUID, error) {
	userIdStr := ctx.GetHeader("x-user-id")
	if strings.Contains(userIdStr, "|") {
		userIdStr = strings.Split(userIdStr, "|")[0]
	}
	res, err := uuid.Parse(userIdStr)
	if err != nil {
		return uuid.Nil, err
	}
	return res, nil
}
