package util

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	"fmt"
	confita "github.com/heetch/confita"
	confita_backend_file "github.com/heetch/confita/backend/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_FileBackend_OriginalDecoratesErrorIsPreserved(t *testing.T) {
	loader := confita.NewLoader(
		NewFileBackend("testdata/unsupported.cfg"),
	)

	dummyTo := struct{}{}
	err := loader.Load(context.Background(), &dummyTo)
	require.Error(t, err)
	// NOTE: not a best idea ever to rely on an implementation detail; currently there is no better way
	require.Contains(t, fmt.Sprintf("%s", err), "unsupported extension")
}

func Test_FileBackend_PreservesLoadAndEvaluationOrder(t *testing.T) {
	type FakeConfig struct {
		First  string `config:"first"`
		Second string `config:"second"`
		Third  string `config:"third"`
	}
	cfg := FakeConfig{Second: "predefined second value"}

	loader := confita.NewLoader(
		confita_backend_file.NewBackend("testdata/main.yaml"),
		NewFileBackend("testdata/non existing optional.yaml"),
		confita_backend_file.NewBackend("testdata/extra.json"),
	)
	err := loader.Load(context.Background(), &cfg)
	require.NoError(t, err)

	assert.Equal(t, "value from first file", cfg.First)
	assert.Equal(t, "predefined second value", cfg.Second)
	assert.Equal(t, "third value from json", cfg.Third)
}

func Test_FileBackend_SuppressesPathError(t *testing.T) {
	nonExistingFn := "intentionally-non-existing.yaml"

	_, existsErr := os.Stat(nonExistingFn)
	require.True(t, os.IsNotExist(existsErr))

	t.Run("original file backend errors on IO", func(t *testing.T) {
		loader := confita.NewLoader(
			confita_backend_file.NewBackend(nonExistingFn),
		)

		err := loader.Load(context.Background(), nil)
		require.Error(t, err)
	})
	t.Run("maybe file suppresses IO error", func(t *testing.T) {
		loader := confita.NewLoader(
			NewFileBackend(nonExistingFn),
		)

		dummyTo := struct{}{}
		err := loader.Load(context.Background(), &dummyTo)
		assert.NoError(t, err)
	})
}
