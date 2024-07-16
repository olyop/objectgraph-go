package enums

import (
	"errors"
	"strings"
)

type ContactType int

var contactTypes = [...]string{
	"EmailAddress",
	"MobileNumber",
}

func NewContactType(t string) ContactType {
	var c ContactType
	c.Deserialize(t)
	return c
}

func (c ContactType) String() string {
	return contactTypes[c]
}

func (c *ContactType) Deserialize(str string) {
	str = formatContactType(str)

	var found bool
	for i, v := range contactTypes {
		if str != v {
			continue
		}

		*c = ContactType(i)
		found = true
		break
	}

	if !found {
		*c = -1
	}
}

func formatContactType(s string) string {
	words := strings.Split(s, "_")

	for i, word := range words {
		firstLetter := strings.ToUpper(string(word[0]))
		rest := word[1:]

		words[i] = firstLetter + rest
	}

	return strings.Join(words, "")
}

func (ContactType) ImplementsGraphQLType(name string) bool {
	return name == "ContactType"
}

func (c *ContactType) UnmarshalGraphQL(input any) error {
	var err error
	switch input := input.(type) {
	case string:
		c.Deserialize(input)
	default:
		err = errors.New("invalid enum input for ContactType")
	}
	return err
}
