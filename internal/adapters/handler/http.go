package handler

import (
	"github.com/bondzai/portfolio-backend/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	cs ports.CertService
	ps ports.ProjectService
	ss ports.SkillService
}

func NewHttpHandler(cs ports.CertService, ps ports.ProjectService, ss ports.SkillService) *httpHandler {
	return &httpHandler{
		cs: cs,
		ps: ps,
		ss: ss,
	}
}

func (h *httpHandler) GetCerts(c *fiber.Ctx) error {
	data, err := h.cs.ReadCerts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(data)
}

func (h *httpHandler) GetSkills(c *fiber.Ctx) error {
	data, err := h.ss.ReadSkills()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(data)
}
