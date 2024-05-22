package services

import (
	"github.com/bondzai/portfolio-backend/internal/core/models"
	"github.com/bondzai/portfolio-backend/internal/core/ports"
)

type ProjectService struct {
	repo ports.ProjectRepo
}

func NewProjectService(repo ports.ProjectRepo) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (m *ProjectService) ReadProjects() ([]*models.Project, error) {
	return m.repo.ReadProjects()
}
