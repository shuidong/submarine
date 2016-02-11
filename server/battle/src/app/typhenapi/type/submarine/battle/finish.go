// This file was generated by typhen-api

package battle

import (
	"app/typhenapi/core"
	"errors"
)

var _ = errors.New

// Finish is a kind of TyphenAPI type.
type Finish struct {
	WinnerUserId int64 `codec:"winner_user_id"`
	FinishedAt   int64 `codec:"finished_at"`
}

// Coerce the fields.
func (t *Finish) Coerce() error {
	return nil
}

// Bytes creates the byte array.
func (t *Finish) Bytes(serializer typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
