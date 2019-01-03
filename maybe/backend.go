package maybe

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	confita_backend "github.com/heetch/confita/backend"
)

// Returns a builder of the optional backend which should be configured with an error suppress routine.
// Builder is created over existing valid Confita's Backend which is supposed to be decorated to suppress errors.
func Backend(backend confita_backend.Backend) optionalBackendBuilder {
	return optionalBackendBuilder{backend}
}

// Alias for a custom error suppressing routine.
//
// `ctx` is an original context passed by the Confita's loader.
//
// `key` is a name of key failed to load; is always empty when is called on a failure of `Unmarshal()`
//
// `err` is a pristine error with a reason to stop further execution of Confita's loader
type IsSuppressedFunc func(ctx context.Context, key string, err error) bool

// Optional's backend details
type optionalBackend struct {
	decoratee        confita_backend.Backend
	isSuppressedFunc IsSuppressedFunc
}

func (ob optionalBackend) isSuppressed(context context.Context, key string, err error) bool {
	return ob.isSuppressedFunc(context, key, err)
}

// implementation of `confita.Backend` interface
func (ob optionalBackend) Get(context context.Context, key string) ([]byte, error) {
	res, err := ob.decoratee.Get(context, key)
	if nil != err && ob.isSuppressed(context, key, err) {
		return nil, confita_backend.ErrNotFound
	}

	return res, err
}

// implementation of `confita.Backend` interface
func (ob optionalBackend) Name() string {
	return "optional:" + ob.decoratee.Name()
}
