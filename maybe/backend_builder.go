package maybe

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	confita "github.com/heetch/confita"
	confita_backend "github.com/heetch/confita/backend"
	errs "github.com/pkg/errors"
)

type optionalBackendBuilder struct {
	decoratee confita_backend.Backend
}

func makeIsSuppressed(checkCause bool, errors ...error) IsSuppressedFunc {
	errSet := newErrorSet(errors...)

	return func(context context.Context, key string, err error) bool {
		if !checkCause {
			return errSet.contains(err)
		}
		if causeErr := errs.Cause(err); causeErr != err {
			return errSet.contains(causeErr)
		}
		return false
	}
}

func (bb optionalBackendBuilder) WithSuppressedErrors(errors ...error) confita_backend.Backend {
	var checkCauseErr bool = false
	return bb.new(makeIsSuppressed(checkCauseErr, errors...))
}

func (bb optionalBackendBuilder) WithSuppressedUnderlyingErrors(errors ...error) confita_backend.Backend {
	var checkCauseErr bool = true
	return bb.new(makeIsSuppressed(checkCauseErr, errors...))
}

func (bb optionalBackendBuilder) WithSuppression(isSuppressed IsSuppressedFunc) confita_backend.Backend {
	return bb.new(isSuppressed)
}

func (bb optionalBackendBuilder) new(isSuppressedFunc IsSuppressedFunc) confita_backend.Backend {
	return bb.wrap(optionalBackend{
		decoratee:        bb.decoratee,
		isSuppressedFunc: isSuppressedFunc,
	})
}

func (bb optionalBackendBuilder) wrap(backend optionalBackend) confita_backend.Backend {
	switch bb.decoratee.(type) {
	case confita.Unmarshaler:
		return optionalUnmarshalingBackend{backend}
	case confita.StructLoader:
		panic("StructLoader is not implemented")
	}

	return backend
}
