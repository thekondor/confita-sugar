package maybe

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	confita_backend_file "github.com/heetch/confita/backend/file"
	"github.com/pkg/errors"
	"os"
)

func ExampleBackend_topLevel() {
	// Create a backend which wouldn't fail confita loader if `config.json' could not be opened due to or ErrPermission either ErrNotExist TOP LEVEL (!) errors returned by `os.Open() call
	optionalConfig := /*maybe.*/ Backend(
		confita_backend_file.NewBackend("config.json")).
		WithSuppressedErrors(os.ErrPermission, os.ErrNotExist)
	// NOTE: util/ subpackage should be used instead of this one to handle such cases in a more appropriate way
	_ = optionalConfig
}

func ExampleBackend() {
	// Create a backend which wouldn't fail confita loader if `config.json' could not be opened due to or ErrPermission either ErrNotExist UNDERLYING (!) errors returned by `os.Open() call
	optionalConfig := /*maybe.*/ Backend(
		confita_backend_file.NewBackend("config.json")).
		WithSuppressedUnderlyingErrors(os.ErrInvalid, os.ErrNotExist)
	// NOTE: util/ subpackage should be used instead of this one to handle such cases in a more appropriate way
	_ = optionalConfig
}

func ExampleIsSuppressedFunc() {
	isNonCriticalError := func(ctx context.Context, key string, err error) bool {
		return "fallback_password" == key && os.ErrNotExist == errors.Cause(err)
	}

	// Confita loader wouldn't fail if there is failure during retrieval of a key `fallback_password' due to `os.ErrNotExist' error
	optionalConfig := /*maybe.*/ Backend(
		confita_backend_file.NewBackend("config.json")).
		WithSuppression(isNonCriticalError)
	_ = optionalConfig
}
