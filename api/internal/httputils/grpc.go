package httputils

import (
	"errors"
	"google.golang.org/grpc/codes"
)

var mapGrpcError = map[codes.Code]*backendError{
	codes.NotFound:         &ErrMessageNotFound,
	codes.Internal:         &ErrMessageInternal,
	codes.Unavailable:      &ErrMessageUnavailable,
	codes.DeadlineExceeded: &ErrMessageTimeout,
}

func CodeToError(code codes.Code) error {
	err, ok := mapGrpcError[code]
	if !ok {
		return errors.New("code not found")
	}
	return err
}
