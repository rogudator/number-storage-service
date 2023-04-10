package transport

import (
	"context"
	"errors"

	"github.com/rogudator/number-storage-service/proto/number_storage"

	"google.golang.org/protobuf/types/known/emptypb"
)

var ErrNumberIsNegative = errors.New("Number is negative")

// This handle receives call from REST or gRPC to get stored number. 
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

// This handle receives call from REST or gRPC to add recieved number to the current number.
func (t *Transport) UpdateNumber(ctx context.Context, req *number_storage.NumberRequest) (*emptypb.Empty, error) {
	number := req.GetNumber()

	// Negative numbers are not allowed.
	if number < 0 {
		return nil, ErrNumberIsNegative
	}

	err := t.services.UpdateNumber(number)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}