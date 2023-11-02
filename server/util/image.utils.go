package util

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// IsImageFile  kiểm tra xem tệp có phải là hình ảnh không bằng phần mở rộng tệp
// hoặc kiểm tra loại MIME
func IsImageFile(filename string) bool {
	extension := []string{".jpg", ".jpeg", ".png", ".gif"}
	ext := filepath.Ext(filename)
	for _, validExt := range extension {
		if ext == validExt {
			return true
		}
	}
	// Kiểm tra loại MIME (cần truy cập tệp trên đĩa để kiểm tra loại MIME)
	mimeType, err := getFileMimeType(filename)
	if err == nil && strings.HasPrefix(mimeType, "image/") {
		return true
	}

	return false
}

// getFileMimeType lấy loại MIME của tệp
func getFileMimeType(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(buffer)
	return mimeType, nil
}
