package transportGrpc

import (
	"number-storage-service/internal/service"
	"number-storage-service/proto/number_storage"
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