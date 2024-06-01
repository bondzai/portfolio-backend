package services

import (
	"github.com/bondzai/portfolio-backend/internal/core/models"
	"github.com/bondzai/portfolio-backend/internal/core/ports"
)

type projectService struct {
	repo ports.ProjectRepo
}

func NewProjectService(repo ports.ProjectRepo) *projectService {
	return &projectService{
		repo: repo,
	}
}

func (m *projectService) ReadProjects() ([]*models.Project, error) {
	return m.repo.ReadProjects()
}
