package errutils

import "errors"

var (
	ErrInsufficientBalance    = errors.New("insufficient balance")
	ErrKeyNotFound            = errors.New("key not found in the database")
	ErrInvalidKey             = errors.New("invalid key")
	ErrInactiveAccount        = errors.New("inactive account")
	ErrReceiverAccountBlocked = errors.New("account blocked")
)
