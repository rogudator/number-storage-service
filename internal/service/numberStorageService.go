package service

import "github.com/rogudator/number-storage-service/internal/repository"

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

// Service layer method just calls repo method for stored number.
func (s *NumberStorageService) GetNumber() (int64, error) {
	return s.repo.GetNumber()
}

// The logic behind changing number is on this service layer.
// In case logic will change, you will need to reprogram only this method.
func (s *NumberStorageService) UpdateNumber(number int64) (error) {
	// Gets current number
	currentNumber, err := s.repo.GetNumber()
	if err != nil {
		return err
	}
	// And adds to it recieved number.
	updatedNumber := currentNumber + number
	return s.repo.UpdateNumber(updatedNumber)
}