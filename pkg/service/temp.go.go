package service

import (
	"temp/pkg/repository"
)

type TempService struct {
	repo repository.Temp
}

func NewTempService(repo repository.Temp) *TempService {
	return &TempService{repo: repo}
}

func (r *TempService) Temp() {

}
