package services

import (
	"slices"

	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/bondzai/portfolio-backend/internal/ports"
)

type certService struct {
	repo ports.CertRepo
}

func NewCertService(repo ports.CertRepo) *certService {
	return &certService{
		repo: repo,
	}
}

func (s *certService) ReadCerts() ([]domain.Certification, error) {
	data, err := s.repo.ReadCerts()
	if err != nil {
		return []domain.Certification{}, err
	}

	slices.Reverse(data)

	return data, nil
}
