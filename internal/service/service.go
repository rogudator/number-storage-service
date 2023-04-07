package service

import "number-storage-service/internal/repository"

type Service struct {
	NumberStorageService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NumberStorageService: *NewNumberStorageService(*repos),
	}
}