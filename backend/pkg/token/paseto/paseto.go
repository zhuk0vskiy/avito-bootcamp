package paseto

import (
	"fmt"
	"time"

	tokenUtils "backend/pkg/token"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
)

var KEY = "hgtpgf33hgtpgf33hgtpgf33hgtpgf33"

type Paseto struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

type Interface interface {
	CreateToken(UserID uuid.UUID, isModerator bool, duration time.Duration) (string, error)
	VerifyToken(token string) (*tokenUtils.Payload, error)
}

func NewPaseto(symmetricKey string) (*Paseto, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &Paseto{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (p *Paseto) CreateToken(UserID uuid.UUID, isModerator bool, duration time.Duration) (string, error) {
	payload, err := tokenUtils.NewPayload(UserID, isModerator, duration)
	if err != nil {
		return "", err
	}

	return p.paseto.Encrypt(p.symmetricKey, payload, nil)
}

func (maker *Paseto) VerifyToken(token string) (*tokenUtils.Payload, error) {
	payload := &tokenUtils.Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, errors.Errorf("token has expired")
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
