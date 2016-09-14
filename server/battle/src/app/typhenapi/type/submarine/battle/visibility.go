// This file was generated by typhen-api

package battle

import (
	"app/typhenapi/core"
	"errors"
)

var _ = errors.New

// Visibility is a kind of TyphenAPI type.
type Visibility struct {
	ActorId   int64     `codec:"actor_id"`
	IsVisible bool      `codec:"is_visible"`
	Movement  *Movement `codec:"movement"`
}

// Coerce the fields.
func (t *Visibility) Coerce() error {
	if t.Movement == nil {
		return errors.New("Movement should not be empty")
	}
	return nil
}

// Bytes creates the byte array.
func (t *Visibility) Bytes(serializer typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
