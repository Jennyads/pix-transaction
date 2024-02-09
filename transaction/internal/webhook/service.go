package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Service interface {
	Send(ctx context.Context, data Webhook, url string) error
}

type service struct {
	client http.Client
}

func (s *service) Send(ctx context.Context, data Webhook, url string) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return err
	}

	hreq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return err
	}

	hreq.Header.Set("Content-Type", "application/json")

	result, err := s.client.Do(hreq)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	if result.StatusCode < 200 || result.StatusCode >= 300 {
		return errors.New("error when send webhook")
	}

	return nil
}

func NewService() Service {
	return &service{}
}
