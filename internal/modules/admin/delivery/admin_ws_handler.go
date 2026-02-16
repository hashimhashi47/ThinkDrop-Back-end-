package delivery

import "github.com/gofiber/websocket/v2"




func (h *AdminController) Handle(c *websocket.Conn) {
	h.hub.AddClient(c)

	defer func() {
		h.hub.RemoveClient(c)
		c.Close()
	}()

	for {
		if _, _, err := c.ReadMessage(); err != nil {
			break
		}
	}
}
