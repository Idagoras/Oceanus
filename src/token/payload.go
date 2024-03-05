package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Expiration time.Time `json:"expiration"`
	IssuedTime time.Time `json:"issued_time"`
	jwt.RegisteredClaims
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := Payload{
		tokenID,
		username,
		time.Now().Add(duration),
		time.Now(),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return &payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.Expiration) {
		return ErrExpiredToken
	}
	return nil
}
