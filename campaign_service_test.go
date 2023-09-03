package listmonk

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type campaignServiceTestSuite struct {
	suite.Suite
}

func TestCampaignService(t *testing.T) {
	suite.Run(t, new(campaignServiceTestSuite))
}

func (s *campaignServiceTestSuite) TestGetCampaigns() {
	mockResponse := []byte(`{
        "data": {
			"results": [
				{
					"id": 1,
					"created_at": "2020-03-14T17:36:41.29451+01:00",
					"updated_at": "2020-03-14T17:36:41.29451+01:00",
					"CampaignID": 0,
					"views": 0,
					"clicks": 0,
					"lists": [
						{
							"id": 1,
							"name": "Default list"
						}
					],
					"started_at": null,
					"to_send": 0,
					"sent": 0,
					"uuid": "57702beb-6fae-4355-a324-c2fd5b59a549",
					"type": "regular",
					"name": "Test campaign",
					"subject": "Welcome to listmonk",
					"from_email": "No Reply <noreply@yoursite.com>",
					"body": "<h3>Hi {{ .Subscriber.FirstName }}!</h3>\n\t\t\tThis is a test e-mail campaign. Your second name is {{ .Subscriber.LastName }} and you are from {{ .Subscriber.Attribs.city }}.",
					"send_at": "2020-03-15T17:36:41.293233+01:00",
					"status": "draft",
					"content_type": "richtext",
					"tags": [
						"test-campaign"
					],
					"template_id": 1,
					"messenger": "email"
				}
			],
			"query": "",
			"total": 1,
			"per_page": 20,
			"page": 1
		}
    }`)
	mockClient := newMockClient(mockResponse)
	service := &GetCampaignsService{
		c: mockClient,
	}

	result, err := service.Do(context.Background())

	s.Nil(err)
	s.Equal(len(result), 1)
	s.Equal(uint(1), result[0].Id)
	s.Equal("Test campaign", result[0].Name)
	s.Equal(1, len(result[0].Lists))
	s.Equal(uint(1), result[0].Lists[0].Id)
	s.Equal("Default list", result[0].Lists[0].Name)
}

