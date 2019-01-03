package maybe

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type BackendMock struct {
	mock.Mock
}

func (mock BackendMock) Get(ctx context.Context, key string) ([]byte, error) {
	args := mock.Called(ctx, key)
	return args.Get(0).([]byte), args.Error(1)
}

func (mock BackendMock) Name() string {
	return mock.Called().String()
}

type UnmarshalingBackendMock struct {
	BackendMock
}

func (mock BackendMock) asUnmarshaler() UnmarshalingBackendMock {
	return UnmarshalingBackendMock{BackendMock: mock}
}

func (mock UnmarshalingBackendMock) Unmarshal(ctx context.Context, to interface{}) error {
	args := mock.Called(ctx, to)
	return args.Error(0)
}
