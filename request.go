package listmonk

import (
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
	form     url.Values
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

func (r *request) setParamList(baseParam string, params ...interface{}) *request {
	for index, value := range params {
		r.setParam(fmt.Sprintf("%s[%d]", baseParam, index), value)
	}
	return r
}

// setFormParam set param with key/value to request form body
func (r *request) setFormParam(key string, value interface{}) *request {
	if r.form == nil {
		r.form = url.Values{}
	}
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) validate() {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.form == nil {
		r.form = url.Values{}
	}
	if r.header == nil {
		r.header = http.Header{}
	}
}

func (r *request) toHttpRequest(baseUrl string, username, password *string, ctx context.Context, opts ...requestOption) (req *http.Request, err error) {
	r.validate()

	var body io.Reader

	if len(r.form) > 0 {
		body = strings.NewReader(r.form.Encode())
		r.header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		body = r.body
	}

	req, err = http.NewRequest(r.method, fmt.Sprintf("%s/api/%s", baseUrl, strings.TrimPrefix(r.endpoint, "/")), body)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = r.query.Encode()
	req.Header = r.header

	if username != nil && password != nil {
		req.SetBasicAuth(*username, *password)
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
