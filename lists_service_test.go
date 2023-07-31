package listmonk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type listServiceTestSuite struct {
	suite.Suite
}

func TestListService(t *testing.T) {
	suite.Run(t, new(listServiceTestSuite))
}

func (s *listServiceTestSuite) TestGetListsService() {
	mockResponse := []byte(`{
		"data": {
			"results": [
				{
					"id": 1,
					"created_at": "2020-02-10T23:07:16.194843+01:00",
					"updated_at": "2020-03-06T22:32:01.118327+01:00",
					"uuid": "ce13e971-c2ed-4069-bd0c-240e9a9f56f9",
					"name": "Default list",
					"type": "public",
					"optin": "double",
					"tags": [
						"test"
					],
					"subscriber_count": 2
				},
				{
					"id": 2,
					"created_at": "2020-03-04T21:12:09.555013+01:00",
					"updated_at": "2020-03-06T22:34:46.405031+01:00",
					"uuid": "f20a2308-dfb5-4420-a56d-ecf0618a102d",
					"name": "get",
					"type": "private",
					"optin": "single",
					"tags": [],
					"subscriber_count": 0
				}
			],
			"total": 5,
			"per_page": 20,
			"page": 1
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &GetListsService{
		c: mockClient,
	}

	lists, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(lists))
	assert.Equal(s.T(), "public", lists[0].Type)
	assert.Equal(s.T(), "double", lists[0].Optin)
	assert.Equal(s.T(), "test", lists[0].Tags[0])
	assert.Equal(s.T(), "Default list", lists[0].Name)
}

func (s *listServiceTestSuite) TestGetListService() {
	mockResponse := []byte(`{
		"data": {
			"id": 1,
			"created_at": "2020-02-10T23:07:16.194843+01:00",
			"updated_at": "2020-03-06T22:32:01.118327+01:00",
			"uuid": "ce13e971-c2ed-4069-bd0c-240e9a9f56f9",
			"name": "Default list",
			"type": "public",
			"optin": "double",
			"tags": [
				"test"
			],
			"subscriber_count": 2
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &GetListService{
		c: mockClient,
	}

	list, err := service.Id(1).Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "public", list.Type)
	assert.Equal(s.T(), "double", list.Optin)
	assert.Equal(s.T(), "test", list.Tags[0])
	assert.Equal(s.T(), "Default list", list.Name)
}

func (s *listServiceTestSuite) TestCreateListService() {
	mockResponse := []byte(`{
		"data": {
			"id": 1,
			"created_at": "2020-02-10T23:07:16.194843+01:00",
			"updated_at": "2020-03-06T22:32:01.118327+01:00",
			"uuid": "ce13e971-c2ed-4069-bd0c-240e9a9f56f9",
			"name": "Default list",
			"type": "public",
			"optin": "double",
			"tags": [
				"test"
			],
			"subscriber_count": 2
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &CreateListService{
		c:    mockClient,
		name: "Default list",
	}

	list, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "public", list.Type)
	assert.Equal(s.T(), "double", list.Optin)
	assert.Equal(s.T(), "test", list.Tags[0])
	assert.Equal(s.T(), "Default list", list.Name)
}

func (s *listServiceTestSuite) TestUpdateListService() {
	mockResponse := []byte(`{
		"data": {
			"id": 1,
			"created_at": "2020-02-10T23:07:16.194843+01:00",
			"updated_at": "2020-03-06T22:32:01.118327+01:00",
			"uuid": "ce13e971-c2ed-4069-bd0c-240e9a9f56f9",
			"name": "Default list",
			"type": "public",
			"optin": "double",
			"tags": [
				"test"
			],
			"subscriber_count": 2
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &UpdateListService{
		c:     mockClient,
		id:    1,
		type_: "public",
		optin: "single",
		tags:  []string{"test"},
	}

	list, err := service.Id(1).Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "public", list.Type)
	assert.Equal(s.T(), "double", list.Optin)
	assert.Equal(s.T(), "test", list.Tags[0])
	assert.Equal(s.T(), "Default list", list.Name)
}

func (s *listServiceTestSuite) TestDeleteListService() {
	mockResponse := []byte(`{}`)

	mockClient := newMockClient(mockResponse)
	service := &DeleteListService{
		c:  mockClient,
		id: 1,
	}

	err := service.Do(context.Background())

	assert.Nil(s.T(), err)
}
