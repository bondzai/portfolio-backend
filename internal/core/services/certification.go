package services

import (
	"github.com/bondzai/portfolio-backend/internal/core/models"
	"github.com/bondzai/portfolio-backend/internal/core/ports"
)

type certService struct {
	repo ports.CertRepo
}

func NewCertService(repo ports.CertRepo) *certService {
	return &certService{
		repo: repo,
	}
}

func (m *certService) ReadCerts() ([]models.Certification, error) {
	return m.repo.ReadCerts()
}
