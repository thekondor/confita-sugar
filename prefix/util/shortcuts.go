// NOTE: No shortcuts for Consul, Etcd since they are already shipped with builtin prefix support.
package util

import (
	confita_backend "github.com/heetch/confita/backend"
	confita_backend_env "github.com/heetch/confita/backend/env"
	confita_sugar_prefix "github.com/thekondor/confita-sugar/prefix"
)

// Returns Env backend which extends key provided for lookup with a `prefix`
// This is a shortcut for confita/backend/env + confita-sugar/prefix
func NewEnvBackend(prefix string) confita_backend.Backend {
	return confita_sugar_prefix.WithPrefix(prefix, confita_backend_env.NewBackend())
}
