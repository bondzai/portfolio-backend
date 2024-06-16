package usecase

import (
	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/bondzai/portfolio-backend/internal/repository"
)

type SkillService interface {
	ReadSkills() ([]domain.Skill, error)
}

type skillService struct {
	repo repository.MongoDBClientInterface
}

func NewSkillService(repo repository.MongoDBClientInterface) *skillService {
	return &skillService{
		repo: repo,
	}
}

func (u *skillService) ReadSkills() ([]domain.Skill, error) {
	data, err := u.repo.ReadSkills()
	if err != nil {
		return []domain.Skill{}, err
	}

	return data, nil
}
