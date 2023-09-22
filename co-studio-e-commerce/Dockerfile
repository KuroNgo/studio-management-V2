# Sử dụng hình ảnh Alpine Go
FROM golang:1.20-alpine

# Thiết lập thư mục làm việc là /app
WORKDIR /app

# Sao chép mã nguồn của ứng dụng Go vào thư mục /app trong container
COPY ../.. .

# Tải và cài đặt các phụ thuộc (nếu cần)
RUN go mod download

# Biên dịch ứng dụng Go
RUN go build -o main .

# Expose port 8000 cho ứng dụng Go
EXPOSE 8080

# Chạy ứng dụng Go khi container được khởi chạy và cung cấp các biến môi trường cho kết nối đến PostgreSQL
CMD ["./main"]

# Các biến môi trường cho kết nối đến PostgreSQL
ENV DB_HOST=localhost
ENV DB_NAME=postgres
ENV DB_PASSWORD=mysecretpassword
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV PORT=8000
ENV SSLMODE=disable
