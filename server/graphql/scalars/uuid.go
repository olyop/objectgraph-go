package scalars

import (
	"fmt"

	"github.com/google/uuid"
)

type UUID struct {
	Value uuid.UUID
}

func NewUUID(uuid uuid.UUID) UUID {
	return UUID{Value: uuid}
}

func (UUID) ImplementsGraphQLType(name string) bool {
	return name == "UUID"
}

func (v *UUID) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		u, err := uuid.Parse(input)
		if err != nil {
			return err
		}
		v.Value = u
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (t *UUID) UnmarshalJSON(data []byte) error {
	content := string(data)
	str := content[1 : len(content)-1]

	u, err := uuid.Parse(str)
	if err != nil {
		return err
	}

	t.Value = u

	return nil
}

func (t UUID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Value.String())), nil
}
