package handler

import (
	"github.com/bondzai/portfolio-backend/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	service ports.CertService
}

func NewHttpHandler(service ports.CertService) *httpHandler {
	return &httpHandler{
		service: service,
	}
}

func (h *httpHandler) GetCerts(c *fiber.Ctx) error {
	data, err := h.service.ReadCerts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(data)
}
