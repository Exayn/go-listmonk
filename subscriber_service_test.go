package listmonk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type subscriberServiceTestSuite struct {
	suite.Suite
}

func TestSubscriberService(t *testing.T) {
	suite.Run(t, new(subscriberServiceTestSuite))
}

func (s *subscriberServiceTestSuite) TestGetSubscriberListService() {
	mockResponse := []byte(`{
        "data": {
            "results": [
                {
                    "id": 1,
                    "created_at": "2020-02-10T23:07:16.199433+01:00",
                    "updated_at": "2020-02-10T23:07:16.199433+01:00",
                    "uuid": "ea06b2e7-4b08-4697-bcfc-2a5c6dde8f1c",
                    "email": "john@example.com",
                    "name": "John Doe",
                    "attribs": {
                        "city": "Bengaluru",
                        "good": true,
                        "type": "known"
                    },
                    "status": "enabled",
                    "lists": [
                        {
                            "subscription_status": "unconfirmed",
                            "id": 1,
                            "uuid": "ce13e971-c2ed-4069-bd0c-240e9a9f56f9",
                            "name": "Default list",
                            "type": "public",
                            "tags": [
                                "test"
                            ],
                            "created_at": "2020-02-10T23:07:16.194843+01:00",
                            "updated_at": "2020-02-10T23:07:16.194843+01:00"
                        }
                    ]
                },
                {
                    "id": 2,
                    "created_at": "2020-02-19T19:10:49.36636+01:00",
                    "updated_at": "2020-02-19T19:10:49.36636+01:00",
                    "uuid": "5d940585-3cc8-4add-b9c5-76efba3c6edd",
                    "email": "sugar@example.com",
                    "name": "sugar",
                    "attribs": {},
                    "status": "enabled",
                    "lists": []
                }
            ],
            "query": "",
            "total": 3,
            "per_page": 20,
            "page": 1
        }
    }`)
	mockClient := newMockClient(mockResponse)
	service := &GetSubscribersService{
		c: mockClient,
		// set other fields if needed
	}

	result, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Greater(s.T(), len(result), 0)
	assert.Equal(s.T(), uint(1), result[0].Id)
	assert.Equal(s.T(), "john@example.com", result[0].Email)
	assert.Equal(s.T(), uint(2), result[1].Id)
	assert.Equal(s.T(), "sugar@example.com", result[1].Email)
}

func (s *subscriberServiceTestSuite) TestGetSubscriberService() {
	mockResponse := []byte(`{
        "data": {
            "id": 1,
            "created_at": "2020-02-10T23:07:16.199433+01:00",
            "updated_at": "2020-02-10T23:07:16.199433+01:00",
            "uuid": "ea06b2e7-4b08-4697-bcfc-2a5c6dde8f1c",
            "email": "john@example.com",
            "name": "John Doe",
            "attribs": {
                "city": "Bengaluru",
                "good": true,
                "type": "known"
            },
            "status": "enabled",
            "lists": [
                {
                    "subscription_status": "unconfirmed",
                    "id": 1,
                    "uuid": "ce13e971-c2ed-4069-bd0c-240e9a9f56f9",
                    "name": "Default list",
                    "type": "public",
                    "tags": [
                        "test"
                    ],
                    "created_at": "2020-02-10T23:07:16.194843+01:00",
                    "updated_at": "2020-02-10T23:07:16.194843+01:00"
                }
            ]
        }
    }`)

	mockClient := newMockClient(mockResponse)

	service := &GetSubscriberService{
		c: mockClient,
	}

	result, err := service.Id(1).Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), uint(1), result.Id)
	assert.Equal(s.T(), "john@example.com", result.Email)
}

