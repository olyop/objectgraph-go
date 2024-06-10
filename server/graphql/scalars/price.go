package scalars

import (
	"strconv"
)

type Price struct {
	Value int
}

func NewPrice(value int) Price {
	return Price{Value: value}
}

func NewNilPrice(value *int) *Price {
	if value == nil {
		return nil
	}

	return &Price{Value: *value}
}

func (Price) ImplementsGraphQLType(name string) bool {
	return name == "Price"
}

func (p *Price) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		// convert from string representation of a int
		value, err := strconv.Atoi(input)
		if err != nil {
			return err
		}

		p.Value = value

		return nil
	default:
		return nil
	}
}

func (p *Price) UnmarshalJSON(data []byte) error {
	value, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}

	p.Value = value

	return nil
}

func (p Price) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(p.Value)), nil
}
