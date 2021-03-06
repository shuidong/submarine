// This file was generated by typhen-api

package submarine

import (
	"errors"
	"github.com/shiwano/submarine/server/battle/lib/typhenapi"
)

var _ = errors.New

// GetRoomsObject is a kind of TyphenAPI type.
type GetRoomsObject struct {
	Rooms []*Room `json:"rooms" msgpack:"rooms"`
}

// Coerce the fields.
func (t *GetRoomsObject) Coerce() error {
	if t.Rooms == nil {
		return errors.New("Rooms should not be empty")
	}
	return nil
}

// Bytes creates the byte array.
func (t *GetRoomsObject) Bytes(serializer typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
