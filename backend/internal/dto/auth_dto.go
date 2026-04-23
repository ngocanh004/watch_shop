package dto

type RegisterRequest struct {
	HoTen          string `json:"ho_ten"`
	SoDienThoai    string `json:"so_dien_thoai"`
	Email          string `json:"email"`
	TenDangNhap    string `json:"ten_dang_nhap"`
	MatKhau        string `json:"mat_khau"`
	XacNhanMatKhau string `json:"xac_nhan_mat_khau"`
}

type LoginRequest struct {
	TenDangNhap string `json:"ten_dang_nhap"`
	MatKhau     string `json:"mat_khau"`
}

type ChangePasswordRequest struct {
	MatKhauCu      string `json:"mat_khau_cu"`
	MatKhauMoi     string `json:"mat_khau_moi"`
	XacNhanMatKhau string `json:"xac_nhan_mat_khau"`
}

type UpdateProfileRequest struct {
	HoTen       string `json:"ho_ten"`
	SoDienThoai string `json:"so_dien_thoai"`
	Email       string `json:"email"`
	DiaChi      string `json:"dia_chi"`
}