func (s *campaignServiceTestSuite) TestGetCampaign() {
	mockResponse := []byte(`{
		"data": {
			"id": 1,
			"created_at": "2020-03-14T17:36:41.29451+01:00",
			"updated_at": "2020-03-14T17:36:41.29451+01:00",
			"CampaignID": 0,
			"views": 0,
			"clicks": 0,
			"lists": [
				{
					"id": 1,
					"name": "Default list"
				}
			],
			"started_at": null,
			"to_send": 0,
			"sent": 0,
			"uuid": "57702beb-6fae-4355-a324-c2fd5b59a549",
			"type": "regular",
			"name": "Test campaign",
			"subject": "Welcome to listmonk",
			"from_email": "No Reply <noreply@yoursite.com>",
			"body": "<h3>Hi {{ .Subscriber.FirstName }}!</h3>\n\t\t\tThis is a test e-mail campaign. Your second name is {{ .Subscriber.LastName }} and you are from {{ .Subscriber.Attribs.city }}.",
			"send_at": "2020-03-15T17:36:41.293233+01:00",
			"status": "draft",
			"content_type": "richtext",
			"tags": [
				"test-campaign"
			],
			"template_id": 1,
			"messenger": "email"
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &GetCampaignService{
		c: mockClient,
	}

	result, err := service.Id(1).Do(context.Background())

	s.Nil(err)
	s.Equal(uint(1), result.Id)
	s.Equal("Test campaign", result.Name)
	s.Equal(1, len(result.Lists))
	s.Equal(uint(1), result.Lists[0].Id)
	s.Equal("Default list", result.Lists[0].Name)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "GET")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/campaigns/%d", 1))
}

func (s *campaignServiceTestSuite) TestCreateCampaign() {
	mockResponse := []byte(`{
		"data": {
			"id": 1,
			"created_at": "2020-03-14T17:36:41.29451+01:00",
			"updated_at": "2020-03-14T17:36:41.29451+01:00",
			"CampaignID": 0,
			"views": 0,
			"clicks": 0,
			"lists": [
				{
					"id": 1,
					"name": "Default list"
				}
			],
			"started_at": null,
			"to_send": 0,
			"sent": 0,
			"uuid": "57702beb-6fae-4355-a324-c2fd5b59a549",
			"type": "regular",
			"name": "Test campaign",
			"subject": "Welcome to listmonk",
			"from_email": "No Reply <noreply@yoursite.com>",
			"body": "<h3>Hi {{ .Subscriber.FirstName }}!</h3>\n\t\t\tThis is a test e-mail campaign. Your second name is {{ .Subscriber.LastName }} and you are from {{ .Subscriber.Attribs.city }}.",
			"send_at": "2020-03-15T17:36:41.293233+01:00",
			"status": "draft",
			"content_type": "richtext",
			"tags": [
				"test-campaign"
			],
			"template_id": 1,
			"messenger": "email"
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &CreateCampaignService{
		c: mockClient,
	}

	result, err := service.
		Name("Test campaign").
		Body("<h3>Hi {{ .Subscriber.FirstName }}!</h3>\n\t\t\tThis is a test e-mail campaign. Your second name is {{ .Subscriber.LastName }} and you are from {{ .Subscriber.Attribs.city }}.").
		Subject("Welcome to listmonk").
		FromEmail("No Reply <noreply@yoursite.com>").
		ContentType("html").
		Lists([]uint{1}).
		Do(context.Background())

	s.Nil(err)
	s.Equal(uint(1), result.Id)
	s.Equal("Test campaign", result.Name)
	s.Equal(1, len(result.Lists))
	s.Equal(uint(1), result.Lists[0].Id)
	s.Equal("Default list", result.Lists[0].Name)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "POST")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, "/campaigns")
}

func (s *campaignServiceTestSuite) TestUpdateCampaignStatus() {
	mockResponse := []byte(`{
		"data": {
			"id": 1,
			"created_at": "2020-03-14T17:36:41.29451+01:00",
			"updated_at": "2020-03-14T17:36:41.29451+01:00",
			"CampaignID": 0,
			"views": 0,
			"clicks": 0,
			"lists": [
				{
					"id": 1,
					"name": "Default list"
				}
			],
			"started_at": null,
			"to_send": 0,
			"sent": 0,
			"uuid": "57702beb-6fae-4355-a324-c2fd5b59a549",
			"type": "regular",
			"name": "Test campaign",
			"subject": "Welcome to listmonk",
			"from_email": "No Reply <noreply@yoursite.com>",
			"body": "<h3>Hi {{ .Subscriber.FirstName }}!</h3>\n\t\t\tThis is a test e-mail campaign. Your second name is {{ .Subscriber.LastName }} and you are from {{ .Subscriber.Attribs.city }}.",
			"send_at": "2020-03-15T17:36:41.293233+01:00",
			"status": "draft",
			"content_type": "richtext",
			"tags": [
				"test-campaign"
			],
			"template_id": 1,
			"messenger": "email"
		}
	}`)

	mockClient := newMockClient(mockResponse)
	service := &UpdateCampaignStatusService{
		c: mockClient,
	}

	result, err := service.Id(1).Status("draft").Do(context.Background())

	s.Nil(err)
	s.Equal(uint(1), result.Id)
	s.Equal("Test campaign", result.Name)
	s.Equal(1, len(result.Lists))
	s.Equal(uint(1), result.Lists[0].Id)
	s.Equal("Default list", result.Lists[0].Name)

	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "PUT")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/campaigns/%d/status", 1))
}

func (s *campaignServiceTestSuite) TestDeleteCampaign() {
	mockResponse := []byte(`{}`)

	mockClient := newMockClient(mockResponse)
	service := &DeleteCampaignService{
		c:  mockClient,
		id: 1,
	}

	err := service.Do(context.Background())

	s.Nil(err)
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).method, "DELETE")
	s.Equal(mockClient.Calls[0].Arguments.Get(1).(*request).endpoint, fmt.Sprintf("/campaigns/%d", 1))
}
