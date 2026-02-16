package usecase

import (
	"thinkdrop-backend/internal/modules/admin/domain"
	"thinkdrop-backend/internal/modules/admin/websocket"
)

type AdminService struct {
	repo domain.AdminRepo
	hub  *websocket.Hub
}

func NewAdminService(r domain.AdminRepo, h *websocket.Hub) *AdminService {
	return &AdminService{repo: r, hub: h}
}
