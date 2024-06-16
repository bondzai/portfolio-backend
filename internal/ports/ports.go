package ports

import (
	"github.com/bondzai/portfolio-backend/internal/domain"
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
