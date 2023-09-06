package listmonk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"context"
)

type SubscriberList struct {
	Id                 uint      `json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Uuid               string    `json:"uuid"`
	Name               string    `json:"name"`
	Type               string    `json:"type"`
	Tags               []string  `json:"tags"`
	SubscriptionStatus string    `json:"subscription_status"`
}

type Subscriber struct {
	Id         uint             `json:"id"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
	Uuid       string           `json:"uuid"`
	Email      string           `json:"email"`
	Name       string           `json:"name"`
	Status     string           `json:"status"`
	Lists      []SubscriberList `json:"lists"`
	Attributes string           `json:"attributes"`
}

type GetSubscribersService struct {
	c       ClientInterface
	page    *uint
	perPage *string
	query   *string
	listIds []uint
}

func (s *GetSubscribersService) Page(page uint) *GetSubscribersService {
	s.page = &page
	return s
}

func (s *GetSubscribersService) PerPage(perPage string) *GetSubscribersService {
	s.perPage = &perPage
	return s
}

func (s *GetSubscribersService) Query(query string) *GetSubscribersService {
	s.query = &query
	return s
}

func (s *GetSubscribersService) ListIds(listIds []uint) *GetSubscribersService {
	s.listIds = listIds
	return s
}

func (s *GetSubscribersService) Do(ctx context.Context, opts ...requestOption) ([]*Subscriber, error) {
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
	if len(s.listIds) > 0 {
		r.setParamList("list_id", s.listIds)
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

	finalRes := make([]*Subscriber, 0)
	err = json.Unmarshal(finalBytes, &finalRes)

	if err != nil {
		return nil, err
	}

	return finalRes, nil
}

type GetSubscriberService struct {
	c  ClientInterface
	id uint
}

func (s *GetSubscriberService) Id(id uint) *GetSubscriberService {
	s.id = id
	return s
}

func (s *GetSubscriberService) Do(ctx context.Context, opts ...requestOption) (*Subscriber, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/subscribers/%d", s.id),
	}

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string]interface{}
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	finalBytes, err := json.Marshal(res["data"])

	if err != nil {
		return nil, err
	}

	finalRes := &Subscriber{}
	err = json.Unmarshal(finalBytes, finalRes)

	if err != nil {
		return nil, err
	}

	return finalRes, nil
}

type CreateSubscriberService struct {
	c                       ClientInterface
	email                   string
	name                    string
	status                  string
	listIds                 []uint
	attributes              *string
	preconfirmSubscriptions *bool
}

func (s *CreateSubscriberService) Email(email string) *CreateSubscriberService {
	s.email = email
	return s
}

func (s *CreateSubscriberService) Name(name string) *CreateSubscriberService {
	s.name = name
	return s
}

func (s *CreateSubscriberService) Status(status string) *CreateSubscriberService {
	s.status = status
	return s
}

func (s *CreateSubscriberService) ListIds(listIds []uint) *CreateSubscriberService {
	s.listIds = listIds
	return s
}

func (s *CreateSubscriberService) Attributes(attributes string) *CreateSubscriberService {
	s.attributes = &attributes
	return s
}

func (s *CreateSubscriberService) PreconfirmSubscriptions(preconfirmSubscriptions bool) *CreateSubscriberService {
	s.preconfirmSubscriptions = &preconfirmSubscriptions
	return s
}

func (s *CreateSubscriberService) Do(ctx context.Context, opts ...requestOption) (*Subscriber, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/subscribers",
	}

	r.setJsonParam("email", s.email)
	r.setJsonParam("name", s.name)
	r.setJsonParam("status", s.status)

	if len(s.listIds) > 0 {
		r.setJsonParam("lists", s.listIds)
	}
	if s.attributes != nil {
		r.setJsonParam("attributes", s.attributes)
	}
	if s.preconfirmSubscriptions != nil {
		r.setJsonParam("preconfirm_subscriptions", s.preconfirmSubscriptions)
	}

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string]interface{}
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	finalBytes, err := json.Marshal(res["data"])

	if err != nil {
		return nil, err
	}

	finalRes := &Subscriber{}
	err = json.Unmarshal(finalBytes, finalRes)

	if err != nil {
		return nil, err
	}

	return finalRes, nil
}

type UpdateSubscribersListsService struct {
	c       ClientInterface
	ids     []uint
	action  string
	listIds []uint
	status  *string
}

func (s *UpdateSubscribersListsService) Ids(ids []uint) *UpdateSubscribersListsService {
	s.ids = ids
	return s
}

func (s *UpdateSubscribersListsService) Action(action string) *UpdateSubscribersListsService {
	s.action = action
	return s
}

func (s *UpdateSubscribersListsService) ListIds(listIds []uint) *UpdateSubscribersListsService {
	s.listIds = listIds
	return s
}

