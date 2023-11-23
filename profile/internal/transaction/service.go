package transaction

import (
	"context"
	"encoding/json"
	"profile/internal/event"
)

type Service interface {
}

type service struct {
	events event.Client
}

func (s service) SendPix(ctx context.Context, pixEvent PixEvent) error {
	payload, err := json.Marshal(pixEvent)
	if err != nil {
		return err
	}
	return s.events.Publish(ctx, payload)
}
