package service

import (
	"pet"
	"pet/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Sueta interface {
	Create(sueta pet.Sueta) (int, int, error)
	FindAll() ([]pet.Sueta, error)
}

type Service struct {
	Sueta
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Sueta: NewSuetaService(repos.Sueta),
	}
}
