package listmonk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type importServiceTestSuite struct {
	suite.Suite
}

func TestImportService(t *testing.T) {
	suite.Run(t, new(importServiceTestSuite))
}

func (s *importServiceTestSuite) TestGetImportStatusService() {
	mockResponse := []byte(`{
		"data": {
			"name": "",
			"total": 0,
			"imported": 0,
			"status": "none"
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &GetImportStatusService{
		c: mockClient,
	}

	status, err := service.Do(context.Background())

	s.Nil(err)
	s.Equal("", status.Name)
	s.Equal("none", status.Status)
	s.Equal(uint(0), status.Total)
	s.Equal(uint(0), status.Imported)
}

func (s *importServiceTestSuite) TestGetImportLogsService() {
	mockResponse := []byte(`{
		"data": "test"
	}`)

	mockClient := newMockClient(mockResponse)
	service := &GetImportLogsService{
		c: mockClient,
	}

	logs, err := service.Do(context.Background())

	s.Nil(err)
	s.Equal("test", *logs)
}

func (s *importServiceTestSuite) TestImportSubscribersService() {
	mockResponse := []byte(`{}`)

	mockClient := newMockClient(mockResponse)
	service := &ImportSubscribersService{
		c: mockClient,
	}

	err := service.Do(context.Background())

	s.Nil(err)
}

func (s *importServiceTestSuite) TestDeleteImportService() {
	mockResponse := []byte(`{
		"data": {
			"name": "",
			"total": 0,
			"imported": 0,
			"status": "none"
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &DeleteImportService{
		c: mockClient,
	}

	err := service.Do(context.Background())

	s.Nil(err)
}
