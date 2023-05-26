package utils

import (
	"encoding/json"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertGrpcToJsonError(c codes.Code, e any) error {
	if e == nil {
		return nil
	}

	data, err := json.Marshal(e)
	if err != nil {
		return status.Error(codes.Internal, fmt.Sprintf("failed to convert error with grpc code to json: %v", err))
	}
	return status.Error(c, string(data))
}
