package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"time"
)

type PasetoMaker struct {
	paseto        *paseto.V2
	symmetricKey  []byte
	asymmetricKey []byte
	footer        string
}

func NewPasetoMaker(symmetricKey string, asymmetricKey string, footer string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:        paseto.NewV2(),
		symmetricKey:  []byte(symmetricKey),
		asymmetricKey: []byte(asymmetricKey),
		footer:        footer,
	}
	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, maker.footer)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	var jsonToken Payload
	var newFooter string
	err := maker.paseto.Decrypt(token, maker.symmetricKey, &jsonToken, &newFooter)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = jsonToken.Valid()
	if err != nil {
		return nil, ErrExpiredToken
	}
	return &jsonToken, nil
}
