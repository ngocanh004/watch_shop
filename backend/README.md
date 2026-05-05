# Backend - Watch Shop (Go + Gin + PostgreSQL)

Hệ thống API phía server cho website bán đồng hồ.

## Yêu cầu môi trường

- Go 1.22 trở lên
- PostgreSQL 14 trở lên
- (Tuỳ chọn) pgAdmin để quản lý DB

## Cài đặt

1. Tạo database trống trong PostgreSQL:

```sql
CREATE DATABASE watch_shop;
```

2. Sao chép file cấu hình:

```bash
copy .env.example .env
```

Chỉnh sửa lại `DB_USER`, `DB_PASSWORD`, `DB_NAME` cho phù hợp với máy.

3. Tải dependencies:

```bash
go mod tidy
```

4. Chạy server:

```bash
go run ./cmd/server
```

Lần đầu chạy, hệ thống tự động:
- Tạo các bảng trong database (auto migrate)
- Thêm dữ liệu mẫu (admin, người dùng, danh mục, thương hiệu, sản phẩm)

Mặc định server chạy tại `http://localhost:8080`.

## Tài khoản mặc định

| Vai trò | Tên đăng nhập | Mật khẩu  |
|---------|---------------|-----------|
| Admin   | admin01       | Admin@123 |
| User    | user01        | User@123  |

## Cấu trúc thư mục

```
backend/
├── cmd/server/main.go       # Điểm vào chính
├── internal/
│   ├── config/              # Đọc cấu hình
│   ├── database/            # Kết nối, migrate, seed
│   ├── models/              # GORM models
│   ├── dto/                 # Cấu trúc request
│   ├── handlers/            # Xử lý HTTP
│   ├── middleware/          # JWT, phân quyền
│   ├── routes/              # Khai báo routes
│   └── utils/               # JWT, mật khẩu, validate, response
├── uploads/                 # Lưu hình ảnh sản phẩm
└── go.mod
```

## Danh sách API chính

### Công khai

- `POST /api/auth/dang-ky` - Đăng ký
- `POST /api/auth/dang-nhap` - Đăng nhập
- `GET  /api/products` - Danh sách sản phẩm (filter, phân trang)
- `GET  /api/products/:id` - Chi tiết sản phẩm
- `GET  /api/categories` - Danh mục
- `GET  /api/brands` - Thương hiệu

### Người dùng (cần JWT)

- `GET  /api/me` - Thông tin cá nhân
- `PUT  /api/me` - Cập nhật cá nhân
- `POST /api/auth/doi-mat-khau` - Đổi mật khẩu
- `GET/POST/PUT/DELETE /api/cart` - Giỏ hàng
- `GET/POST/DELETE /api/wishlist` - Yêu thích
- `POST /api/orders` - Đặt hàng
- `GET  /api/orders/me` - Lịch sử đơn hàng

### Admin (cần JWT + role=admin)

- `GET  /api/admin/thong-ke` - Thống kê dashboard
- CRUD `/api/admin/users`
- CRUD `/api/admin/categories`
- CRUD `/api/admin/brands`
- CRUD `/api/admin/products`
- `GET  /api/admin/orders` - Danh sách đơn hàng
- `PUT  /api/admin/orders/:id/trang-thai` - Cập nhật trạng thái
- `POST /api/admin/upload` - Upload ảnh

## Filter sản phẩm hỗ trợ

- `keyword`, `category_id`, `brand_id`
- `loai_may`, `day_deo`, `gioi_tinh`
- `gia_min`, `gia_max`
- `noi_bat=1`
- `sort=moi_nhat | gia_tang | gia_giam | ban_chay`
- `page`, `page_size`
