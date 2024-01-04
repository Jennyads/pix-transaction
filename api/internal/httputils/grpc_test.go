package httputils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestCodeToError(t *testing.T) {
	cases := []struct {
		name    string
		code    codes.Code
		wantErr error
	}{
		{
			name:    "NotFound",
			code:    codes.NotFound,
			wantErr: &ErrMessageNotFound,
		},
		{
			name:    "Internal",
			code:    codes.Internal,
			wantErr: &ErrMessageInternal,
		},
		{
			name:    "Unavailable",
			code:    codes.Unavailable,
			wantErr: &ErrMessageUnavailable,
		},
		{
			name:    "DeadlineExceeded",
			code:    codes.DeadlineExceeded,
			wantErr: &ErrMessageTimeout,
		},
		{
			name:    "DeadlineExceeded",
			code:    codes.Code(17),
			wantErr: errors.New("code not found"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := CodeToError(tc.code)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
