package usecase

import (
	"slices"

	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/bondzai/portfolio-backend/internal/repository"
	"github.com/bondzai/portfolio-backend/internal/utils/errs"
)

type ProjectService interface {
	ReadProjects() ([]domain.Project, error)
}

type projectService struct {
	repo repository.MongoDBClientInterface
}

func NewProjectService(repo repository.MongoDBClientInterface) *projectService {
	return &projectService{
		repo: repo,
	}
}

func (u *projectService) ReadProjects() ([]domain.Project, error) {
	data, err := u.repo.ReadProjects()
	if err != nil {
		return []domain.Project{}, errs.NewUnExpectedError()
	}

	slices.Reverse(data)

	return data, nil
}
