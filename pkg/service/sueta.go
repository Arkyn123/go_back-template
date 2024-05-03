package service

import (
	"net/http"
	"pet"
	"pet/pkg/repository"
)

type SuetaService struct {
	repo repository.Sueta
}

func NewSuetaService(repo repository.Sueta) *SuetaService {
	return &SuetaService{repo: repo}
}

func (r *SuetaService) Create(sueta pet.Sueta) (int, int, error) {

	id, status, err := r.repo.Create(sueta)
	if err != nil {
		return 0, status, err
	}

	return id, http.StatusOK, nil
}

func (r *SuetaService) FindAll() ([]pet.Sueta, error) {

	suetas, err := r.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return suetas, nil
}
