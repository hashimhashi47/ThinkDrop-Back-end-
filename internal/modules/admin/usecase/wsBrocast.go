package usecase

import (
	"encoding/json"
	"thinkdrop-backend/internal/modules/admin/domain"
)

func (a *AdminService) Broadcast(module, eventType string, data interface{}) error {

	event := domain.AdminEvent{
		Module: module,
		Type:   eventType,
		Data:   data,
	}

	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	a.hub.Broadcast(jsonData)
	return nil
}
