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

func (h *ChatHandler) HandleWS(c *websocket.Conn,UserID uint) {
	log.Println(UserID)
	client := &ws.Client{
		UserID: UserID,
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

func (h *ChatHandler) GetChatMessages(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	convoIDStr := c.Params("id")
	convoID64, err := strconv.ParseUint(convoIDStr, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid conversation id"})
	}
	convoID := uint(convoID64)

	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	data, err := h.Service.GetMessages(userID, convoID, limit, offset)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

func (h *ChatHandler) StartConversation(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var body struct {
		ReceiverID uint `json:"receiver_id"`
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	convo, err := h.Service.StartConversation(userID, body.ReceiverID)
	if err != nil {
		return err
	}

	return c.JSON(convo)
}
