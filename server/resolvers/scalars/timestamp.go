package scalars

import (
	"fmt"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func (Timestamp) ImplementsGraphQLType(name string) bool {
	return name == "Timestamp"
}

func (t *Timestamp) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		// convert from string representation of a int
		intValue, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return err
		}

		t.Time = time.Unix(intValue, 0)

		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", t.Unix())), nil
}
