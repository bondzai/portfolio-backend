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

func (s *skillService) ReadSkills() ([]models.Skill, error) {
	data, err := s.repo.ReadSkills()
	if err != nil {
		return []models.Skill{}, err
	}

	return data, nil
}