func (s *UpdateSubscribersListsService) Status(status string) *UpdateSubscribersListsService {
	s.status = &status
	return s
}

func (s *UpdateSubscribersListsService) Do(ctx context.Context, opts ...requestOption) (*bool, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/subscribers/lists",
	}

	r.setJsonParam("ids", s.ids)
	r.setJsonParam("action", s.action)
	r.setJsonParam("target_list_ids", s.listIds)

	if s.status != nil {
		r.setJsonParam("status", s.status)
	}

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string]*bool
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	return res["data"], nil
}

type UpdateSubscriberService struct {
	c                       ClientInterface
	id                      uint
	email                   string
	name                    string
	status                  string
	listIds                 []uint
	attributes              *string
	preconfirmSubscriptions *bool
}

func (s *UpdateSubscriberService) Id(id uint) *UpdateSubscriberService {
	s.id = id
	return s
}

func (s *UpdateSubscriberService) Email(email string) *UpdateSubscriberService {
	s.email = email
	return s
}

func (s *UpdateSubscriberService) Name(name string) *UpdateSubscriberService {
	s.name = name
	return s
}

func (s *UpdateSubscriberService) Status(status string) *UpdateSubscriberService {
	s.status = status
	return s
}

func (s *UpdateSubscriberService) ListIds(listIds []uint) *UpdateSubscriberService {
	s.listIds = listIds
	return s
}

func (s *UpdateSubscriberService) Attributes(attributes string) *UpdateSubscriberService {
	s.attributes = &attributes
	return s
}

func (s *UpdateSubscriberService) PreconfirmSubscriptions(preconfirmSubscriptions bool) *UpdateSubscriberService {
	s.preconfirmSubscriptions = &preconfirmSubscriptions
	return s
}

func (s *UpdateSubscriberService) Do(ctx context.Context, opts ...requestOption) (*Subscriber, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/subscribers/%d", s.id),
	}

	r.setJsonParam("email", s.email)
	r.setJsonParam("name", s.name)
	r.setJsonParam("status", s.status)

	if len(s.listIds) > 0 {
		r.setJsonParam("lists", s.listIds)
	}
	if s.attributes != nil {
		r.setJsonParam("attributes", s.attributes)
	}
	if s.preconfirmSubscriptions != nil {
		r.setJsonParam("preconfirm_subscriptions", s.preconfirmSubscriptions)
	}

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string]interface{}
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	finalBytes, err := json.Marshal(res["data"])

	if err != nil {
		return nil, err
	}

	finalRes := &Subscriber{}
	err = json.Unmarshal(finalBytes, finalRes)

	if err != nil {
		return nil, err
	}

	return finalRes, nil
}

type BlocklistsSubscriberService struct {
	c  ClientInterface
	id uint
}

func (s *BlocklistsSubscriberService) Id(id uint) *BlocklistsSubscriberService {
	s.id = id
	return s
}

func (s *BlocklistsSubscriberService) Do(ctx context.Context, opts ...requestOption) (*bool, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/subscribers/%d/blocklists", s.id),
	}

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string]*bool
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	return res["data"], nil
}

type BlocklistsQuerySubscriberService struct {
	c     ClientInterface
	query string
}

func (s *BlocklistsQuerySubscriberService) Query(query string) *BlocklistsQuerySubscriberService {
	s.query = query
	return s
}

func (s *BlocklistsQuerySubscriberService) Do(ctx context.Context, opts ...requestOption) (*bool, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/subscribers/blocklists",
	}

	r.setJsonParam("query", s.query)

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string]*bool
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	return res["data"], nil
}

type DeleteSubscriberService struct {
	c  ClientInterface
	id uint
}

func (s *DeleteSubscriberService) Id(id uint) *DeleteSubscriberService {
	s.id = id
	return s
}

func (s *DeleteSubscriberService) Do(ctx context.Context, opts ...requestOption) (*bool, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/subscribers/%d", s.id),
	}

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string]*bool
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	return res["data"], nil
}

type DeleteSubscribersService struct {
	c   ClientInterface
	ids []uint
}

func (s *DeleteSubscribersService) Ids(ids []uint) *DeleteSubscribersService {
	s.ids = ids
	return s
}

func (s *DeleteSubscribersService) Do(ctx context.Context, opts ...requestOption) (*bool, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/subscribers",
	}

	r.setParamList("id", s.ids)

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string]*bool
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	return res["data"], nil
}

type DeleteSubscribersQueryService struct {
	c     ClientInterface
	query string
}

func (s *DeleteSubscribersQueryService) Query(query string) *DeleteSubscribersQueryService {
	s.query = query
	return s
}

func (s *DeleteSubscribersQueryService) Do(ctx context.Context, opts ...requestOption) (*bool, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/subscribers/query/delete",
	}

	r.setParam("query", s.query)

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string]*bool
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	return res["data"], nil
}
