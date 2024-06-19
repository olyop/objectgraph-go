package trace

import (
	"github.com/olyop/graphqlops-go/graphqlops/graphql/errors"
	"github.com/olyop/graphqlops-go/graphqlops/graphql/trace/tracer"
)

// Deprecated: this type has been deprecated. Use tracer.ValidationFinishFunc instead.
type TraceValidationFinishFunc = func([]*errors.QueryError)

// Deprecated: use ValidationTracerContext.
type ValidationTracer = tracer.LegacyValidationTracer //nolint:staticcheck

// Deprecated: this type has been deprecated. Use tracer.ValidationTracer instead.
type ValidationTracerContext = tracer.ValidationTracer

// Deprecated: use a tracer that implements ValidationTracerContext.
type NoopValidationTracer = tracer.LegacyNoopValidationTracer //nolint:staticcheck
