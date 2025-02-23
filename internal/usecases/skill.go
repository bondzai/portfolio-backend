package usecases

import (
	"github.com/bondzai/portfolio-backend/internal/models"
	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/bondzai/portfolio-backend/internal/utils/errs"
)

type (
	SkillService interface {
		ReadSkills() ([]models.Skill, error)
	}

	skillService struct {
		repo repositories.MongoDBClient
	}
)

func NewSkillService(repo repositories.MongoDBClient) SkillService {
	return &skillService{
		repo: repo,
	}
}

func (u *skillService) ReadSkills() ([]models.Skill, error) {
	data, err := u.repo.ReadSkills()
	if err != nil {
		return []models.Skill{}, errs.NewUnExpectedError()
	}

	return data, nil
}
