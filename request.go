package listmonk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// request define an API request
type request struct {
	method   string
	endpoint string
	query    url.Values
	json     map[string]interface{}
	header   http.Header
	body     io.Reader
}

// setParam set param with key/value to query string
func (r *request) setParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}

	if reflect.TypeOf(value).Kind() == reflect.Slice {
		v, err := json.Marshal(value)
		if err == nil {
			value = string(v)
		}
	}

	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) setParamList(baseParam string, params []uint) *request {
	for _, value := range params {
		r.setParam(baseParam, fmt.Sprintf("%d", value))
	}
	return r
}

// setFormParam set param with key/value to request form body
func (r *request) setJsonParam(key string, value interface{}) *request {
	if r.json == nil {
		r.json = map[string]interface{}{}
	}

	r.json[key] = value

	return r
}

func (r *request) validate() {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.json == nil {
		r.json = map[string]interface{}{}
	}
	if r.header == nil {
		r.header = http.Header{}
	}
}

func (r *request) toHttpRequest(baseUrl string, username, password *string, ctx context.Context, opts ...requestOption) (req *http.Request, err error) {
	r.validate()

	var body io.Reader

	if r.json != nil {
		jsonBytes, err := json.Marshal(r.json)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonBytes)
		r.header.Set("Content-Type", "application/json")
	}

	req, err = http.NewRequest(r.method, fmt.Sprintf("%s/api/%s", baseUrl, strings.TrimPrefix(r.endpoint, "/")), body)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = r.query.Encode()

	if username != nil && password != nil {
		req.SetBasicAuth(*username, *password)
	}

	for key, values := range r.header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	for _, opt := range opts {
		opt(r)
	}

	req = req.WithContext(ctx)

	return req, nil
}

// RequestOption define option type for request
type requestOption func(*request)

// WithHeader set or add a header value to the request
func WithHeader(key, value string, replace bool) requestOption {
	return func(r *request) {
		if r.header == nil {
			r.header = http.Header{}
		}
		if replace {
			r.header.Set(key, value)
		} else {
			r.header.Add(key, value)
		}
	}
}

// WithHeaders set or replace the headers of the request
func WithHeaders(header http.Header) requestOption {
	return func(r *request) {
		r.header = header.Clone()
	}
}
