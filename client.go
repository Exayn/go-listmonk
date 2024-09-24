package listmonk

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ClientInterface interface {
	callAPI(ctx context.Context, r *request, opts ...requestOption) ([]byte, error)
}

type Client struct {
	ClientInterface
	baseURL    string
	username   *string
	password   *string
	httpClient *http.Client
}

func NewClient(baseURL string, username, password *string) *Client {
	return &Client{
		baseURL:    baseURL,
		username:   username,
		password:   password,
		httpClient: &http.Client{},
	}
}

func NewClientWithCustomHTTPClient(baseURL string, username, password *string, httpClient *http.Client) *Client {
	return &Client{
		baseURL:    baseURL,
		username:   username,
		password:   password,
		httpClient: httpClient,
	}
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...requestOption) ([]byte, error) {
	req, err := r.toHttpRequest(c.baseURL, c.username, c.password, ctx, opts...)

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

func (c *Client) NewGetHealthService() *GetHealthService {
	return &GetHealthService{
		c: c,
	}
}

func (c *Client) NewGetSubscribersService() *GetSubscribersService {
	return &GetSubscribersService{
		c: c,
	}
}

func (c *Client) NewGetSubscriberService() *GetSubscriberService {
	return &GetSubscriberService{
		c: c,
	}
}

func (c *Client) NewCreateSubscriberService() *CreateSubscriberService {
	return &CreateSubscriberService{
		c: c,
	}
}

func (c *Client) NewUpdateSubscribersListsService() *UpdateSubscribersListsService {
	return &UpdateSubscribersListsService{
		c: c,
	}
}

func (c *Client) NewUpdateSubscriberService() *UpdateSubscriberService {
	return &UpdateSubscriberService{
		c: c,
	}
}

func (c *Client) NewBlocklistsSubscriberService() *BlocklistsSubscriberService {
	return &BlocklistsSubscriberService{
		c: c,
	}
}

func (c *Client) NewBlocklistsQuerySubscriberService() *BlocklistsQuerySubscriberService {
	return &BlocklistsQuerySubscriberService{
		c: c,
	}
}

func (c *Client) NewDeleteSubscriberService() *DeleteSubscriberService {
	return &DeleteSubscriberService{
		c: c,
	}
}

func (c *Client) NewDeleteSubscribersService() *DeleteSubscribersService {
	return &DeleteSubscribersService{
		c: c,
	}
}

func (c *Client) NewDeleteSubscribersQueryService() *DeleteSubscribersQueryService {
	return &DeleteSubscribersQueryService{
		c: c,
	}
}

func (c *Client) NewGetListsService() *GetListsService {
	return &GetListsService{
		c: c,
	}
}

func (c *Client) NewGetListService() *GetListService {
	return &GetListService{
		c: c,
	}
}

func (c *Client) NewCreateListService() *CreateListService {
	return &CreateListService{
		c: c,
	}
}

func (c *Client) NewUpdateListService() *UpdateListService {
	return &UpdateListService{
		c: c,
	}
}

func (c *Client) NewDeleteListService() *DeleteListService {
	return &DeleteListService{
		c: c,
	}
}

func (c *Client) NewGetCampaignsService() *GetCampaignsService {
	return &GetCampaignsService{
		c: c,
	}
}

func (c *Client) NewGetCampaignService() *GetCampaignService {
	return &GetCampaignService{
		c: c,
	}
}

func (c *Client) NewCreateCampaignService() *CreateCampaignService {
	return &CreateCampaignService{
		c: c,
	}
}

func (c *Client) NewUpdateCampaignStatusService() *UpdateCampaignStatusService {
	return &UpdateCampaignStatusService{
		c: c,
	}
}

func (c *Client) NewDeleteCampaignService() *DeleteCampaignService {
	return &DeleteCampaignService{
		c: c,
	}
}

func (c *Client) NewGetMediaService() *GetMediaService {
	return &GetMediaService{
		c: c,
	}
}

func (c *Client) NewCreateMediaService() *CreateMediaService {
	return &CreateMediaService{
		c: c,
	}
}

func (c *Client) NewDeleteMediaService() *DeleteMediaService {
	return &DeleteMediaService{
		c: c,
	}
}

func (c *Client) NewGetTemplatesService() *GetTemplatesService {
	return &GetTemplatesService{
		c: c,
	}
}

func (c *Client) NewGetTemplateService() *GetTemplateService {
	return &GetTemplateService{
		c: c,
	}
}

func (c *Client) NewGetTemplatePreviewService() *GetTemplatePreviewService {
	return &GetTemplatePreviewService{
		c: c,
	}
}

func (c *Client) NewUpdateTemplateAsDefaultService() *UpdateTemplateAsDefaultService {
	return &UpdateTemplateAsDefaultService{
		c: c,
	}
}

func (c *Client) NewDeleteTemplateService() *DeleteTemplateService {
	return &DeleteTemplateService{
		c: c,
	}
}

func (c *Client) NewPostTransactionalService() *PostTransactionalService {
	return &PostTransactionalService{
		c: c,
	}
}
