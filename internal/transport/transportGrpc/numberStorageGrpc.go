package transportGrpc

import (
	"context"
	"number-storage-service/proto/number_storage"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (gRPC *TransportGrpc) GetNumber(ctx context.Context, req *emptypb.Empty) (*number_storage.NumberResponse, error) {
	number, err := gRPC.services.GetNumber()
	if err != nil {
		return nil, err
	}

	res := number_storage.NumberResponse{
		Number: number,
	}

	return &res, nil
}

func (gRPC *TransportGrpc) UpdateNumber(ctx context.Context, req *number_storage.NumberRequest) (*number_storage.NumberResponse, error) {
	number := req.GetNumber()

	err := gRPC.services.UpdateNumber(number)
	if err != nil {
		return nil, err
	}

	return &number_storage.NumberResponse{
		Number: number,
	}, nil
}