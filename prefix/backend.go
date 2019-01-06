package prefix

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	confita_backend "github.com/heetch/confita/backend"
	"strings"
)

// Vendored backend details
type vendoredBackend struct {
	decoratee    confita_backend.Backend
	vendorPrefix string
}

const (
	// Delimiter used to build a final key name to look up.
	PrefixDelimiter = "_"
)

// Decorates provided Confita's backend to look every key as `prefix' + `_' + key.
// `Get` query is always forwarded to the decorated original backend. No case operations are performed for a key name additionally.
func WithPrefix(prefix string, backend confita_backend.Backend) confita_backend.Backend {
	if 0 == len(prefix) {
		panic("Empty vendor's prefix")
	}

	return vendoredBackend{decoratee: backend, vendorPrefix: prefix}
}

// implementation of `confita.Backend` interface
func (vb vendoredBackend) Get(context context.Context, key string) ([]byte, error) {
	fullKey := strings.Join([]string{vb.vendorPrefix, key}, PrefixDelimiter)
	return vb.decoratee.Get(context, fullKey)
}

// implementation of `confita.Backend` interface
func (vb vendoredBackend) Name() string {
	return vb.vendorPrefix + ":" + vb.decoratee.Name()
}
