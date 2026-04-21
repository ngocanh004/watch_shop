package models

import "time"

const (
	OrderStatusPending   = "cho_xu_ly"
	OrderStatusConfirmed = "da_xac_nhan"
	OrderStatusShipping  = "dang_giao"
	OrderStatusDone      = "hoan_thanh"
	OrderStatusCancelled = "da_huy"
)

type Order struct {
	ID            uint        `gorm:"primaryKey" json:"id"`
	MaDonHang     string      `gorm:"size:30;uniqueIndex" json:"ma_don_hang"`
	UserID        uint        `gorm:"index" json:"user_id"`
	HoTen         string      `gorm:"size:100" json:"ho_ten"`
	SoDienThoai   string      `gorm:"size:20" json:"so_dien_thoai"`
	DiaChi        string      `gorm:"size:255" json:"dia_chi"`
	GhiChu        string      `gorm:"type:text" json:"ghi_chu"`
	TongTien      float64     `gorm:"type:decimal(15,2)" json:"tong_tien"`
	PhuongThuc    string      `gorm:"size:30;default:cod" json:"phuong_thuc"`
	TrangThai     string      `gorm:"size:30;default:cho_xu_ly" json:"trang_thai"`
	NgayDatHang   time.Time   `json:"ngay_dat_hang"`
	ChiTietDonHang []OrderItem `gorm:"foreignKey:OrderID" json:"chi_tiet_don_hang,omitempty"`
	User          User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderItem struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	OrderID    uint    `gorm:"index" json:"order_id"`
	ProductID  uint    `gorm:"index" json:"product_id"`
	TenSanPham string  `gorm:"size:200" json:"ten_san_pham"`
	HinhAnh    string  `gorm:"size:255" json:"hinh_anh"`
	Gia        float64 `gorm:"type:decimal(15,2)" json:"gia"`
	SoLuong    int     `json:"so_luong"`
	ThanhTien  float64 `gorm:"type:decimal(15,2)" json:"thanh_tien"`
	Product    Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
