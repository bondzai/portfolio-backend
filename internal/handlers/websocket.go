package handlers

import (
	"log"

	"github.com/bondzai/portfolio-backend/internal/usecases"
	"github.com/gofiber/websocket/v2"
)

type (
	WsHandler interface {
		HandleConnection(c *websocket.Conn)
	}

	wsHandler struct {
		wsService usecases.WsService
	}
)

func NewWsHandler(wsService usecases.WsService) WsHandler {
	return &wsHandler{
		wsService: wsService,
	}
}

func (h *wsHandler) HandleConnection(c *websocket.Conn) {
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
