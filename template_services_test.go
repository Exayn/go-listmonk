package listmonk

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type templateServiceTestSuite struct {
	suite.Suite
}

func TestTemplateService(t *testing.T) {
	suite.Run(t, new(templateServiceTestSuite))
}

func (s *templateServiceTestSuite) TestGetTemplatesService() {
	mockResponse := []byte(`{
		"data": [
			{
				"id": 1,
				"created_at": "2020-02-10T23:07:16.199433+01:00",
				"updated_at": "2020-02-10T23:07:16.199433+01:00",
				"name": "Default template",
				"body": "Hello {{.name}}!",
				"type": "html",
				"is_default": true
			},
			{
				"id": 2,
				"created_at": "2020-02-19T19:10:49.36636+01:00",
				"updated_at": "2020-02-19T19:10:49.36636+01:00",
				"name": "Another template",
				"body": "Hello {{.name}}!",
				"type": "html",
				"is_default": false
			}
		]
	}`)

	mockClient := newMockClient(mockResponse)
	service := &GetTemplatesService{
		c: mockClient,
		// set other fields if needed
	}

	res, err := service.Do(context.Background())

	s.Nil(err)
	s.NotNil(res)
	s.Equal(2, len(res))
	s.Equal(uint(1), res[0].Id)
	s.Equal("Default template", res[0].Name)
	s.Equal("Hello {{.name}}!", res[0].Body)
	s.Equal("html", res[0].Type)
	s.Equal(true, res[0].IsDefault)
	s.Equal(uint(2), res[1].Id)
	s.Equal("Another template", res[1].Name)
	s.Equal("Hello {{.name}}!", res[1].Body)
	s.Equal("html", res[1].Type)
	s.Equal(false, res[1].IsDefault)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "GET")
}

func (s *templateServiceTestSuite) TestGetTemplateService() {
	mockResponse := []byte(`{
		"data": {
			"id": 1,
			"created_at": "2020-02-10T23:07:16.199433+01:00",
			"updated_at": "2020-02-10T23:07:16.199433+01:00",
			"name": "Default template",
			"body": "Hello {{.name}}!",
			"type": "html",
			"is_default": true
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &GetTemplateService{
		c: mockClient,
		// set other fields if needed
	}

	service.Id(1)

	res, err := service.Do(context.Background())

	s.Nil(err)
	s.NotNil(res)
	s.Equal(uint(1), res.Id)
	s.Equal("Default template", res.Name)
	s.Equal("Hello {{.name}}!", res.Body)
	s.Equal("html", res.Type)
	s.Equal(true, res.IsDefault)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "GET")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/templates/%d", 1))
}

func (s *templateServiceTestSuite) TestGetTemplatePreviewService() {
	mockResponse := []byte("<p>Hello world!</p>")

	mockClient := newMockClient(mockResponse)
	service := &GetTemplatePreviewService{
		c: mockClient,
		// set other fields if needed
	}

	service.Id(1)

	res, err := service.Do(context.Background())

	s.Nil(err)
	s.NotNil(res)
	s.Equal("<p>Hello world!</p>", res)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "GET")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/templates/%d/preview", 1))
}

func (s *templateServiceTestSuite) TestUpdateTemplateAsDefaultService() {
	mockResponse := []byte(`{
		"data": {
			"id": 1,
			"created_at": "2020-02-10T23:07:16.199433+01:00",
			"updated_at": "2020-02-10T23:07:16.199433+01:00",
			"name": "Default template",
			"body": "Hello {{.name}}!",
			"type": "html",
			"is_default": true
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &UpdateTemplateAsDefaultService{
		c: mockClient,
		// set other fields if needed
	}

	service.Id(1)

	res, err := service.Do(context.Background())

	s.Nil(err)
	s.NotNil(res)
	s.Equal(uint(1), res.Id)
	s.Equal("Default template", res.Name)
	s.Equal("Hello {{.name}}!", res.Body)
	s.Equal("html", res.Type)
	s.Equal(true, res.IsDefault)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "PUT")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/templates/%d/default", 1))
}

func (s *templateServiceTestSuite) TestDeleteTemplateService() {
	mockResponse := []byte(`{
		"data": {
			"id": 1,
			"created_at": "2020-02-10T23:07:16.199433+01:00",
			"updated_at": "2020-02-10T23:07:16.199433+01:00",
			"name": "Default template",
			"body": "Hello {{.name}}!",
			"type": "html",
			"is_default": true
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &DeleteTemplateService{
		c: mockClient,
		// set other fields if needed
	}

	service.Id(1)

	err := service.Do(context.Background())

	s.Nil(err)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "DELETE")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/templates/%d", 1))
}
