package handlers

import (
	"strconv"

	"watch-shop/backend/internal/models"
	"watch-shop/backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	db *gorm.DB
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{db: db}
}

type productInput struct {
	TenSanPham string  `json:"ten_san_pham"`
	MoTa       string  `json:"mo_ta"`
	Gia        float64 `json:"gia"`
	GiaGoc     float64 `json:"gia_goc"`
	SoLuong    int     `json:"so_luong"`
	HinhAnh    string  `json:"hinh_anh"`
	CategoryID uint    `json:"category_id"`
	BrandID    uint    `json:"brand_id"`
	LoaiMay    string  `json:"loai_may"`
	DayDeo     string  `json:"day_deo"`
	MatKinh    string  `json:"mat_kinh"`
	DuongKinh  string  `json:"duong_kinh"`
	KhangNuoc  string  `json:"khang_nuoc"`
	GioiTinh   string  `json:"gioi_tinh"`
	XuatXu     string  `json:"xuat_xu"`
	BaoHanh    string  `json:"bao_hanh"`
	NoiBat     *bool   `json:"noi_bat"`
	TrangThai  *bool   `json:"trang_thai"`
	HinhAnhs   []string `json:"hinh_anhs"`
}

func (h *ProductHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 12
	}

	q := h.db.Model(&models.Product{}).Preload("Category").Preload("Brand").Preload("HinhAnhs")

	if c.Query("admin") != "1" {
		q = q.Where("trang_thai = ?", true)
	}
	if kw := c.Query("keyword"); kw != "" {
		q = q.Where("ten_san_pham ILIKE ?", "%"+kw+"%")
	}
	if v := c.Query("category_id"); v != "" {
		q = q.Where("category_id = ?", v)
	}
	if v := c.Query("brand_id"); v != "" {
		q = q.Where("brand_id = ?", v)
	}
	if v := c.Query("loai_may"); v != "" {
		q = q.Where("loai_may = ?", v)
	}
	if v := c.Query("day_deo"); v != "" {
		q = q.Where("day_deo = ?", v)
	}
	if v := c.Query("gioi_tinh"); v != "" {
		q = q.Where("gioi_tinh = ?", v)
	}
	if v := c.Query("gia_min"); v != "" {
		q = q.Where("gia >= ?", v)
	}
	if v := c.Query("gia_max"); v != "" {
		q = q.Where("gia <= ?", v)
	}
	if c.Query("noi_bat") == "1" {
		q = q.Where("noi_bat = ?", true)
	}

	sort := c.DefaultQuery("sort", "moi_nhat")
	switch sort {
	case "gia_tang":
		q = q.Order("gia ASC")
	case "gia_giam":
		q = q.Order("gia DESC")
	case "ban_chay":
		q = q.Order("da_ban DESC")
	default:
		q = q.Order("id DESC")
	}

	var total int64
	q.Count(&total)

	var items []models.Product
	q.Limit(size).Offset((page - 1) * size).Find(&items)
	utils.Page(c, items, total, page, size)
}

func (h *ProductHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p models.Product
	if err := h.db.Preload("Category").Preload("Brand").Preload("HinhAnhs").First(&p, id).Error; err != nil {
		utils.NotFound(c, "Khong tim thay san pham")
		return
	}
	utils.OK(c, "OK", p)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var in productInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	if in.TenSanPham == "" || in.Gia <= 0 {
		utils.BadRequest(c, "Vui long nhap ten va gia san pham")
		return
	}
	p := models.Product{
		TenSanPham: in.TenSanPham,
		MoTa:       in.MoTa,
		Gia:        in.Gia,
		GiaGoc:     in.GiaGoc,
		SoLuong:    in.SoLuong,
		HinhAnh:    in.HinhAnh,
		CategoryID: in.CategoryID,
		BrandID:    in.BrandID,
		LoaiMay:    in.LoaiMay,
		DayDeo:     in.DayDeo,
		MatKinh:    in.MatKinh,
		DuongKinh:  in.DuongKinh,
		KhangNuoc:  in.KhangNuoc,
		GioiTinh:   in.GioiTinh,
		XuatXu:     in.XuatXu,
		BaoHanh:    in.BaoHanh,
		TrangThai:  true,
	}
	if in.NoiBat != nil {
		p.NoiBat = *in.NoiBat
	}
	if in.TrangThai != nil {
		p.TrangThai = *in.TrangThai
	}
	if err := h.db.Create(&p).Error; err != nil {
		utils.ServerError(c, "Khong the them san pham")
		return
	}
	for i, url := range in.HinhAnhs {
		h.db.Create(&models.ProductImage{ProductID: p.ID, URL: url, ThuTu: i})
	}
	utils.Created(c, "Them san pham thanh cong", p)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p models.Product
	if err := h.db.First(&p, id).Error; err != nil {
		utils.NotFound(c, "Khong tim thay san pham")
		return
	}
	var in productInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	if in.TenSanPham != "" {
		p.TenSanPham = in.TenSanPham
	}
	p.MoTa = in.MoTa
	if in.Gia > 0 {
		p.Gia = in.Gia
	}
	p.GiaGoc = in.GiaGoc
	p.SoLuong = in.SoLuong
	if in.HinhAnh != "" {
		p.HinhAnh = in.HinhAnh
	}
	if in.CategoryID != 0 {
		p.CategoryID = in.CategoryID
	}
	if in.BrandID != 0 {
		p.BrandID = in.BrandID
	}
	p.LoaiMay = in.LoaiMay
	p.DayDeo = in.DayDeo
	p.MatKinh = in.MatKinh
	p.DuongKinh = in.DuongKinh
	p.KhangNuoc = in.KhangNuoc
	p.GioiTinh = in.GioiTinh
	p.XuatXu = in.XuatXu
	p.BaoHanh = in.BaoHanh
	if in.NoiBat != nil {
		p.NoiBat = *in.NoiBat
	}
	if in.TrangThai != nil {
		p.TrangThai = *in.TrangThai
	}
	h.db.Save(&p)

	if in.HinhAnhs != nil {
		h.db.Where("product_id = ?", p.ID).Delete(&models.ProductImage{})
		for i, url := range in.HinhAnhs {
			h.db.Create(&models.ProductImage{ProductID: p.ID, URL: url, ThuTu: i})
		}
	}
	utils.OK(c, "Cap nhat thanh cong", p)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.db.Where("product_id = ?", id).Delete(&models.ProductImage{})
	if err := h.db.Delete(&models.Product{}, id).Error; err != nil {
		utils.ServerError(c, "Xoa that bai")
		return
	}
	utils.OK(c, "Xoa thanh cong", nil)
}
