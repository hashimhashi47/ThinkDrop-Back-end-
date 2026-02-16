package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"thinkdrop-backend/internal/modules/chat/usecase"
	ws "thinkdrop-backend/internal/modules/chat/websocket"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type ChatHandler struct {
	Service *usecase.ChatService
	Hub     *ws.Hub
}

func NewChatHandler(service *usecase.ChatService, hub *ws.Hub) *ChatHandler {
	return &ChatHandler{
		Service: service,
		Hub:     hub,
	}
}

type IncomingMessage struct {
	ReceiverID uint   `json:"receiver_id"`
	Content    string `json:"content"`
}

func (h *ChatHandler) HandleWS(c *websocket.Conn) {

	userIDStr := c.Query("user_id")
	userID64, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		log.Println("invalid user_id")
		return
	}
	userID := uint(userID64)

	client := &ws.Client{
		UserID: userID,
		Conn:   c,
		Send:   make(chan interface{}, 256),
		Hub:    h.Hub,
	}

	h.Hub.Register(client)

	go client.WritePump()

	client.ReadPump(func(senderID uint, data []byte) {

		var msg IncomingMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			return
		}

		savedMsg, err := h.Service.SendMessage(senderID, msg.ReceiverID, msg.Content)
		if err != nil {
			return
		}

		// send to receiver
		h.Hub.SendToUser(msg.ReceiverID, savedMsg)

		// send back to sender
		h.Hub.SendToUser(senderID, savedMsg)
	})
}

func (h *ChatHandler) GetChats(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	data, err := h.Service.Getallchat(UserID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})
}
