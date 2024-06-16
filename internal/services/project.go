package services

import (
	"slices"

	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/bondzai/portfolio-backend/internal/ports"
)

type projectService struct {
	repo ports.ProjectRepo
}

func NewProjectService(repo ports.ProjectRepo) *projectService {
	return &projectService{
		repo: repo,
	}
}

func (s *projectService) ReadProjects() ([]domain.Project, error) {
	data, err := s.repo.ReadProjects()
	if err != nil {
		return []domain.Project{}, err
	}

	slices.Reverse(data)

	return data, nil
}
