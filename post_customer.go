package fortnox

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewPostCustomerRequest() PostCustomerRequest {
	r := PostCustomerRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewPostCustomerQueryParams()
	r.pathParams = r.NewPostCustomerPathParams()
	r.requestBody = r.NewPostCustomerRequestBody()
	return r
}

type PostCustomerRequest struct {
	client      *Client
	queryParams *PostCustomerQueryParams
	pathParams  *PostCustomerPathParams
	method      string
	headers     http.Header
	requestBody PostCustomerRequestBody
}

func (r PostCustomerRequest) NewPostCustomerQueryParams() *PostCustomerQueryParams {
	return &PostCustomerQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type PostCustomerQueryParams struct {
}

func (p PostCustomerQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostCustomerRequest) QueryParams() *PostCustomerQueryParams {
	return r.queryParams
}

func (r PostCustomerRequest) NewPostCustomerPathParams() *PostCustomerPathParams {
	return &PostCustomerPathParams{}
}

type PostCustomerPathParams struct {
}

func (p *PostCustomerPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostCustomerRequest) PathParams() *PostCustomerPathParams {
	return r.pathParams
}

func (r *PostCustomerRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostCustomerRequest) Method() string {
	return r.method
}

func (r PostCustomerRequest) NewPostCustomerRequestBody() PostCustomerRequestBody {
	return PostCustomerRequestBody{}
}

type PostCustomerRequestBody struct {
	Customer Customer
}

func (r *PostCustomerRequest) RequestBody() *PostCustomerRequestBody {
	return &r.requestBody
}

func (r *PostCustomerRequest) SetRequestBody(body PostCustomerRequestBody) {
	r.requestBody = body
}

func (r *PostCustomerRequest) NewResponseBody() *PostCustomerResponseBody {
	return &PostCustomerResponseBody{}
}

type PostCustomerResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Customer        Customer
}

func (r *PostCustomerRequest) URL() url.URL {
	return r.client.GetEndpointURL("/customers", r.PathParams())
}

func (r *PostCustomerRequest) Do(ctx context.Context) (PostCustomerResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r.Method(), r.URL(), r.RequestBody())
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
