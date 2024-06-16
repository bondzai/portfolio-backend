package usecase

import (
	"slices"

	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/bondzai/portfolio-backend/internal/repository"
	"github.com/bondzai/portfolio-backend/internal/utils/errs"
)

type CertService interface {
	ReadCerts() ([]domain.Certification, error)
}

type certService struct {
	repo repository.MongoDBClient
}

func NewCertService(repo repository.MongoDBClient) *certService {
	return &certService{
		repo: repo,
	}
}

func (u *certService) ReadCerts() ([]domain.Certification, error) {
	data, err := u.repo.ReadCerts()
	if err != nil {
		return []domain.Certification{}, errs.NewUnExpectedError()
	}

	slices.Reverse(data)

	return data, nil
}
