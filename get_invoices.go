package fortnox

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetInvoicesRequest() GetInvoicesRequest {
	r := GetInvoicesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetInvoicesQueryParams()
	r.pathParams = r.NewGetInvoicesPathParams()
	r.requestBody = r.NewGetInvoicesRequestBody()
	return r
}

type GetInvoicesRequest struct {
	client      *Client
	queryParams *GetInvoicesQueryParams
	pathParams  *GetInvoicesPathParams
	method      string
	headers     http.Header
	requestBody GetInvoicesRequestBody
}

func (r GetInvoicesRequest) NewGetInvoicesQueryParams() *GetInvoicesQueryParams {
	return &GetInvoicesQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetInvoicesQueryParams struct {
	FromDate Date `schema:"fromdate"`
	ToDate   Date `schema:"todate"`
}

func (p GetInvoicesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoicesRequest) QueryParams() *GetInvoicesQueryParams {
	return r.queryParams
}

func (r GetInvoicesRequest) NewGetInvoicesPathParams() *GetInvoicesPathParams {
	return &GetInvoicesPathParams{}
}

type GetInvoicesPathParams struct {
}

func (p *GetInvoicesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetInvoicesRequest) PathParams() *GetInvoicesPathParams {
	return r.pathParams
}

func (r *GetInvoicesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoicesRequest) Method() string {
	return r.method
}

func (r GetInvoicesRequest) NewGetInvoicesRequestBody() GetInvoicesRequestBody {
	return GetInvoicesRequestBody{}
}

type GetInvoicesRequestBody struct{}

func (r *GetInvoicesRequest) RequestBody() *GetInvoicesRequestBody {
	return &r.requestBody
}

func (r *GetInvoicesRequest) SetRequestBody(body GetInvoicesRequestBody) {
	r.requestBody = body
}

func (r *GetInvoicesRequest) NewResponseBody() *GetInvoicesResponseBody {
	return &GetInvoicesResponseBody{}
}

type GetInvoicesResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoices        Invoices
}

func (r *GetInvoicesRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices", r.PathParams())
}

func (r *GetInvoicesRequest) Do(ctx context.Context) (GetInvoicesResponseBody, error) {
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
