package listmonk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PostTransactionalService struct {
	c                ClientInterface
	subscriberEmail  string
	subscriberId     uint
	subscriberEmails []string
	subscriberIds    []uint
	templateId       uint
	fromEmail        string
	data             map[string]string
	headers          map[string]string
	messenger        string
	content_type     string
}

func (s *PostTransactionalService) SubscriberEmail(email string) *PostTransactionalService {
	s.subscriberEmail = email
	return s
}

func (s *PostTransactionalService) SubscriberId(id uint) *PostTransactionalService {
	s.subscriberId = id
	return s
}

func (s *PostTransactionalService) SubscriberEmails(emails []string) *PostTransactionalService {
	s.subscriberEmails = emails
	return s
}

func (s *PostTransactionalService) SubscriberIds(ids []uint) *PostTransactionalService {
	s.subscriberIds = ids
	return s
}

func (s *PostTransactionalService) TemplateId(id uint) *PostTransactionalService {
	s.templateId = id
	return s
}

func (s *PostTransactionalService) FromEmail(email string) *PostTransactionalService {
	s.fromEmail = email
	return s
}

func (s *PostTransactionalService) Data(data map[string]string) *PostTransactionalService {
	s.data = data
	return s
}

func (s *PostTransactionalService) Headers(headers map[string]string) *PostTransactionalService {
	s.headers = headers
	return s
}

func (s *PostTransactionalService) Messenger(messenger string) *PostTransactionalService {
	s.messenger = messenger
	return s
}

func (s *PostTransactionalService) ContentType(content_type string) *PostTransactionalService {
	s.content_type = content_type
	return s
}

func (s *PostTransactionalService) Do(ctx context.Context, opts ...requestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/tx",
	}

	if s.subscriberEmail != "" {
		r.setJsonParam("subscriber_email", s.subscriberEmail)
	}

	if s.subscriberId != 0 {
		r.setJsonParam("subscriber_id", s.subscriberId)
	}

	if len(s.subscriberEmails) > 0 {
		quoted := make([]string, len(s.subscriberEmails))
		for i, s := range s.subscriberEmails {
			quoted[i] = fmt.Sprintf("\"%s\"", s)
		}

		emails := "[" + strings.Join(quoted, ", ") + "]"

		r.setJsonParam("subscriber_emails", emails)
	}

	if len(s.subscriberIds) > 0 {
		strIds := make([]string, len(s.subscriberIds))

		for i, id := range s.subscriberIds {
			strIds[i] = fmt.Sprintf("%d", id)
		}

		ids := "[" + strings.Join(strIds, ", ") + "]"

		r.setJsonParam("subscriber_ids", ids)
	}

	if s.templateId != 0 {
		r.setJsonParam("template_id", s.templateId)
	}

	if s.fromEmail != "" {
		r.setJsonParam("from_email", s.fromEmail)
	}

	if len(s.data) > 0 {
		dataJson, err := json.Marshal(s.data)

		if err != nil {
			return err
		}

		r.setJsonParam("data", string(dataJson))
	}

	if len(s.headers) > 0 {
		elements := make([]string, 0)

		for k, v := range s.headers {
			elements = append(elements, fmt.Sprintf("{%s: %s}", fmt.Sprintf("\"%s\"", k), fmt.Sprintf("\"%s\"", v)))
		}

		headers := "[" + strings.Join(elements, ", ") + "]"

		r.setJsonParam("headers", headers)
	}

	if s.messenger != "" {
		r.setJsonParam("messenger", s.messenger)
	}

	if s.content_type != "" {
		r.setJsonParam("content_type", s.content_type)
	}

	_, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return err
	}

	return nil
}
