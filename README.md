# Watch Shop - Hệ thống bán đồng hồ trực tuyến

Đồ án tốt nghiệp: **Xây dựng hệ thống quản lý và bán đồng hồ trực tuyến với kiến trúc Backend Golang và Web Admin Vue 3**

## Kiến trúc

```
watch-shop/
├── backend/     # API server - Go + Gin + GORM + PostgreSQL
├── admin/       # Trang quản trị - Vue 3 + Vite + Pinia
└── client/      # Trang khách hàng - Vue 3 + Vite + Pinia
```

## Công nghệ sử dụng

**Backend:**
- Ngôn ngữ: Go 1.22
- Framework: Gin
- ORM: GORM
- Database: PostgreSQL 14+
- Xác thực: JWT (golang-jwt/jwt/v5)
- Mã hoá mật khẩu: bcrypt

**Frontend (Admin + Client):**
- Framework: Vue 3 (Composition API)
- Build: Vite 5
- State: Pinia
- Router: Vue Router 4
- HTTP: Axios

## Cách chạy

### 1. Backend

```bash
cd backend
copy .env.example .env
# Sửa thông tin DB trong .env
go mod tidy
go run ./cmd/server
```

Server: `http://localhost:8080`. Khi chạy lần đầu sẽ tự migrate và seed dữ liệu mẫu.

### 2. Trang quản trị

```bash
cd admin
copy .env.example .env
npm install
npm run dev
```

Truy cập: `http://localhost:5174`

Đăng nhập: `admin01` / `Admin@123`

### 3. Trang khách hàng

```bash
cd client
copy .env.example .env
npm install
npm run dev
```

Truy cập: `http://localhost:5173`

Đăng nhập thử: `user01` / `User@123` (hoặc đăng ký tài khoản mới)

## Tính năng

### Client (khách hàng)
- Trang chủ với banner, sản phẩm nổi bật, sản phẩm mới
- Tìm kiếm, lọc sản phẩm nâng cao (danh mục, thương hiệu, giới tính, loại máy, loại dây, giá)
- Xem chi tiết sản phẩm, sản phẩm liên quan
- Đăng ký, đăng nhập, đổi mật khẩu, cập nhật thông tin
- Giỏ hàng, yêu thích
- Đặt hàng (Checkout), thanh toán COD/chuyển khoản
- Lịch sử đơn hàng, chi tiết, huỷ đơn

### Admin (quản trị)
- Dashboard thống kê (doanh thu, đơn hàng, sản phẩm, người dùng)
- Quản lý sản phẩm (CRUD, upload ảnh)
- Quản lý danh mục, thương hiệu
- Quản lý người dùng, reset mật khẩu
- Quản lý đơn hàng, cập nhật trạng thái
- Cập nhật thông tin cá nhân, đổi mật khẩu

## Cấu trúc CSDL

- `users`: Người dùng (admin/user)
- `categories`: Danh mục đồng hồ
- `brands`: Thương hiệu
- `products`: Sản phẩm đồng hồ
- `product_images`: Ảnh phụ của sản phẩm
- `cart_items`: Giỏ hàng
- `wishlist_items`: Yêu thích
- `orders`: Đơn hàng
- `order_items`: Chi tiết đơn hàng

## Tài liệu API

Xem chi tiết tại [backend/README.md](backend/README.md).
