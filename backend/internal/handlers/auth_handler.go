package handlers

import (
	"strings"

	"watch-shop/backend/internal/config"
	"watch-shop/backend/internal/dto"
	"watch-shop/backend/internal/models"
	"watch-shop/backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthHandler(db *gorm.DB, cfg *config.Config) *AuthHandler {
	return &AuthHandler{db: db, cfg: cfg}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}

	req.HoTen = strings.TrimSpace(req.HoTen)
	req.TenDangNhap = strings.TrimSpace(req.TenDangNhap)

	if req.HoTen == "" || req.SoDienThoai == "" || req.TenDangNhap == "" ||
		req.MatKhau == "" || req.XacNhanMatKhau == "" {
		utils.BadRequest(c, "Vui long nhap day du thong tin!")
		return
	}
	if len(req.HoTen) > 50 {
		utils.BadRequest(c, "Ho ten chi duoc chua chu va khong qua 50 ky tu")
		return
	}
	if !utils.ValidatePhone(req.SoDienThoai) {
		utils.BadRequest(c, "So dien thoai khong hop le. Dinh dang: 0********* (10 so)")
		return
	}
	if !utils.ValidateUsername(req.TenDangNhap) {
		utils.BadRequest(c, "Ten dang nhap phai chua ca chu va so. (Khong chua ky tu dac biet)")
		return
	}
	if !utils.ValidatePassword(req.MatKhau) {
		utils.BadRequest(c, "Mat khau phai gom chu hoa, chu thuong, so, ky tu dac biet (8-20 ky tu)")
		return
	}
	if req.MatKhau != req.XacNhanMatKhau {
		utils.BadRequest(c, "Mat khau xac nhan khong khop")
		return
	}
	if req.Email != "" && !utils.ValidateEmail(req.Email) {
		utils.BadRequest(c, "Email khong hop le")
		return
	}

	var count int64
	h.db.Model(&models.User{}).Where("ten_dang_nhap = ?", req.TenDangNhap).Count(&count)
	if count > 0 {
		utils.BadRequest(c, "Ten dang nhap da ton tai")
		return
	}

	hash, err := utils.HashPassword(req.MatKhau)
	if err != nil {
		utils.ServerError(c, "Khong the ma hoa mat khau")
		return
	}

	user := models.User{
		HoTen:       req.HoTen,
		SoDienThoai: req.SoDienThoai,
		Email:       req.Email,
		TenDangNhap: req.TenDangNhap,
		MatKhau:     hash,
		VaiTro:      "user",
		TrangThai:   true,
	}
	if err := h.db.Create(&user).Error; err != nil {
		utils.ServerError(c, "Khong the tao tai khoan")
		return
	}

	utils.Created(c, "Dang ky thanh cong! Vui long dang nhap.", gin.H{"id": user.ID})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	if req.TenDangNhap == "" || req.MatKhau == "" {
		utils.BadRequest(c, "Vui long nhap day du thong tin!")
		return
	}

	var user models.User
	if err := h.db.Where("ten_dang_nhap = ?", req.TenDangNhap).First(&user).Error; err != nil {
		utils.Unauthorized(c, "Ten dang nhap hoac mat khau khong dung!")
		return
	}
	if !user.TrangThai {
		utils.Forbidden(c, "Tai khoan da bi khoa")
		return
	}
	if !utils.CheckPassword(user.MatKhau, req.MatKhau) {
		utils.Unauthorized(c, "Ten dang nhap hoac mat khau khong dung!")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.VaiTro, h.cfg.JWTSecret, h.cfg.JWTExpireHours)
	if err != nil {
		utils.ServerError(c, "Khong the tao token")
		return
	}

	utils.OK(c, "Dang nhap thanh cong", gin.H{
		"token": token,
		"user":  user,
	})
}

func (h *AuthHandler) ChangePassword(c *gin.Context) {
	uid, _ := c.Get("user_id")
	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	if req.MatKhauCu == "" || req.MatKhauMoi == "" || req.XacNhanMatKhau == "" {
		utils.BadRequest(c, "Vui long nhap day du thong tin!")
		return
	}
	if !utils.ValidatePassword(req.MatKhauMoi) {
		utils.BadRequest(c, "Mat khau moi phai gom chu hoa, chu thuong, so, ky tu dac biet (8-20 ky tu)")
		return
	}
	if req.MatKhauMoi != req.XacNhanMatKhau {
		utils.BadRequest(c, "Mat khau xac nhan khong khop")
		return
	}

	var user models.User
	if err := h.db.First(&user, uid).Error; err != nil {
		utils.NotFound(c, "Khong tim thay tai khoan")
		return
	}
	if !utils.CheckPassword(user.MatKhau, req.MatKhauCu) {
		utils.BadRequest(c, "Mat khau cu khong chinh xac")
		return
	}

	hash, _ := utils.HashPassword(req.MatKhauMoi)
	h.db.Model(&user).Update("mat_khau", hash)
	utils.OK(c, "Cap nhat thanh cong!", nil)
}

func (h *AuthHandler) Profile(c *gin.Context) {
	uid, _ := c.Get("user_id")
	var user models.User
	if err := h.db.First(&user, uid).Error; err != nil {
		utils.NotFound(c, "Khong tim thay tai khoan")
		return
	}
	utils.OK(c, "OK", user)
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	uid, _ := c.Get("user_id")
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}

	updates := map[string]interface{}{}
	if req.HoTen != "" {
		updates["ho_ten"] = strings.TrimSpace(req.HoTen)
	}
	if req.SoDienThoai != "" {
		if !utils.ValidatePhone(req.SoDienThoai) {
			utils.BadRequest(c, "So dien thoai khong hop le")
			return
		}
		updates["so_dien_thoai"] = req.SoDienThoai
	}
	if req.Email != "" {
		if !utils.ValidateEmail(req.Email) {
			utils.BadRequest(c, "Email khong hop le")
			return
		}
		updates["email"] = req.Email
	}
	if req.DiaChi != "" {
		updates["dia_chi"] = req.DiaChi
	}

	if err := h.db.Model(&models.User{}).Where("id = ?", uid).Updates(updates).Error; err != nil {
		utils.ServerError(c, "Cap nhat that bai")
		return
	}
	utils.OK(c, "Cap nhat thanh cong!", nil)
}
