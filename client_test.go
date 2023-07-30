package listmonk

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type mockClient struct {
	mock.Mock
}

func newMockClient(mockResponse []byte) ClientInterface {
	client := new(mockClient)
	client.On("callAPI", mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, nil)
	return client
}

func (m *mockClient) callAPI(ctx context.Context, r *request, opts ...requestOption) ([]byte, error) {
	args := m.Called(ctx, r, opts)
	return args.Get(0).([]byte), args.Error(1)
}
