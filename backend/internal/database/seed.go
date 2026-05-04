package database

import (
	"log"

	"watch-shop/backend/internal/models"
	"watch-shop/backend/internal/utils"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		log.Println("Du lieu da co, bo qua seed")
		return
	}

	hash, _ := utils.HashPassword("Admin@123")
	admin := models.User{
		HoTen:       "Quan Tri Vien",
		SoDienThoai: "0123456789",
		Email:       "admin@watchshop.vn",
		TenDangNhap: "admin01",
		MatKhau:     hash,
		VaiTro:      "admin",
		TrangThai:   true,
	}
	db.Create(&admin)

	userHash, _ := utils.HashPassword("User@123")
	user := models.User{
		HoTen:       "Nguyen Van A",
		SoDienThoai: "0987654321",
		Email:       "user@watchshop.vn",
		TenDangNhap: "user01",
		MatKhau:     userHash,
		VaiTro:      "user",
		TrangThai:   true,
	}
	db.Create(&user)

	cats := []models.Category{
		{TenDanhMuc: "Dong ho nam", MoTa: "Dong ho danh cho nam gioi", Slug: "dong-ho-nam", TrangThai: true},
		{TenDanhMuc: "Dong ho nu", MoTa: "Dong ho danh cho nu gioi", Slug: "dong-ho-nu", TrangThai: true},
		{TenDanhMuc: "Dong ho doi", MoTa: "Dong ho cap doi", Slug: "dong-ho-doi", TrangThai: true},
		{TenDanhMuc: "Dong ho the thao", MoTa: "Dong ho the thao", Slug: "dong-ho-the-thao", TrangThai: true},
	}
	db.Create(&cats)

	brands := []models.Brand{
		{TenThuongHieu: "Casio", QuocGia: "Nhat Ban", MoTa: "Thuong hieu Nhat Ban noi tieng", TrangThai: true},
		{TenThuongHieu: "Citizen", QuocGia: "Nhat Ban", MoTa: "Thuong hieu lau doi cua Nhat", TrangThai: true},
		{TenThuongHieu: "Seiko", QuocGia: "Nhat Ban", MoTa: "Seiko - Nhat Ban", TrangThai: true},
		{TenThuongHieu: "Orient", QuocGia: "Nhat Ban", MoTa: "Dong ho co Orient", TrangThai: true},
		{TenThuongHieu: "Rolex", QuocGia: "Thuy Si", MoTa: "Dong ho cao cap Thuy Si", TrangThai: true},
		{TenThuongHieu: "Tissot", QuocGia: "Thuy Si", MoTa: "Tissot - Thuy Si", TrangThai: true},
	}
	db.Create(&brands)

	products := []models.Product{
		{
			TenSanPham: "Casio MTP-V001GL-7B", MoTa: "Dong ho Casio classic cho nam, day da nau, mat trang.",
			Gia: 990000, GiaGoc: 1200000, SoLuong: 50, CategoryID: cats[0].ID, BrandID: brands[0].ID,
			LoaiMay: "Quartz", DayDeo: "Da", MatKinh: "Mineral", DuongKinh: "38mm", KhangNuoc: "3ATM",
			GioiTinh: "Nam", XuatXu: "Thai Lan", BaoHanh: "12 thang",
			HinhAnh: "https://images.unsplash.com/photo-1524592094714-0f0654e20314?w=500",
			NoiBat: true, TrangThai: true,
		},
		{
			TenSanPham: "Citizen NH8350-08A", MoTa: "Citizen tu dong, day kim loai, lich ngay.",
			Gia: 4500000, GiaGoc: 5200000, SoLuong: 20, CategoryID: cats[0].ID, BrandID: brands[1].ID,
			LoaiMay: "Automatic", DayDeo: "Kim loai", MatKinh: "Sapphire", DuongKinh: "41mm", KhangNuoc: "5ATM",
			GioiTinh: "Nam", XuatXu: "Nhat Ban", BaoHanh: "24 thang",
			HinhAnh: "https://images.unsplash.com/photo-1547996160-81dfa63595aa?w=500",
			NoiBat: true, TrangThai: true,
		},
		{
			TenSanPham: "Seiko SUR306P1", MoTa: "Seiko classic, day kim loai, mat xanh sang trong.",
			Gia: 6200000, GiaGoc: 7000000, SoLuong: 15, CategoryID: cats[0].ID, BrandID: brands[2].ID,
			LoaiMay: "Quartz", DayDeo: "Kim loai", MatKinh: "Sapphire", DuongKinh: "40mm", KhangNuoc: "10ATM",
			GioiTinh: "Nam", XuatXu: "Nhat Ban", BaoHanh: "24 thang",
			HinhAnh: "https://images.unsplash.com/photo-1542496658-e33a6d0d50f6?w=500",
			NoiBat: false, TrangThai: true,
		},
		{
			TenSanPham: "Orient Bambino RA-AC0M04Y", MoTa: "Orient Bambino dong ho co dien cuc dep.",
			Gia: 4800000, GiaGoc: 5500000, SoLuong: 18, CategoryID: cats[0].ID, BrandID: brands[3].ID,
			LoaiMay: "Automatic", DayDeo: "Da", MatKinh: "Mineral", DuongKinh: "38mm", KhangNuoc: "3ATM",
			GioiTinh: "Nam", XuatXu: "Nhat Ban", BaoHanh: "24 thang",
			HinhAnh: "https://images.unsplash.com/photo-1622434641406-a158123450f9?w=500",
			NoiBat: true, TrangThai: true,
		},
		{
			TenSanPham: "Casio LTP-V002L-1B", MoTa: "Casio dong ho nu day da, gon nhe.",
			Gia: 850000, GiaGoc: 1000000, SoLuong: 35, CategoryID: cats[1].ID, BrandID: brands[0].ID,
			LoaiMay: "Quartz", DayDeo: "Da", MatKinh: "Mineral", DuongKinh: "30mm", KhangNuoc: "3ATM",
			GioiTinh: "Nu", XuatXu: "Thai Lan", BaoHanh: "12 thang",
			HinhAnh: "https://images.unsplash.com/photo-1606293459243-cb9c6df1ce75?w=500",
			NoiBat: false, TrangThai: true,
		},
		{
			TenSanPham: "Citizen EM0500-73A", MoTa: "Citizen Eco Drive cho nu, nang luong anh sang.",
			Gia: 5800000, GiaGoc: 6500000, SoLuong: 12, CategoryID: cats[1].ID, BrandID: brands[1].ID,
			LoaiMay: "Eco-Drive", DayDeo: "Kim loai", MatKinh: "Sapphire", DuongKinh: "29mm", KhangNuoc: "5ATM",
			GioiTinh: "Nu", XuatXu: "Nhat Ban", BaoHanh: "24 thang",
			HinhAnh: "https://images.unsplash.com/photo-1639037687665-37e2ec068a31?w=500",
			NoiBat: true, TrangThai: true,
		},
		{
			TenSanPham: "Tissot T101.210.16.116.00", MoTa: "Tissot dong ho nu Thuy Si sang trong.",
			Gia: 8900000, GiaGoc: 9800000, SoLuong: 8, CategoryID: cats[1].ID, BrandID: brands[5].ID,
			LoaiMay: "Quartz", DayDeo: "Kim loai", MatKinh: "Sapphire", DuongKinh: "32mm", KhangNuoc: "5ATM",
			GioiTinh: "Nu", XuatXu: "Thuy Si", BaoHanh: "24 thang",
			HinhAnh: "https://images.unsplash.com/photo-1548171915-e5cb7b6f6e10?w=500",
			NoiBat: true, TrangThai: true,
		},
		{
			TenSanPham: "Casio G-Shock GA-2100-1A", MoTa: "Casio G-Shock the thao manh me, ben bi.",
			Gia: 3200000, GiaGoc: 3800000, SoLuong: 40, CategoryID: cats[3].ID, BrandID: brands[0].ID,
			LoaiMay: "Quartz", DayDeo: "Cao su", MatKinh: "Mineral", DuongKinh: "44mm", KhangNuoc: "20ATM",
			GioiTinh: "Nam", XuatXu: "Thai Lan", BaoHanh: "12 thang",
			HinhAnh: "https://images.unsplash.com/photo-1434056886845-dac89ffe9b56?w=500",
			NoiBat: true, TrangThai: true,
		},
	}
	db.Create(&products)

	log.Println("Seed du lieu mau thanh cong")
	log.Println("Tai khoan admin: admin01 / Admin@123")
	log.Println("Tai khoan user:  user01 / User@123")
}
