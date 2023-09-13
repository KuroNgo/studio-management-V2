
# Giải thích cấu trúc file
## Util
Nơi thực hiện lưu hằng số

Nơi thực hiện lưu các hàm hỗ trợ

## Model
Lưu trữ struct

## Config
Khởi tạo app
Connect DB
Config ENV

## Route
Định hướng URL tới function đích

## Handler
Xử lý request | response
`
Cách nhận biết: Có ctx *gin.Context là handler
ctx *gin.Context là một kiểu dữ liệu đại diện cho context của một request
Nó chứa các phương thức và thuốc tính cho phpes thao tác với request và response dễ dàng hơn
`


## Middleware
Bước trung gian giữa route và handler
VD: Check xem user có đủ quyền dùng app đó không

## Service
Xử lý logic chính

## Repository
Tương tác với database

# Giải thích (lý do và hướng chạy) các gói đã tạo trong dự án
### github.com/gin-gonic/gin
`
Gin được xây dựng dựa trên httprouter, một router nhanh và nhẹ được viết bằng Go.
Gin cung cấp một cách đơn giản để tạo các ứng dụng web hiệu suất cao và có thể mở rộng.
`

### gorm.io/gorm
`
GORM là một ORM (Object Relational Mapping) cho Golang.
`

### golang.org/x/crypto/bcrypt
`
Package bcrypt cung cấp hàm băm bcrypt cho Go (cách băm này khác
với MD5 ở chỗ .
`

- Reactivex
`
ReactiveX là một thư viện lập trình bất đồng bộ được thiết kế để xử lý luồng dữ liệu bất đồng bộ.
Giúp xử lý các tác vụ bất đồng bộ một cách dễ dàng hơn mà không cần
nhiều kiến thứ về nó
Làm cho code dễ đọc

`

## Activity_log
 - Ghi log sự kiện: Cho phép ứng dụng ghi lại các sự kiện quan trọng,
ví dụ như đăng nhập, đăng ký, thay đổi mật khẩu,thao tác cơ sở dữ liệu quan trọng,
hoặc bất kỳ hoạt động nào mà bạn muốn theo dõi.

 - Lọc và truy vấn log: Cung cấp khả năng lọc và truy vấn các sự kiện
trong log dựa trên các thông tin như thời gian, loại sự kiện, người dùng thực hiện, v.v.
Điều này giúp trong việc tìm kiếm và xem lại các sự kiện cụ thể.

 - Bảo mật và quản lý quyền truy cập: Xác thực và ủy quyền người dùng cho việc xem và ghi log.
Có thể có các vai trò khác nhau cho việc quản lý log.

- Báo cáo và theo dõi: Tạo các báo cáo hoặc biểu đồ từ dữ liệu log để theo dõi xu hướng,
phát hiện sự cố hoặc kiểm tra hiệu suất hệ thống.

 - Dự phòng và sao lưu log: Đảm bảo rằng log được sao lưu định kỳ để đối phó với sự cố hoặc phục hồi dữ liệu.

- Xử lý log trực tiếp hoặc thông qua dịch vụ bên ngoài: Có thể cần tích hợp log với các dịch vụ
và công cụ bên ngoài như Elasticsearch, Logstash, hoặc Splunk để xử lý và phân tích log một cách hiệu quả.

- Giới thiệu người dùng về sự cố: Ghi log có thể được sử dụng để cảnh báo người dùng hoặc
quản trị viên về các sự cố quan trọng, ví dụ như lỗi hệ thống hoặc hoạt động bất thường.

 - Quản lý kích thước và thời gian lưu trữ: Xác định cách dữ liệu log được quản lý,
bao gồm việc định kỳ xoá dữ liệu cũ hoặc lưu trữ dữ liệu trong một khoảng thời gian nhất định.

 - Tích hợp với công cụ giám sát: Tích hợp log với các công cụ giám sát như Prometheus
hoặc Grafana để theo dõi sức khỏe và hiệu suất hệ thống.