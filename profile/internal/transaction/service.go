package transaction

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"profile/internal/event"
)

type Service interface {
	SendPix(ctx context.Context, req *Pix) error
}

type service struct {
	events event.Client
}

func (s service) SendPix(ctx context.Context, req *Pix) error {
	pixEvent := PixEvent{
		PixData: req,
	}

	payload, err := json.Marshal(pixEvent)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if err := s.events.Publish(ctx, payload); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func NewService(event event.Client) Service {
	return &service{
		events: event,
	}
}
