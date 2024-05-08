package service

import (
	"temp/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Temp interface {
	Temp()
}

type Service struct {
	Temp
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Temp: NewTempService(repos.Temp),
	}
}
