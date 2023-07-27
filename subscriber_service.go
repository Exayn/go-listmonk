package listmonk

import (
	"encoding/json"
	"net/http"
	"time"

	"context"
)

type Subscriber struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Uuid       string    `json:"uuid"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	Attributes string    `json:"attributes"`
}

type GetSubscriberListService struct {
	c       *Client
	page    *uint
	perPage *uint
	query   *string
}

func (s *GetSubscriberListService) Page(page uint) *GetSubscriberListService {
	s.page = &page
	return s
}

func (s *GetSubscriberListService) PerPage(perPage uint) *GetSubscriberListService {
	s.perPage = &perPage
	return s
}

func (s *GetSubscriberListService) Query(query string) *GetSubscriberListService {
	s.query = &query
	return s
}

func (s *GetSubscriberListService) Do(ctx context.Context, opts ...RequestOption) ([]*Subscriber, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/subscribers",
	}

	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.perPage != nil {
		r.setParam("per_page", *s.perPage)
	}
	if s.query != nil {
		r.setParam("query", *s.query)
	}

	data, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	res := make([]*Subscriber, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