func (s *subscriberServiceTestSuite) TestCreateSubscriberService() {
	mockResponse := []byte(`
    {
        "data": {
            "id": 3,
            "created_at": "2019-07-03T12:17:29.735507+05:30",
            "updated_at": "2019-07-03T12:17:29.735507+05:30",
            "uuid": "eb420c55-4cfb-4972-92ba-c93c34ba475d",
            "email": "subsriber@domain.com",
            "name": "The Subscriber",
            "attribs": {},
            "status": "enabled",
            "lists": [1]
        }
    }`)

	mockClient := newMockClient(mockResponse)

	service := &CreateSubscriberService{
		c:       mockClient,
		email:   "subsriber@domain.com",
		name:    "The Subscriber",
		status:  "enabled",
		listIds: []uint{1},
	}

	result, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), uint(3), result.Id)
	assert.Equal(s.T(), "The Subscriber", result.Name)
	assert.Equal(s.T(), "subsriber@domain.com", result.Email)
}

func (s *subscriberServiceTestSuite) TestUpdateSubscribersListsService() {
	mockResponse := []byte(`
    {
        "data": true 
    }`)

	mockClient := newMockClient(mockResponse)

	service := &UpdateSubscribersListsService{
		c:       mockClient,
		listIds: []uint{1},
		action:  "remove",
		ids:     []uint{1, 2},
	}

	result, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), true, *result)
}

func (s *subscriberServiceTestSuite) TestUpdateSubscriberService() {
	mockResponse := []byte(`
    {
        "data": {
            "id": 3,
            "created_at": "2019-07-03T12:17:29.735507+05:30",
            "updated_at": "2019-07-03T12:17:29.735507+05:30",
            "uuid": "eb420c55-4cfb-4972-92ba-c93c34ba475d",
            "email": "subsriber@domain.com",
            "name": "The Subscriber",
            "attribs": {},
            "status": "enabled",
            "lists": [1]
        }
    }`)

	mockClient := newMockClient(mockResponse)

	service := &UpdateSubscriberService{
		c:       mockClient,
		id:      3,
		email:   "subsriber@domain.com",
		name:    "The Subscriber",
		status:  "enabled",
		listIds: []uint{1},
	}

	result, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), uint(3), result.Id)
	assert.Equal(s.T(), "The Subscriber", result.Name)
	assert.Equal(s.T(), "subsriber@domain.com", result.Email)
}

func (s *subscriberServiceTestSuite) TestBlocklistsSubscriberService() {
	mockResponse := []byte(`
    {
        "data": true 
    }`)

	mockClient := newMockClient(mockResponse)

	service := &BlocklistsSubscriberService{
		c:  mockClient,
		id: 1,
	}

	result, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), true, *result)
}

func (s *subscriberServiceTestSuite) TestBlocklistsQuerySubscriberService() {
	mockResponse := []byte(`
    {
        "data": true 
    }`)

	mockClient := newMockClient(mockResponse)

	service := &BlocklistsQuerySubscriberService{
		c:     mockClient,
		query: "test",
	}

	result, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), true, *result)
}

func (s *subscriberServiceTestSuite) TestDeleteSubscriberService() {
	mockResponse := []byte(`
    {
        "data": true 
    }`)

	mockClient := newMockClient(mockResponse)

	service := &DeleteSubscriberService{
		c:  mockClient,
		id: 1,
	}

	result, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), true, *result)
}

func (s *subscriberServiceTestSuite) TestDeleteSubscribersService() {
	mockResponse := []byte(`
    {
        "data": true 
    }`)

	mockClient := newMockClient(mockResponse)

	service := &DeleteSubscribersService{
		c:   mockClient,
		ids: []uint{1, 2},
	}

	result, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), true, *result)
}

func (s *subscriberServiceTestSuite) TestDeleteSubscribersQueryService() {
	mockResponse := []byte(`
    {
        "data": true 
    }`)

	mockClient := newMockClient(mockResponse)

	service := &DeleteSubscribersQueryService{
		c:     mockClient,
		query: "test",
	}

	result, err := service.Do(context.Background())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), true, *result)
}
