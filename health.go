package listmonk

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetHealthService struct {
	c ClientInterface
}

func (s *GetHealthService) Do(ctx context.Context, opts ...requestOption) (*bool, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/health",
	}

	bytes, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	var res map[string]bool
	err = json.Unmarshal(bytes, &res)

	if err != nil {
		return nil, err
	}
	var result bool = res["data"]
	return &result, nil
}
