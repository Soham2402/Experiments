package stew

import "errors"

var (
	ErrEmptyKey   = errors.New("key cannot be empty")
	ErrExpiredTTL = errors.New("expiry time is invalid")
)
