package pix

import (
	"context"
	"encoding/json"
	"log"
	"transaction/internal/event"
	"transaction/internal/profile"
	"transaction/internal/transactions"
)

type Service interface {
	Handler(ctx context.Context, payload []byte) ([]byte, error)
}

type service struct {
	transaction transactions.Service
	profile     profile.Service
	events      event.Client
}

func (s service) Handler(ctx context.Context, payload []byte) ([]byte, error) {

	var pixEvent PixEvent
	err := json.Unmarshal(payload, &pixEvent)
	if err != nil {
		log.Println("error unmarshalling pix event")
		return nil, err
	}

	err = s.Validations(ctx, pixEvent.PixData)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s service) Validations(ctx context.Context, pix *Pix) error {
	balance, err := s.profile.FindAccountBalance(ctx, pix.AccountId, pix.UserID)
	if err != nil {
		return err
	}

	if balance < pix.Amount {
		return ErrInsufficientBalance
	}

	// TODO ...

	return nil
}

//func SendPixWorkflow(ctx workflow.Context, pixEvent *PixEvent) error {
//	//	return nil
//	//}
