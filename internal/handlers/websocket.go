package handlers

import (
	"log"

	"github.com/bondzai/portfolio-backend/internal/usecases"
	"github.com/gofiber/websocket/v2"
)

type WsHandler struct {
	wsService *usecases.WsService
}

func NewWsHandler(wsService *usecases.WsService) *WsHandler {
	return &WsHandler{
		wsService: wsService,
	}
}

func (h *WsHandler) HandleConnection(c *websocket.Conn) {
	h.wsService.AddConnection(c)
	defer h.wsService.RemoveConnection(c)

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
	}
}
