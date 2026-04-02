package fortnox

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetCustomersRequest() GetCustomersRequest {
	r := GetCustomersRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetCustomersQueryParams()
	r.pathParams = r.NewGetCustomersPathParams()
	r.requestBody = r.NewGetCustomersRequestBody()
	return r
}

type GetCustomersRequest struct {
	client      *Client
	queryParams *GetCustomersQueryParams
	pathParams  *GetCustomersPathParams
	method      string
	headers     http.Header
	requestBody GetCustomersRequestBody
}

func (r GetCustomersRequest) NewGetCustomersQueryParams() *GetCustomersQueryParams {
	return &GetCustomersQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetCustomersQueryParams struct {
	CustomerNumber string `schema:"customernumber"`
	Name           string `schema:"name"`
	Email          string `schema:"email"`
}

func (p GetCustomersQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomersRequest) QueryParams() *GetCustomersQueryParams {
	return r.queryParams
}

func (r GetCustomersRequest) NewGetCustomersPathParams() *GetCustomersPathParams {
	return &GetCustomersPathParams{}
}

type GetCustomersPathParams struct {
}

func (p *GetCustomersPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomersRequest) PathParams() *GetCustomersPathParams {
	return r.pathParams
}

func (r *GetCustomersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomersRequest) Method() string {
	return r.method
}

func (r GetCustomersRequest) NewGetCustomersRequestBody() GetCustomersRequestBody {
	return GetCustomersRequestBody{}
}

type GetCustomersRequestBody struct{}

func (r *GetCustomersRequest) RequestBody() *GetCustomersRequestBody {
	return &r.requestBody
}

func (r *GetCustomersRequest) SetRequestBody(body GetCustomersRequestBody) {
	r.requestBody = body
}

func (r *GetCustomersRequest) NewResponseBody() *GetCustomersResponseBody {
	return &GetCustomersResponseBody{}
}

type GetCustomersResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Customers       Customers
}

func (r *GetCustomersRequest) URL() url.URL {
	return r.client.GetEndpointURL("/customers", r.PathParams())
}

func (r *GetCustomersRequest) Do(ctx context.Context) (GetCustomersResponseBody, error) {
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
