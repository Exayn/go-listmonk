package listmonk

import (
	"context"
	"fmt"
	"testing"

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

	s.Nil(err)
	s.Greater(len(result), 0)
	s.Equal(uint(1), result[0].Id)
	s.Equal("john@example.com", result[0].Email)
	s.Equal(uint(2), result[1].Id)
	s.Equal("sugar@example.com", result[1].Email)

	s.Equal(1, len(result[0].Lists))
	s.Equal(uint(1), result[0].Lists[0].Id)
	s.Equal(1, len(result[0].Lists[0].Tags))
	s.Equal("test", result[0].Lists[0].Tags[0])

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "GET")
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

	s.Nil(err)
	s.Equal(uint(1), result.Id)
	s.Equal("john@example.com", result.Email)

	s.Equal(uint(1), result.Lists[0].Id)
	s.Equal(1, len(result.Lists[0].Tags))
	s.Equal("test", result.Lists[0].Tags[0])

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "GET")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/subscribers/%d", 1))
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
            "status": "enabled"
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

	s.Nil(err)
	s.Equal(uint(3), result.Id)
	s.Equal("The Subscriber", result.Name)
	s.Equal("subsriber@domain.com", result.Email)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "POST")
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

	s.Nil(err)
	s.Equal(true, *result)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "PUT")
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
            "status": "enabled"
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

	s.Nil(err)
	s.Equal(uint(3), result.Id)
	s.Equal("The Subscriber", result.Name)
	s.Equal("subsriber@domain.com", result.Email)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "PUT")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/subscribers/%d", 3))
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

	s.Nil(err)
	s.Equal(true, *result)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "PUT")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/subscribers/%d/blocklists", 1))
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

	s.Nil(err)
	s.Equal(true, *result)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "PUT")
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

	s.Nil(err)
	s.Equal(true, *result)
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

	s.Nil(err)
	s.Equal(true, *result)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "DELETE")
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

	s.Nil(err)
	s.Equal(true, *result)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "DELETE")
}
