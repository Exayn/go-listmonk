package listmonk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "", status.Name)
	assert.Equal(s.T(), "none", status.Status)
	assert.Equal(s.T(), uint(0), status.Total)
	assert.Equal(s.T(), uint(0), status.Imported)
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

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "test", *logs)
}

func (s *importServiceTestSuite) TestImportSubscribersService() {
	mockResponse := []byte(`{}`)

	mockClient := newMockClient(mockResponse)
	service := &ImportSubscribersService{
		c: mockClient,
	}

	err := service.Do(context.Background())

	assert.Nil(s.T(), err)
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

	assert.Nil(s.T(), err)
}
