package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	HoTen       string    `gorm:"size:100;not null" json:"ho_ten"`
	SoDienThoai string    `gorm:"size:20" json:"so_dien_thoai"`
	Email       string    `gorm:"size:100;uniqueIndex" json:"email"`
	TenDangNhap string    `gorm:"size:50;uniqueIndex;not null" json:"ten_dang_nhap"`
	MatKhau     string    `gorm:"size:255;not null" json:"-"`
	VaiTro      string    `gorm:"size:20;default:user" json:"vai_tro"`
	DiaChi      string    `gorm:"size:255" json:"dia_chi"`
	TrangThai   bool      `gorm:"default:true" json:"trang_thai"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
