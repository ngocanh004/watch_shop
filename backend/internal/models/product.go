package models

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	TenSanPham  string  `gorm:"size:200;not null" json:"ten_san_pham"`
	MoTa        string  `gorm:"type:text" json:"mo_ta"`
	Gia         float64 `gorm:"type:decimal(15,2);not null" json:"gia"`
	GiaGoc      float64 `gorm:"type:decimal(15,2)" json:"gia_goc"`
	SoLuong     int     `gorm:"default:0" json:"so_luong"`
	DaBan       int     `gorm:"default:0" json:"da_ban"`
	HinhAnh     string  `gorm:"size:255" json:"hinh_anh"`
	CategoryID  uint    `json:"category_id"`
	BrandID     uint    `json:"brand_id"`
	LoaiMay     string  `gorm:"size:50" json:"loai_may"`
	DayDeo      string  `gorm:"size:50" json:"day_deo"`
	MatKinh     string  `gorm:"size:50" json:"mat_kinh"`
	DuongKinh   string  `gorm:"size:20" json:"duong_kinh"`
	KhangNuoc   string  `gorm:"size:20" json:"khang_nuoc"`
	GioiTinh    string  `gorm:"size:20" json:"gioi_tinh"`
	XuatXu      string  `gorm:"size:100" json:"xuat_xu"`
	BaoHanh     string  `gorm:"size:50" json:"bao_hanh"`
	NoiBat      bool    `gorm:"default:false" json:"noi_bat"`
	TrangThai   bool    `gorm:"default:true" json:"trang_thai"`

	Category Category       `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Brand    Brand          `gorm:"foreignKey:BrandID" json:"brand,omitempty"`
	HinhAnhs []ProductImage `gorm:"foreignKey:ProductID" json:"hinh_anhs,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Product) TableName() string {
	return "products"
}

type ProductImage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"index" json:"product_id"`
	URL       string    `gorm:"size:255" json:"url"`
	ThuTu     int       `gorm:"default:0" json:"thu_tu"`
	CreatedAt time.Time `json:"created_at"`
}

func (ProductImage) TableName() string {
	return "product_images"
}
