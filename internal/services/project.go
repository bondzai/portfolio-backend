package services

import (
	"slices"

	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/bondzai/portfolio-backend/internal/repository"
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

func (s *projectService) ReadProjects() ([]domain.Project, error) {
	data, err := s.repo.ReadProjects()
	if err != nil {
		return []domain.Project{}, err
	}

	slices.Reverse(data)

	return data, nil
}
