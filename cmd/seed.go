package main

import (
	"log/slog"

	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/bondzai/portfolio-backend/internal/utils"
)

func RunSeed() {
	mockRepo := repositories.NewMock()
	mongoRepo := initMongoDB()

	var err error

	certifications, _ := mockRepo.ReadCerts()
	err = mongoRepo.InsertMany("certifications", utils.ConvertToInterfaceSlice(certifications))
	if err != nil {
		slog.Error("Error seeded certifications data", err)
	} else {
		slog.Info("Successfully seeded certifications data to MongoDB")
	}

	projects, _ := mockRepo.ReadProjects()
	err = mongoRepo.InsertMany("projects", utils.ConvertToInterfaceSlice(projects))
	if err != nil {
		slog.Error("Error seeded projects data", err)
	} else {
		slog.Info("Successfully seeded projects data to MongoDB")
	}

	skills, _ := mockRepo.ReadSkills()
	err = mongoRepo.InsertMany("skills", utils.ConvertToInterfaceSlice(skills))
	if err != nil {
		slog.Error("Error seeded skills data", err)
	} else {
		slog.Info("Successfully seeded skills data to MongoDB")
	}
}
