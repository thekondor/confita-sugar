package prefix

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	confita_backend "github.com/heetch/confita/backend"
	confita_backend_env "github.com/heetch/confita/backend/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

// Note: for the sake of simplicity, default Environment backend is used in the tests below instead of generic Backend mock (as should be).

func Test_Panics_OnEmptyPrefix(t *testing.T) {
	assert.Panics(t, func() {
		WithPrefix("", confita_backend_env.NewBackend())
	})
}

func Test_Errors_OnNonExistingKey(t *testing.T) {
	sut := WithPrefix("VENDOR_PREFIX", confita_backend_env.NewBackend())

	_, err := sut.Get(context.Background(), "NON_EXISTING_ENV_VAR")
	assert.Equal(t, confita_backend.ErrNotFound, err)
}

func Test_ReturnsValue_OnExistingKey_WithSameCase(t *testing.T) {
	sut := WithPrefix("VENDOR_PREFIX", confita_backend_env.NewBackend())
	os.Setenv("VENDOR_PREFIX_ENV_VAR", "overwritten value")

	value, err := sut.Get(context.Background(), "ENV_VAR")
	require.NoError(t, err)
	assert.Equal(t, []byte("overwritten value"), value)
}

func Test_ReturnsValue_OnExistingKey_WithDifferentCase(t *testing.T) {
	sut := WithPrefix("VENDOR_PREFIX", confita_backend_env.NewBackend())

	os.Setenv("VENDOR_PREFIX_EnV_VaR", "overwritten value")

	value, err := sut.Get(context.Background(), "ENV_VAR")
	require.NoError(t, err)
	assert.Equal(t, []byte("overwritten value"), value)
}
