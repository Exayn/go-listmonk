package listmonk

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type mediaServiceTestSuite struct {
	suite.Suite
}

func TestMediaService(t *testing.T) {
	suite.Run(t, new(mediaServiceTestSuite))
}

func (s *mediaServiceTestSuite) TestGetMedia() {
	mockResponse := []byte(`{
        "data": [
			{
				"id": 1,
				"uuid": "ec7b45ce-1408-4e5c-924e-965326a20287",
				"filename": "Media file",
				"created_at": "2020-04-08T22:43:45.080058+01:00",
				"thumb_url": "/uploads/image_thumb.jpg",
				"uri": "/uploads/image.jpg"
			}
		]
    }`)

	mockClient := newMockClient(mockResponse)
	service := &GetMediaService{
		c: mockClient,
	}

	result, err := service.Do(context.Background())

	s.Nil(err)
	s.Equal(uint(1), result[0].Id)
	s.Equal("Media file", result[0].Filename)
	s.Equal("ec7b45ce-1408-4e5c-924e-965326a20287", result[0].Uuid)
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "GET")
}

func (s *mediaServiceTestSuite) TestCreateCampaign() {
	mockResponse := []byte(`{
        "data": {
			"id": 1,
			"uuid": "ec7b45ce-1408-4e5c-924e-965326a20287",
			"filename": "Media file",
			"created_at": "2020-04-08T22:43:45.080058+01:00",
			"thumb_url": "/uploads/image_thumb.jpg",
			"uri": "/uploads/image.jpg"
		}
    }`)

	mockClient := newMockClient(mockResponse)
	service := &CreateMediaService{
		c: mockClient,
	}

	result, err := service.File([]byte("file")).Do(context.Background())

	s.Nil(err)
	s.Equal(uint(1), result.Id)
	s.Equal("Media file", result.Filename)
	s.Equal("ec7b45ce-1408-4e5c-924e-965326a20287", result.Uuid)
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "POST")
}

func (s *mediaServiceTestSuite) TestDeleteMedia() {
	mockResponse := []byte(`{}`)

	mockClient := newMockClient(mockResponse)
	service := &DeleteMediaService{
		c:  mockClient,
		id: 1,
	}

	err := service.Do(context.Background())

	s.Nil(err)
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "DELETE")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/media/%d", 1))
}
