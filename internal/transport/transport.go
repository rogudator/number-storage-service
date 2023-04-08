package transport

import (
	"github.com/rogudator/number-storage-service/internal/service"
	"github.com/rogudator/number-storage-service/proto/number_storage"
)

type Transport struct {
	services *service.Service
	number_storage.NumberStorageServiceServer
}

func NewTransport(services *service.Service) *Transport {
	return &Transport{
		services: services,
	}
}