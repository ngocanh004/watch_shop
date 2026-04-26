package handlers

import (
	"strconv"

	"watch-shop/backend/internal/models"
	"watch-shop/backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 10
	}

	q := h.db.Model(&models.User{})
	if kw := c.Query("keyword"); kw != "" {
		like := "%" + kw + "%"
		q = q.Where("ho_ten ILIKE ? OR ten_dang_nhap ILIKE ? OR email ILIKE ?", like, like, like)
	}
	if r := c.Query("vai_tro"); r != "" {
		q = q.Where("vai_tro = ?", r)
	}

	var total int64
	q.Count(&total)

	var items []models.User
	q.Order("id DESC").Limit(size).Offset((page - 1) * size).Find(&items)
	utils.Page(c, items, total, page, size)
}

func (h *UserHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var u models.User
	if err := h.db.First(&u, id).Error; err != nil {
		utils.NotFound(c, "Khong tim thay nguoi dung")
		return
	}
	utils.OK(c, "OK", u)
}

type userUpdateInput struct {
	HoTen       string `json:"ho_ten"`
	SoDienThoai string `json:"so_dien_thoai"`
	Email       string `json:"email"`
	DiaChi      string `json:"dia_chi"`
	VaiTro      string `json:"vai_tro"`
	TrangThai   *bool  `json:"trang_thai"`
}

func (h *UserHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var u models.User
	if err := h.db.First(&u, id).Error; err != nil {
		utils.NotFound(c, "Khong tim thay nguoi dung")
		return
	}
	var in userUpdateInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	if in.HoTen != "" {
		u.HoTen = in.HoTen
	}
	if in.SoDienThoai != "" {
		u.SoDienThoai = in.SoDienThoai
	}
	if in.Email != "" {
		u.Email = in.Email
	}
	if in.DiaChi != "" {
		u.DiaChi = in.DiaChi
	}
	if in.VaiTro != "" {
		u.VaiTro = in.VaiTro
	}
	if in.TrangThai != nil {
		u.TrangThai = *in.TrangThai
	}
	h.db.Save(&u)
	utils.OK(c, "Cap nhat thanh cong", u)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.db.Delete(&models.User{}, id).Error; err != nil {
		utils.ServerError(c, "Xoa that bai")
		return
	}
	utils.OK(c, "Xoa thanh cong", nil)
}

type resetPasswordInput struct {
	MatKhauMoi string `json:"mat_khau_moi"`
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var in resetPasswordInput
	if err := c.ShouldBindJSON(&in); err != nil || !utils.ValidatePassword(in.MatKhauMoi) {
		utils.BadRequest(c, "Mat khau moi khong hop le (8-20 ky tu, chua chu hoa, chu thuong, so, ky tu dac biet)")
		return
	}
	hash, _ := utils.HashPassword(in.MatKhauMoi)
	if err := h.db.Model(&models.User{}).Where("id = ?", id).Update("mat_khau", hash).Error; err != nil {
		utils.ServerError(c, "Cap nhat that bai")
		return
	}
	utils.OK(c, "Cap nhat thanh cong", nil)
}
