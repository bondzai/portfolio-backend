package usecase

import (
	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/bondzai/portfolio-backend/internal/utils/errs"
)

type SkillService interface {
	ReadSkills() ([]domain.Skill, error)
}

type skillService struct {
	repo repositories.MongoDBClient
}

func NewSkillService(repo repositories.MongoDBClient) *skillService {
	return &skillService{
		repo: repo,
	}
}

func (u *skillService) ReadSkills() ([]domain.Skill, error) {
	data, err := u.repo.ReadSkills()
	if err != nil {
		return []domain.Skill{}, errs.NewUnExpectedError()
	}

	return data, nil
}
