package services

import (
	"github.com/bondzai/portfolio-backend/internal/core/models"
	"github.com/bondzai/portfolio-backend/internal/core/ports"
)

type SkillService struct {
	repo ports.SkillRepo
}

func NewSkillService(repo ports.SkillRepo) *SkillService {
	return &SkillService{
		repo: repo,
	}
}

func (m *SkillService) ReadSkills() ([]*models.Skill, error) {
	return m.repo.ReadSkills()
}
