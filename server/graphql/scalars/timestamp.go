package scalars

import (
	"fmt"
	"strconv"
	"time"
)

type Timestamp struct {
	Value time.Time
}

func NewTimestamp(t time.Time) Timestamp {
	return Timestamp{Value: t}
}

func NewNilTimestamp(t time.Time) *Timestamp {
	if t.IsZero() {
		return nil
	}

	return &Timestamp{Value: t}
}

func (Timestamp) ImplementsGraphQLType(name string) bool {
	return name == "Timestamp"
}

func (t *Timestamp) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		intValue, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return err
		}

		t.Value = time.UnixMilli(intValue)

		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	content := string(data)
	str := content[1 : len(content)-1]
	println(str)

	intValue, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	t.Value = time.UnixMilli(intValue)

	return nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%d\"", t.Value.UnixMilli())), nil
}
