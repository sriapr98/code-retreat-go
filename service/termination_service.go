package service

import (
	"code-retreat-go/models"
	"code-retreat-go/repository"
)

type TerminationService interface {
	GetAllTerminations() []models.Termination
}

type terminationService struct {
	terminationRepository repository.TerminationRepository
}

func (e terminationService) GetAllTerminations() []models.Termination {
	return e.terminationRepository.GetAllTerminations()
}

func NewTerminationService(terminationRepository repository.TerminationRepository) TerminationService {
	return terminationService{terminationRepository}
}
