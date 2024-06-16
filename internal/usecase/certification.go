package usecase

import (
	"slices"

	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/bondzai/portfolio-backend/internal/repository"
)

type CertService interface {
	ReadCerts() ([]domain.Certification, error)
}

type certService struct {
	repo repository.MongoDBClientInterface
}

func NewCertService(repo repository.MongoDBClientInterface) *certService {
	return &certService{
		repo: repo,
	}
}

func (u *certService) ReadCerts() ([]domain.Certification, error) {
	data, err := u.repo.ReadCerts()
	if err != nil {
		return []domain.Certification{}, err
	}

	slices.Reverse(data)

	return data, nil
}
