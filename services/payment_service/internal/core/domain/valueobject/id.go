package valueobject

import (
	"fmt"

	"github.com/google/uuid"
)

type ID struct {
	value uuid.UUID
}

func NewIDFromString(id string) (ID, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return ID{}, fmt.Errorf("invalid UUID: %s", err.Error())
	}
	return ID{value: parsed}, nil
}

func NewID() ID {
	return ID{value: uuid.New()}
}

func (i ID) String() string {
	return i.value.String()
}

func (i ID) UUID() uuid.UUID {
	return i.value
}
