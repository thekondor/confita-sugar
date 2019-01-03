package maybe

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	confita "github.com/heetch/confita"
	confita_backend "github.com/heetch/confita/backend"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestBackend(t *testing.T) {
	suite.Run(t, new(BackendTest))
}

type BackendTest struct {
	suite.Suite
	ctx         context.Context
	backendMock *BackendMock
}

func (test *BackendTest) SetupTest() {
	test.ctx = context.Background()
	test.backendMock = &BackendMock{}
}

func (test *BackendTest) unmarshal(sutBackend confita_backend.Backend, to interface{}) error {
	return (interface{})(sutBackend).(confita.Unmarshaler).Unmarshal(test.ctx, to)
}

func (test *BackendTest) Test_Unmarshal_SuppressesError() {
	testError := errors.New("marshalling error")
	backendMock := test.backendMock.asUnmarshaler()
	dummyTo := struct{}{}

	test.T().Run("top level error", func(t *testing.T) {
		backendMock.On("Unmarshal", test.ctx, &dummyTo).
			Return(testError).
			Once()

		sut := Backend(backendMock).
			WithSuppressedErrors(testError)
		err := test.unmarshal(sut, &dummyTo)
		test.Assert().NoError(err)
	})

	test.T().Run("underlying error", func(t *testing.T) {
		backendMock.On("Unmarshal", test.ctx, &dummyTo).
			Return(errors.Wrap(testError, "wrapped error")).
			Once()

		sut := Backend(backendMock).
			WithSuppressedUnderlyingErrors(testError)
		err := test.unmarshal(sut, &dummyTo)
		test.Assert().NoError(err)
	})

	test.T().Run("error on condition", func(t *testing.T) {
		backendMock.On("Unmarshal", test.ctx, &dummyTo).
			Return(testError).
			Once()

		sut := Backend(backendMock).
			WithSuppression(
				func(ctx context.Context, key string, err error) bool {
					return testError == err
				})
		err := test.unmarshal(sut, &dummyTo)
		test.Assert().NoError(err)
	})
}

func (test *BackendTest) Test_Suppresses_TopLevelError() {
	testError := errors.New("top level error")

	sut := Backend(test.backendMock).
		WithSuppressedErrors(testError)

	test.backendMock.On("Get", test.ctx, "key name").
		Return([]byte{}, testError).
		Once()

	_, err := sut.Get(test.ctx, "key name")
	test.Require().Error(err)
	test.Assert().Equal(confita_backend.ErrNotFound, err)
}

func (test *BackendTest) Test_Suppresses_UnderlyingError() {
	testError := errors.New("underlying error")

	sut := Backend(test.backendMock).
		WithSuppressedUnderlyingErrors(testError)

	test.backendMock.On("Get", test.ctx, "key name").
		Return([]byte{}, errors.Wrap(testError, "wrapped error")).
		Once()

	_, err := sut.Get(test.ctx, "key name")
	test.Require().Error(err)
	test.Assert().Equal(confita_backend.ErrNotFound, err)
}

func (test *BackendTest) Test_Suppresses_OnComplexError() {
	testError := errors.New("test error")

	isNonCriticalErrorTestFunc := func(ctx context.Context, key string, err error) bool {
		return "key name" == key && testError == err
	}
	sut := Backend(test.backendMock).
		WithSuppression(isNonCriticalErrorTestFunc)

	test.T().Run("unknown regular error", func(t *testing.T) {
		test.backendMock.On("Get", test.ctx, "key name").
			Return([]byte{}, errors.New("Unknown error")).
			Once()
		_, err := sut.Get(test.ctx, "key name")

		require.Error(t, err)
		assert.Equal(t, "Unknown error", errors.Cause(err).Error())
	})

	test.T().Run("expected error on condition", func(t *testing.T) {
		test.backendMock.On("Get", test.ctx, "key name").
			Return([]byte{}, testError).
			Once()

		_, err := sut.Get(test.ctx, "key name")
		require.Error(t, err)
		assert.Equal(t, confita_backend.ErrNotFound, err)
	})
}
