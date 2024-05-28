package scalars

import (
	"database/sql"
	"strconv"
)

type Price struct {
	Value int64
}

func NewPrice(value int64) Price {
	return Price{Value: value}
}

func NewNillPrice(value sql.NullInt64) *Price {
	if !value.Valid {
		return nil
	}

	return &Price{Value: value.Int64}
}

func (Price) ImplementsGraphQLType(name string) bool {
	return name == "Price"
}

func (p *Price) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		// convert from string representation of a int
		value, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return err
		}

		p.Value = value

		return nil
	default:
		return nil
	}
}

func (p Price) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(p.Value, 10)), nil
}
