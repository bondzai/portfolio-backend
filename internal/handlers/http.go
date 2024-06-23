package handlers

import (
	"github.com/bondzai/portfolio-backend/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	certService    usecases.CertService
	projectService usecases.ProjectService
	skillService   usecases.SkillService
	wakaService    usecases.WakaService
}

func NewHttpHandler(
	certService usecases.CertService,
	projectService usecases.ProjectService,
	skillService usecases.SkillService,
	wakaService usecases.WakaService,
) *httpHandler {
	return &httpHandler{
		certService:    certService,
		projectService: projectService,
		skillService:   skillService,
		wakaService:    wakaService,
	}
}

func (h *httpHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Ok")
}

func (h *httpHandler) GetCerts(c *fiber.Ctx) error {
	data, err := h.certService.ReadCerts()
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(data)
}

func (h *httpHandler) GetSkills(c *fiber.Ctx) error {
	data, err := h.skillService.ReadSkills()
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(data)
}

func (h *httpHandler) GetProjects(c *fiber.Ctx) error {
	data, err := h.projectService.ReadProjects()
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(data)
}

func (h *httpHandler) GetWakaStats(c *fiber.Ctx) error {
	data, err := h.wakaService.FetchDataFromAPI()
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(data)
}
