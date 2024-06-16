package ports

import (
	"github.com/bondzai/portfolio-backend/internal/domain"
)

type (
	CertService interface {
		ReadCerts() ([]domain.Certification, error)
	}

	CertRepo interface {
		ReadCerts() ([]domain.Certification, error)
	}
)

type (
	ProjectService interface {
		ReadProjects() ([]domain.Project, error)
	}

	ProjectRepo interface {
		ReadProjects() ([]domain.Project, error)
	}
)

type (
	SkillService interface {
		ReadSkills() ([]domain.Skill, error)
	}

	SkillRepo interface {
		ReadSkills() ([]domain.Skill, error)
	}
)

type (
	WakaService interface {
		FetchDataFromAPI() (map[string]interface{}, error)
	}

	WakaRepo interface {
		FetchDataFromAPI() (map[string]interface{}, error)
	}
)
