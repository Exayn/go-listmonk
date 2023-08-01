package listmonk

import (
	"context"
	"encoding/json"
	"net/http"
)

type ImportStatus struct {
	Name     string `json:"name"`
	Total    uint   `json:"total"`
	Imported uint   `json:"imported"`
	Status   string `json:"status"`
}

type GetImportStatusService struct {
	c ClientInterface
}

func (s *GetImportStatusService) Do(ctx context.Context, opts ...requestOption) (*ImportStatus, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/import/subscribers",
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var result map[string]*ImportStatus
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result["data"], nil
}

type GetImportLogsService struct {
	c ClientInterface
}

func (s *GetImportLogsService) Do(ctx context.Context, opts ...requestOption) (*string, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/import/subscribers/logs",
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var result map[string]*string
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result["data"], nil
}

type ImportSubscribersService struct {
	c      ClientInterface
	file   []byte
	params map[string]interface{}
}

func (s *ImportSubscribersService) File(file []byte) *ImportSubscribersService {
	s.file = file
	return s
}

func (s *ImportSubscribersService) Params(params map[string]interface{}) *ImportSubscribersService {
	s.params = params
	return s
}

func (s *ImportSubscribersService) Do(ctx context.Context, opts ...requestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/import/subscribers",
	}

	param, err := json.Marshal(s.params)

	if err != nil {
		return err
	}

	r.setFormParam("params", param)
	r.setFormParam("file", s.file)

	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

type DeleteImportService struct {
	c ClientInterface
}

func (s *DeleteImportService) Do(ctx context.Context, opts ...requestOption) error {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/import/subscribers",
	}

	_, err := s.c.callAPI(ctx, r, opts...)

	return err
}
