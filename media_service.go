package listmonk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Media struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Uuid      string    `json:"uuid"`
	Filename  string    `json:"filename"`
	ThumbUrl  string    `json:"thumb_url"`
	Uri       string    `json:"uri"`
}

type GetMediaService struct {
	c ClientInterface
}

func (s *GetMediaService) Do(ctx context.Context, opts ...requestOption) ([]*Media, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/media",
	}

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string][]*Media
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	return res["data"], nil
}

type CreateMediaService struct {
	c    ClientInterface
	file []byte
}

func (s *CreateMediaService) File(file []byte) *CreateMediaService {
	s.file = file
	return s
}

func (s *CreateMediaService) Do(ctx context.Context, opts ...requestOption) (*Media, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/media",
	}

	r.setJsonParam("file", s.file)

	bytes, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	var res map[string]*Media
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}

	return res["data"], nil
}

type DeleteMediaService struct {
	c  ClientInterface
	id uint
}

func (s *DeleteMediaService) Id(id uint) *DeleteMediaService {
	s.id = id
	return s
}

func (s *DeleteMediaService) Do(ctx context.Context, opts ...requestOption) error {
	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/media/%d", s.id),
	}

	_, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return err
	}

	return nil
}
