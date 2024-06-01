package services

import (
	"github.com/bondzai/portfolio-backend/internal/core/models"
	"github.com/bondzai/portfolio-backend/internal/core/ports"
)

type skillService struct {
	repo ports.SkillRepo
}

func NewSkillService(repo ports.SkillRepo) *skillService {
	return &skillService{
		repo: repo,
	}
}

func (m *skillService) ReadSkills() ([]models.Skill, error) {
	return m.repo.ReadSkills()
}
