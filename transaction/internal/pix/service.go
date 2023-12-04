package pix

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"transaction/internal/event"
	"transaction/internal/profile"
	"transaction/internal/transactions"
	proto "transaction/proto/v1"
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

	_, err = s.profile.FindKey(ctx, pix.Key, pix.Receiver)
	if err != nil {
		if errors.Is(ErrKeyNotFound, err) {
			return ErrInvalidKey
		}
		return err
	}

	isActive, err := s.profile.IsAccountActive(ctx, &proto.AccountRequest{AccountId: pix.AccountId, UserId: pix.UserID})
	if err != nil {
		return err
	}

	if !isActive.GetValue() {
		return ErrInactiveAccount
	}

	return nil
}

//func SendPixWorkflow(ctx workflow.Context, pixEvent *PixEvent) error {
//	//	return nil
//	//}
