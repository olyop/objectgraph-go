package scalars

import (
	"fmt"

	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

func (UUID) ImplementsGraphQLType(name string) bool {
	return name == "UUID"
}

func (t *UUID) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		u, err := uuid.Parse(input)
		if err != nil {
			return err
		}
		t.UUID = u
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (t UUID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.UUID.String())), nil
}
