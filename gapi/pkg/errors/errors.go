package errors

import (
	"errors"
)

var (
	ErrInvalidUsers     = errors.New("invalid trade: Buyer and Seller are the same")
	ErrProviderNotFound = errors.New("provider not found")
)
