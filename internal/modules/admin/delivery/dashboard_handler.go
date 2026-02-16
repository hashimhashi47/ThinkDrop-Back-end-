package delivery

import (
	"thinkdrop-backend/internal/modules/admin/usecase"
	"thinkdrop-backend/internal/modules/admin/websocket"
)

type AdminController struct {
	service *usecase.AdminService
	hub     *websocket.Hub
}

func NewAdminController(s *usecase.AdminService, h *websocket.Hub) *AdminController {
	return &AdminController{service: s, hub: h}
}

