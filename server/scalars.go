package main

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var uuidScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:         "UUID",
	Serialize:    uuidScalarSerialize,
	ParseValue:   uuidScalarParseValue,
	ParseLiteral: uuidScalarParseLiteral,
})

func uuidScalarSerialize(value any) any {
	switch value := value.(type) {
	case uuid.UUID:
		return value.String()
	case *uuid.UUID:
		return value.String()
	default:
		return nil
	}
}

func uuidScalarParseValue(value any) any {
	switch value := value.(type) {
	case string:
		val, err := uuid.Parse(value)
		if err != nil {
			return nil
		}

		return val
	default:
		return nil
	}
}

func uuidScalarParseLiteral(valueAST ast.Value) any {
	switch valueAST := valueAST.(type) {
	case *ast.StringValue:
		val, err := uuid.Parse(valueAST.Value)
		if err != nil {
			return nil
		}

		return val
	default:
		return nil
	}
}

var timestampScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:         "Timestamp",
	Serialize:    timestampScalarSerialize,
	ParseValue:   timestampScalarParseValue,
	ParseLiteral: timestampScalarParseLiteral,
})

func timestampScalarSerialize(value any) any {
	switch value := value.(type) {
	case time.Time:
		milli := value.UnixMilli()
		return strconv.FormatInt(milli, 10)
	default:
		return nil
	}
}

func timestampScalarParseValue(value any) any {
	switch value := value.(type) {
	case string:
		milli, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil
		}

		return time.UnixMilli(milli)
	default:
		return nil
	}
}

func timestampScalarParseLiteral(valueAST ast.Value) any {
	switch valueAST := valueAST.(type) {
	case *ast.StringValue:
		milli, err := strconv.ParseInt(valueAST.Value, 10, 64)
		if err != nil {
			return nil
		}

		return time.UnixMilli(milli)
	default:
		return nil
	}
}
