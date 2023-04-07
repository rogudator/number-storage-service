package service

import "number-storage-service/internal/repository"

type NumberStorageService struct {
	repo repository.Repository
}

func NewNumberStorageService(repos repository.Repository) *NumberStorageService {
	return &NumberStorageService{
		repo: repos,
	}
}

type NumberStorage interface {
	GetNumber() (int64, error)
	UpdateNumber(number int64) (int64, error)
}

func (s *NumberStorageService) GetNumber() (int64, error) {
	return s.repo.GetNumber()
}

func (s *NumberStorageService) UpdateNumber(number int64) (error) {
	currentNumber, err := s.repo.GetNumber()
	if err != nil {
		return err
	}
	updatedNumber := currentNumber + number
	return s.repo.UpdateNumber(updatedNumber)
}