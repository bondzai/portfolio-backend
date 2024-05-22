package ports

import (
	"github.com/bondzai/portfolio-backend/internal/core/models"
)

// primary port for input data
type CertService interface {
	ReadCerts() ([]*models.Certification, error)
}

// secondary port for output data
type CertRepo interface {
	ReadCerts() ([]*models.Certification, error)
}
