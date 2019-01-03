package util

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	confita_backend "github.com/heetch/confita/backend"
	confita_backend_file "github.com/heetch/confita/backend/file"
	"github.com/pkg/errors"
	"github.com/thekondor/confita-sugar/maybe"
	"os"
)

func isPathError(err error) bool {
	causeErr := errors.Cause(err)
	if causeErr == err {
		return false
	}

	_, isPathErr := causeErr.(*os.PathError)
	return isPathErr
}

// Returns a file backend which suppresses all IO errors occurred during open of `path` file.
// The backend is built over standard's confita's File backend.
// This is a shortcut of:
//   maybe.Backend(confita_backend_file.NewBackend(...)).
//     WithSuppression(isPathError)
func NewFileBackend(path string) confita_backend.Backend {
	return maybe.Backend(confita_backend_file.NewBackend(path)).
		WithSuppression(func(ctx context.Context, key string, err error) bool {
			return isPathError(err)
		})
}
