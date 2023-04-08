package transport

import (
	"context"
	"github.com/rogudator/number-storage-service/proto/number_storage"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *Transport) GetNumber(ctx context.Context, req *emptypb.Empty) (*number_storage.NumberResponse, error) {
	number, err := t.services.GetNumber()
	if err != nil {
		return nil, err
	}

	res := number_storage.NumberResponse{
		Number: number,
	}

	return &res, nil
}

func (t *Transport) UpdateNumber(ctx context.Context, req *number_storage.NumberRequest) (*emptypb.Empty, error) {
	number := req.GetNumber()

	err := t.services.UpdateNumber(number)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}