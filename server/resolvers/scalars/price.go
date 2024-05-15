package scalars

import "strconv"

type Price struct {
	Value int
}

func (Price) ImplementsGraphQLType(name string) bool {
	return name == "Price"
}

func (p *Price) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		// convert from string representation of a int
		intValue, err := strconv.Atoi(input)
		if err != nil {
			return err
		}

		p.Value = intValue

		return nil
	default:
		return nil
	}
}

func (p Price) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(p.Value)), nil
}
