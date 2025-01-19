package usecases

import (
	"slices"

	"github.com/bondzai/portfolio-backend/internal/models"
	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/bondzai/portfolio-backend/internal/utils/errs"
)

type (
	CertService interface {
		ReadCerts() ([]models.Certification, error)
	}

	certService struct {
		repo repositories.MongoDBClient
	}
)

func NewCertService(repo repositories.MongoDBClient) CertService {
	return &certService{
		repo: repo,
	}
}

func (u *certService) ReadCerts() ([]models.Certification, error) {
	data, err := u.repo.ReadCerts()
	if err != nil {
		return []models.Certification{}, errs.NewUnExpectedError()
	}

	slices.Reverse(data)

	return data, nil
}
