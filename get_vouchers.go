package fortnox

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetVouchersRequest() GetVouchersRequest {
	r := GetVouchersRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetVouchersQueryParams()
	r.pathParams = r.NewGetVouchersPathParams()
	r.requestBody = r.NewGetVouchersRequestBody()
	return r
}

type GetVouchersRequest struct {
	client      *Client
	queryParams *GetVouchersQueryParams
	pathParams  *GetVouchersPathParams
	method      string
	headers     http.Header
	requestBody GetVouchersRequestBody
}

func (r GetVouchersRequest) NewGetVouchersQueryParams() *GetVouchersQueryParams {
	return &GetVouchersQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetVouchersQueryParams struct {
}

func (p GetVouchersQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetVouchersRequest) QueryParams() *GetVouchersQueryParams {
	return r.queryParams
}

func (r GetVouchersRequest) NewGetVouchersPathParams() *GetVouchersPathParams {
	return &GetVouchersPathParams{}
}

type GetVouchersPathParams struct {
}

func (p *GetVouchersPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetVouchersRequest) PathParams() *GetVouchersPathParams {
	return r.pathParams
}

func (r *GetVouchersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetVouchersRequest) Method() string {
	return r.method
}

func (r GetVouchersRequest) NewGetVouchersRequestBody() GetVouchersRequestBody {
	return GetVouchersRequestBody{}
}

type GetVouchersRequestBody struct{}

func (r *GetVouchersRequest) RequestBody() *GetVouchersRequestBody {
	return &r.requestBody
}

func (r *GetVouchersRequest) SetRequestBody(body GetVouchersRequestBody) {
	r.requestBody = body
}

func (r *GetVouchersRequest) NewResponseBody() *GetVouchersResponseBody {
	return &GetVouchersResponseBody{}
}

type GetVouchersResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Vouchers        Vouchers
}

func (r *GetVouchersRequest) URL() url.URL {
	return r.client.GetEndpointURL("/vouchers", r.PathParams())
}

func (r *GetVouchersRequest) Do(ctx context.Context) (GetVouchersResponseBody, error) {
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
