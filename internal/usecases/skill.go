package usecases

import (
	"github.com/bondzai/portfolio-backend/internal/models"
	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/bondzai/portfolio-backend/internal/utils/errs"
)

type SkillService interface {
	ReadSkills() ([]models.Skill, error)
}

type skillService struct {
	repo repositories.MongoDBClient
}

func NewSkillService(repo repositories.MongoDBClient) *skillService {
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
