package prefixed_env

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	confita_backend "github.com/heetch/confita/backend"
	temporary_impl_alias "github.com/thekondor/confita-prefixed-env"
)

// See https://github.com/thekondor/confita-prefixed-env#usage for the usage details.
func NewDefaultBackend(prefix string) confita_backend.Backend {
	return temporary_impl_alias.NewDefaultBackend(prefix)
}

// See https://github.com/thekondor/confita-prefixed-env#usage for the usage details.
func NewBackend(prefix string, prefixDelimeter string) confita_backend.Backend {
	return temporary_impl_alias.NewBackend(prefix, prefixDelimeter)
}
