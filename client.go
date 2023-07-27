package listmonk

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func NewClientWithCustomHTTPClient(baseURL string, httpClient *http.Client) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) ([]byte, error) {
	req, err := r.toHttpRequest(c.baseURL, ctx, opts...)

	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(APIError)

		e := json.Unmarshal(data, apiErr)

		if e != nil {
			return nil, e
		}

		apiErr.Code = res.StatusCode

		return nil, apiErr
	}

	return data, nil
}

func (c *Client) NewGetSubscriberListService() *GetSubscriberListService {
	return &GetSubscriberListService{
		c: c,
	}
}
