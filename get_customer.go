package fortnox

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetCustomerRequest() GetCustomerRequest {
	r := GetCustomerRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetCustomerQueryParams()
	r.pathParams = r.NewGetCustomerPathParams()
	r.requestBody = r.NewGetCustomerRequestBody()
	return r
}

type GetCustomerRequest struct {
	client      *Client
	queryParams *GetCustomerQueryParams
	pathParams  *GetCustomerPathParams
	method      string
	headers     http.Header
	requestBody GetCustomerRequestBody
}

func (r GetCustomerRequest) NewGetCustomerQueryParams() *GetCustomerQueryParams {
	return &GetCustomerQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetCustomerQueryParams struct{}

func (p GetCustomerQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomerRequest) QueryParams() *GetCustomerQueryParams {
	return r.queryParams
}

func (r GetCustomerRequest) NewGetCustomerPathParams() *GetCustomerPathParams {
	return &GetCustomerPathParams{}
}

type GetCustomerPathParams struct {
	CustomerNumber string
}

func (p *GetCustomerPathParams) Params() map[string]string {
	return map[string]string{
		"CustomerNumber": p.CustomerNumber,
	}
}

func (r *GetCustomerRequest) PathParams() *GetCustomerPathParams {
	return r.pathParams
}

func (r *GetCustomerRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomerRequest) Method() string {
	return r.method
}

func (r GetCustomerRequest) NewGetCustomerRequestBody() GetCustomerRequestBody {
	return GetCustomerRequestBody{}
}

type GetCustomerRequestBody struct{}

func (r *GetCustomerRequest) RequestBody() *GetCustomerRequestBody {
	return &r.requestBody
}

func (r *GetCustomerRequest) SetRequestBody(body GetCustomerRequestBody) {
	r.requestBody = body
}

func (r *GetCustomerRequest) NewResponseBody() *GetCustomerResponseBody {
	return &GetCustomerResponseBody{}
}

type GetCustomerResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Customer        Customer
}

func (r *GetCustomerRequest) URL() url.URL {
	return r.client.GetEndpointURL("/customers/{{.CustomerNumber}}", r.PathParams())
}

func (r *GetCustomerRequest) Do(ctx context.Context) (GetCustomerResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r.Method(), r.URL(), nil)
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
