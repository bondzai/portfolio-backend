package ports

import (
	"github.com/bondzai/portfolio-backend/internal/core/models"
)

type (
	CertService interface {
		ReadCerts() ([]*models.Certification, error)
	}

	CertRepo interface {
		ReadCerts() ([]*models.Certification, error)
	}
)

type (
	ProjectService interface {
		ReadProjects() ([]*models.Project, error)
	}

	ProjectRepo interface {
		ReadProjects() ([]*models.Project, error)
	}
)

type (
	SkillService interface {
		ReadSkills() ([]*models.Skill, error)
	}

	SkillRepo interface {
		ReadSkills() ([]*models.Skill, error)
	}
)
