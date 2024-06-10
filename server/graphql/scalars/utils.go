package scalars

func NullInt(input *int) *int32 {
	if input == nil {
		return nil
	}

	value := int32(*input)

	return &value
}
