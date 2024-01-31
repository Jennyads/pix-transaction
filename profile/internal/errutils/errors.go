package errutils

import "errors"

var (
	ErrInsufficientBalance    = errors.New("insufficient balance")
	ErrKeyNotFound            = errors.New("key not found in the database")
	ErrInvalidKey             = errors.New("invalid key")
	ErrInactiveAccount        = errors.New("inactive account")
	ErrReceiverAccountBlocked = errors.New("account blocked")
	EmailAlreadyExists        = errors.New("email already exists")
	CpfAlreadyExists          = errors.New("cpf already exists")
	PhoneAlreadyExists        = errors.New("phone already exists")
)
