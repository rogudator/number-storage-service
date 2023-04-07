package transportGrpc

import (
	"github.com/rogudator/number-storage-service/internal/service"
	"github.com/rogudator/number-storage-service/proto/number_storage"
)

type TransportGrpc struct {
	services *service.Service
	number_storage.NumberStorageServiceServer
}

func NewTransportGrpc(services *service.Service) *TransportGrpc {
	return &TransportGrpc{
		services: services,
	}
}