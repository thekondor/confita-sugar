package maybe

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	confita "github.com/heetch/confita"
)

// Optional's marshaling backend details
type optionalUnmarshalingBackend struct {
	optionalBackend
}

// Implementation of `confita.Unmarshaler` interface
func (ub optionalUnmarshalingBackend) Unmarshal(context context.Context, to interface{}) error {
	decoratedUnmarshaler := ub.decoratee.(confita.Unmarshaler)
	err := decoratedUnmarshaler.Unmarshal(context, to)
	if nil != err && ub.isSuppressed(context, "", err) {
		return nil
	}

	return err
}
