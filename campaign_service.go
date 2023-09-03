package listmonk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CampaignList struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type Campaign struct {
	Id          uint           `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	CampaignId  uint           `json:"CampaignID"`
	View        uint           `json:"view"`
	Clicks      uint           `json:"clicks"`
	Lists       []CampaignList `json:"lists"`
	StartedAt   time.Time      `json:"started_at"`
	ToSend      uint           `json:"to_send"`
	Sent        uint           `json:"sent"`
	Uuid        string         `json:"uuid"`
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	Subject     string         `json:"subject"`
	FromEmail   string         `json:"from_email"`
	Body        string         `json:"body"`
	SendAt      time.Time      `json:"send_at"`
	Status      string         `json:"status"`
	ContentType string         `json:"content_type"`
	Tags        []string       `json:"tags"`
	TemplateId  uint           `json:"template_id"`
	Messenger   string         `json:"messenger"`
}

type GetCampaignsService struct {
	c       ClientInterface
	page    *uint
	perPage *string
	query   *string
	orderBy *string
	order   *string
}

func (s *GetCampaignsService) Page(page uint) *GetCampaignsService {
	s.page = &page
	return s
}

func (s *GetCampaignsService) PerPage(perPage string) *GetCampaignsService {
	s.perPage = &perPage
	return s
}

func (s *GetCampaignsService) Query(query string) *GetCampaignsService {
	s.query = &query
	return s
}

func (s *GetCampaignsService) OrderBy(orderBy string) *GetCampaignsService {
	s.orderBy = &orderBy
	return s
}

func (s *GetCampaignsService) Order(order string) *GetCampaignsService {
	s.order = &order
	return s
}

func (s *GetCampaignsService) Do(ctx context.Context, opts ...requestOption) ([]*Campaign, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/campaigns",
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

	if s.orderBy != nil {
		r.setParam("order_by", *s.orderBy)
	}

	if s.order != nil {
		r.setParam("order", *s.order)
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

	finalRes := make([]*Campaign, 0)
	err = json.Unmarshal(finalBytes, &finalRes)

	if err != nil {
		return nil, err
	}

	return finalRes, nil
}

type GetCampaignService struct {
	c  ClientInterface
	id uint
}

func (s *GetCampaignService) Id(id uint) *GetCampaignService {
	s.id = id
	return s
}

func (s *GetCampaignService) Do(ctx context.Context, opts ...requestOption) (*Campaign, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/campaigns/%d", s.id),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var result map[string]*Campaign
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result["data"], nil
}

type GetCampaignPreviewService struct {
	c  ClientInterface
	id uint
}

func (s *GetCampaignPreviewService) Id(id uint) *GetCampaignPreviewService {
	s.id = id
	return s
}

func (s *GetCampaignPreviewService) Do(ctx context.Context, opts ...requestOption) ([]byte, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/campaigns/%d/preview", s.id),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return data, err
}

type CreateCampaignService struct {
	c           ClientInterface
	name        string
	subject     string
	lists       []uint
	fromEmail   *string
	type_       string
	contentType string
	body        string
	altBody     *string
	sendAt      *time.Time
	messenger   *string
	templateId  *uint
	tags        []string
}

func (s *CreateCampaignService) Name(name string) *CreateCampaignService {
	s.name = name
	return s
}

func (s *CreateCampaignService) Subject(subject string) *CreateCampaignService {
	s.subject = subject
	return s
}

func (s *CreateCampaignService) Lists(lists []uint) *CreateCampaignService {
	s.lists = lists
	return s
}

func (s *CreateCampaignService) FromEmail(fromEmail string) *CreateCampaignService {
	s.fromEmail = &fromEmail
	return s
}

func (s *CreateCampaignService) Type(type_ string) *CreateCampaignService {
	s.type_ = type_
	return s
}

func (s *CreateCampaignService) ContentType(contentType string) *CreateCampaignService {
	s.contentType = contentType
	return s
}

func (s *CreateCampaignService) Body(body string) *CreateCampaignService {
	s.body = body
	return s
}

func (s *CreateCampaignService) AltBody(altBody string) *CreateCampaignService {
	s.altBody = &altBody
	return s
}

func (s *CreateCampaignService) SendAt(sendAt time.Time) *CreateCampaignService {
	s.sendAt = &sendAt
	return s
}

func (s *CreateCampaignService) Messenger(messenger string) *CreateCampaignService {
	s.messenger = &messenger
	return s
}

func (s *CreateCampaignService) TemplateId(templateId uint) *CreateCampaignService {
	s.templateId = &templateId
	return s
}

func (s *CreateCampaignService) Tags(tags []string) *CreateCampaignService {
	s.tags = tags
	return s
}

func (s *CreateCampaignService) Do(ctx context.Context, opts ...requestOption) (*Campaign, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/campaigns",
	}

	r.setParam("name", s.name)
	r.setParam("subject", s.subject)
	r.setParam("lists", s.lists)
	r.setParam("type", s.type_)
	r.setParam("content_type", s.contentType)
	r.setParam("body", s.body)

	if s.fromEmail != nil {
		r.setParam("from_email", *s.fromEmail)
	}

	if s.altBody != nil {
		r.setParam("alt_body", *s.altBody)
	}

	if s.sendAt != nil {
		r.setParam("send_at", *s.sendAt)
	}

	if s.messenger != nil {
		r.setParam("messenger", *s.messenger)
	}

	if s.templateId != nil {
		r.setParam("template_id", *s.templateId)
	}

	if s.tags != nil {
		r.setParam("tags", s.tags)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var result map[string]*Campaign
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result["data"], nil
}

type UpdateCampaignStatusService struct {
	c      ClientInterface
	id     uint
	status string
}

func (s *UpdateCampaignStatusService) Id(id uint) *UpdateCampaignStatusService {
	s.id = id
	return s
}

func (s *UpdateCampaignStatusService) Status(status string) *UpdateCampaignStatusService {
	s.status = status
	return s
}

func (s *UpdateCampaignStatusService) Do(ctx context.Context, opts ...requestOption) (*Campaign, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/campaigns/%d/status", s.id),
	}

	r.setParam("status", s.status)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var result map[string]*Campaign
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result["data"], nil
}

type DeleteCampaignService struct {
	c  ClientInterface
	id uint
}

func (s *DeleteCampaignService) Id(id uint) *DeleteCampaignService {
	s.id = id
	return s
}

func (s *DeleteCampaignService) Do(ctx context.Context, opts ...requestOption) error {
	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/campaigns/%d", s.id),
	}

	_, err := s.c.callAPI(ctx, r, opts...)
	return err
}
