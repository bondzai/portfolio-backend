package usecases

import (
	"slices"

	"github.com/bondzai/portfolio-backend/internal/models"
	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/bondzai/portfolio-backend/internal/utils/errs"
)

type (
	ProjectService interface {
		ReadProjects() ([]models.Project, error)
	}

	projectService struct {
		repo repositories.MongoDBClient
	}
)

func NewProjectService(repo repositories.MongoDBClient) *projectService {
	return &projectService{
		repo: repo,
	}
}

func (u *projectService) ReadProjects() ([]models.Project, error) {
	data, err := u.repo.ReadProjects()
	if err != nil {
		return []models.Project{}, errs.NewUnExpectedError()
	}

	slices.Reverse(data)

	return data, nil
}
