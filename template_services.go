package listmonk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Template struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Body      string    `json:"body"`
	Type      string    `json:"type"`
	IsDefault bool      `json:"is_default"`
}

type GetTemplatesService struct {
	c ClientInterface
}

func (s *GetTemplatesService) Do(ctx context.Context, opts ...requestOption) ([]*Template, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/templates",
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

	finalRes := make([]*Template, 0)
	err = json.Unmarshal(finalBytes, &finalRes)

	if err != nil {
		return nil, err
	}

	return finalRes, nil
}

type GetTemplateService struct {
	c  ClientInterface
	id uint
}

func (s *GetTemplateService) Id(id uint) *GetTemplateService {
	s.id = id
	return s
}

func (s *GetTemplateService) Do(ctx context.Context, opts ...requestOption) (*Template, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/templates/%d", s.id),
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

	finalRes := &Template{}
	err = json.Unmarshal(finalBytes, &finalRes)

	if err != nil {
		return nil, err
	}

	return finalRes, nil
}

type GetTemplatePreviewService struct {
	c  ClientInterface
	id uint
}

func (s *GetTemplatePreviewService) Id(id uint) *GetTemplatePreviewService {
	s.id = id
	return s
}

func (s *GetTemplatePreviewService) Do(ctx context.Context, opts ...requestOption) (string, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/templates/%d/preview", s.id),
	}

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

type UpdateTemplateAsDefaultService struct {
	c  ClientInterface
	id uint
}

func (s *UpdateTemplateAsDefaultService) Id(id uint) *UpdateTemplateAsDefaultService {
	s.id = id
	return s
}

func (s *UpdateTemplateAsDefaultService) Do(ctx context.Context, opts ...requestOption) (*Template, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/templates/%d/default", s.id),
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

	finalRes := &Template{}
	err = json.Unmarshal(finalBytes, &finalRes)

	if err != nil {
		return nil, err
	}

	return finalRes, nil
}

type DeleteTemplateService struct {
	c  ClientInterface
	id uint
}

func (s *DeleteTemplateService) Id(id uint) *DeleteTemplateService {
	s.id = id
	return s
}

func (s *DeleteTemplateService) Do(ctx context.Context, opts ...requestOption) error {
	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/templates/%d", s.id),
	}

	_, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return err
	}

	return nil
}
