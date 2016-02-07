// This file was generated by typhen-api

package battle

import (
	"app/typhenapi/core"
	"errors"
)

var _ = errors.New

// TorpedoRequestObject is a kind of TyphenAPI type.
type TorpedoRequestObject struct {
}

// Coerce the fields.
func (t *TorpedoRequestObject) Coerce() error {
	return nil
}

// Bytes creates the byte array.
func (t *TorpedoRequestObject) Bytes(serializer typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
