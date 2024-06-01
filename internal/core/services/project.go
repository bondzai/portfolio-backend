package services

import (
	"slices"

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

func (s *projectService) ReadProjects() ([]models.Project, error) {
	data, err := s.repo.ReadProjects()
	if err != nil {
		return []models.Project{}, err
	}

	slices.Reverse(data)

	return data, nil
}
