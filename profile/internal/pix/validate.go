package pix

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidatePixTransaction(pixTransaction *PixTransaction) error {
	validate := validator.New()
	if err := validate.RegisterValidation("validatePixKey", validatePixKey); err != nil {
		return err
	}
	if err := validate.RegisterValidation("validateSenderBalance", validateSenderBalance); err != nil {
		return err
	}

	err := validate.Struct(pixTransaction)
	if err != nil {
		return errors.New("invalid pix transaction " + err.Error())
	}

	return nil
}

func validatePixKey(fl validator.FieldLevel) bool {
	pixKey := fl.Field().String()
	switch {
	case isValidEmail(pixKey):
		return true
	case isValidPhoneNumber(pixKey):
		return true
	case isValidCPF(pixKey):
		return true
	case isValidRandomKey(pixKey):
		return true
	default:
		panic("invalid pix key")
	}
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidPhoneNumber(phone string) bool {
	phoneRegex := regexp.MustCompile(`^\+[0-9]+$`)
	return phoneRegex.MatchString(phone)
}

func isValidCPF(cpf string) bool {
	cpfRegex := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
	return cpfRegex.MatchString(cpf)
}

func isValidRandomKey(random string) bool {
	randomKeyRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return randomKeyRegex.MatchString(random)
}

func validateSenderBalance(fl validator.FieldLevel) bool {
	return true
}
