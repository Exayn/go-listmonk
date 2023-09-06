package listmonk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type List struct {
	Id               uint      `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Uuid             string    `json:"uuid"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	Optin            string    `json:"optin"`
	Tags             []string  `json:"tags"`
	SubscribersCount uint      `json:"subscribers_count"`
}

type GetListsService struct {
	c       ClientInterface
	query   *string
	orderBy *string
	order   *string
	page    *uint
	perPage *string
}

func (s *GetListsService) Query(query string) *GetListsService {
	s.query = &query
	return s
}

func (s *GetListsService) OrderBy(orderBy string) *GetListsService {
	s.orderBy = &orderBy
	return s
}

func (s *GetListsService) Order(order string) *GetListsService {
	s.order = &order
	return s
}

func (s *GetListsService) Page(page uint) *GetListsService {
	s.page = &page
	return s
}

func (s *GetListsService) PerPage(perPage string) *GetListsService {
	s.perPage = &perPage
	return s
}

func (s *GetListsService) Do(ctx context.Context, opts ...requestOption) ([]*List, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/lists",
	}

	if s.query != nil {
		r.setParam("query", *s.query)
	}

	if s.orderBy != nil {
		r.setParam("order_by", *s.orderBy)
	}

	if s.order != nil {
		r.setParam("order", *s.order)
	}

	if s.page != nil {
		r.setParam("page", *s.page)
	}

	if s.perPage != nil {
		r.setParam("per_page", *s.perPage)
	}

	bytes, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var res map[string]map[string]interface{}
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	finalBytes, err := json.Marshal(res["data"]["results"])

	if err != nil {
		return nil, err
	}

	finalRes := make([]*List, 0)
	err = json.Unmarshal(finalBytes, &finalRes)

	if err != nil {
		return nil, err
	}

	return finalRes, nil
}

type GetListService struct {
	c  ClientInterface
	id uint
}

func (s *GetListService) Id(id uint) *GetListService {
	s.id = id
	return s
}

func (s *GetListService) Do(ctx context.Context, opts ...requestOption) (*List, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/lists/%d", s.id),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var result map[string]*List
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result["data"], nil
}

type CreateListService struct {
	c     ClientInterface
	name  string
	type_ string
	optin string
	tags  []string
}

func (s *CreateListService) Name(name string) *CreateListService {
	s.name = name
	return s
}

func (s *CreateListService) Type(type_ string) *CreateListService {
	s.type_ = type_
	return s
}

func (s *CreateListService) Optin(optin string) *CreateListService {
	s.optin = optin
	return s
}

func (s *CreateListService) Tags(tags []string) *CreateListService {
	s.tags = tags
	return s
}

func (s *CreateListService) Do(ctx context.Context, opts ...requestOption) (*List, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/lists",
	}

	r.setFormParam("name", s.name)
	r.setFormParam("type", s.type_)
	r.setFormParam("optin", s.optin)

	if len(s.tags) > 0 {
		r.setFormParamList("tags", s.tags)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var result map[string]*List
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result["data"], nil
}

type UpdateListService struct {
	c     ClientInterface
	id    uint
	type_ string
	optin string
	tags  []string
}

func (s *UpdateListService) Id(id uint) *UpdateListService {
	s.id = id
	return s
}

func (s *UpdateListService) Type(type_ string) *UpdateListService {
	s.type_ = type_
	return s
}

func (s *UpdateListService) Optin(optin string) *UpdateListService {
	s.optin = optin
	return s
}

func (s *UpdateListService) Tags(tags []string) *UpdateListService {
	s.tags = tags
	return s
}

func (s *UpdateListService) Do(ctx context.Context, opts ...requestOption) (*List, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/lists/%d", s.id),
	}

	r.setFormParam("type", s.type_)
	r.setFormParam("optin", s.optin)

	if s.tags != nil && len(s.tags) > 0 {
		r.setFormParam("tags", s.tags)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var result map[string]*List
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result["data"], nil
}

type DeleteListService struct {
	c  ClientInterface
	id uint
}

func (s *DeleteListService) Id(id uint) *DeleteListService {
	s.id = id
	return s
}

func (s *DeleteListService) Do(ctx context.Context, opts ...requestOption) error {
	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/lists/%d", s.id),
	}

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}
