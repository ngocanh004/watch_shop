package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"watch-shop/backend/internal/config"
	"watch-shop/backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadHandler struct {
	cfg *config.Config
}

func NewUploadHandler(cfg *config.Config) *UploadHandler {
	return &UploadHandler{cfg: cfg}
}

var allowedExt = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".webp": true, ".gif": true,
}

func (h *UploadHandler) Image(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.BadRequest(c, "Khong tim thay file")
		return
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExt[ext] {
		utils.BadRequest(c, "Dinh dang file khong duoc ho tro")
		return
	}
	if file.Size > 5*1024*1024 {
		utils.BadRequest(c, "File qua lon (toi da 5MB)")
		return
	}

	now := time.Now()
	subDir := fmt.Sprintf("%d/%02d", now.Year(), now.Month())
	dir := filepath.Join(h.cfg.UploadDir, subDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		utils.ServerError(c, "Khong the tao thu muc luu")
		return
	}
	name := uuid.New().String() + ext
	dst := filepath.Join(dir, name)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		utils.ServerError(c, "Luu file that bai")
		return
	}

	urlPath := fmt.Sprintf("/uploads/%s/%s", subDir, name)
	utils.OK(c, "Upload thanh cong", gin.H{
		"url":      urlPath,
		"full_url": h.cfg.PublicURL + urlPath,
	})
}
