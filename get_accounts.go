package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetAccountsRequest() GetAccountsRequest {
	r := GetAccountsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetAccountsQueryParams()
	r.pathParams = r.NewGetAccountsPathParams()
	r.requestBody = r.NewGetAccountsRequestBody()
	return r
}

type GetAccountsRequest struct {
	client      *Client
	queryParams *GetAccountsQueryParams
	pathParams  *GetAccountsPathParams
	method      string
	headers     http.Header
	requestBody GetAccountsRequestBody
}

func (r GetAccountsRequest) NewGetAccountsQueryParams() *GetAccountsQueryParams {
	return &GetAccountsQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetAccountsQueryParams struct {
	Page int `schema:"page,omitempty"`
}

func (p GetAccountsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAccountsRequest) QueryParams() *GetAccountsQueryParams {
	return r.queryParams
}

func (r GetAccountsRequest) NewGetAccountsPathParams() *GetAccountsPathParams {
	return &GetAccountsPathParams{}
}

type GetAccountsPathParams struct {
}

func (p *GetAccountsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAccountsRequest) PathParams() *GetAccountsPathParams {
	return r.pathParams
}

func (r *GetAccountsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAccountsRequest) Method() string {
	return r.method
}

func (r GetAccountsRequest) NewGetAccountsRequestBody() GetAccountsRequestBody {
	return GetAccountsRequestBody{}
}

type GetAccountsRequestBody struct{}

func (r *GetAccountsRequest) RequestBody() *GetAccountsRequestBody {
	return &r.requestBody
}

func (r *GetAccountsRequest) SetRequestBody(body GetAccountsRequestBody) {
	r.requestBody = body
}

func (r *GetAccountsRequest) NewResponseBody() *GetAccountsResponseBody {
	return &GetAccountsResponseBody{}
}

type GetAccountsResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Accounts        Accounts
}

func (r *GetAccountsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/accounts", r.PathParams())
}

func (r *GetAccountsRequest) Do() (GetAccountsResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *GetAccountsRequest) All() (GetAccountsResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	for resp.MetaInformation.CurrentPage < resp.MetaInformation.TotalPages {
		r.QueryParams().Page = resp.MetaInformation.CurrentPage + 1
		page, err := r.Do()
		if err != nil {
			return resp, err
		}

		resp.Accounts = append(resp.Accounts, page.Accounts...)

		resp.MetaInformation.CurrentPage++
	}

	return resp, nil
}
