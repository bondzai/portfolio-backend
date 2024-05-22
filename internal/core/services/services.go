package services

import (
	"github.com/bondzai/portfolio-backend/internal/core/models"
	"github.com/bondzai/portfolio-backend/internal/core/ports"
)

type CertService struct {
	repo ports.CertRepo
}

func NewCertService(repo ports.CertRepo) *CertService {
	return &CertService{
		repo: repo,
	}
}

func (m *CertService) ReadCerts() ([]*models.Certification, error) {
	return m.repo.ReadCerts()
}
