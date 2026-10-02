package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_At"`
	ExpiresAt time.Time `json:"expires_At"`
}

var (
	ErrExpiredToken error = errors.New("Token has expired!")
	ErrInvalidToken error = errors.New("Token is invalid!")
)

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	newToken, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        newToken,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
