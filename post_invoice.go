package fortnox

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewPostInvoiceRequest() PostInvoiceRequest {
	r := PostInvoiceRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewPostInvoiceQueryParams()
	r.pathParams = r.NewPostInvoicePathParams()
	r.requestBody = r.NewPostInvoiceRequestBody()
	return r
}

type PostInvoiceRequest struct {
	client      *Client
	queryParams *PostInvoiceQueryParams
	pathParams  *PostInvoicePathParams
	method      string
	headers     http.Header
	requestBody PostInvoiceRequestBody
}

func (r PostInvoiceRequest) NewPostInvoiceQueryParams() *PostInvoiceQueryParams {
	return &PostInvoiceQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type PostInvoiceQueryParams struct {
}

func (p PostInvoiceQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostInvoiceRequest) QueryParams() *PostInvoiceQueryParams {
	return r.queryParams
}

func (r PostInvoiceRequest) NewPostInvoicePathParams() *PostInvoicePathParams {
	return &PostInvoicePathParams{}
}

type PostInvoicePathParams struct {
}

func (p *PostInvoicePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostInvoiceRequest) PathParams() *PostInvoicePathParams {
	return r.pathParams
}

func (r *PostInvoiceRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostInvoiceRequest) Method() string {
	return r.method
}

func (r PostInvoiceRequest) NewPostInvoiceRequestBody() PostInvoiceRequestBody {
	return PostInvoiceRequestBody{}
}

type PostInvoiceRequestBody struct {
	Invoice Invoice
}

func (r *PostInvoiceRequest) RequestBody() *PostInvoiceRequestBody {
	return &r.requestBody
}

func (r *PostInvoiceRequest) SetRequestBody(body PostInvoiceRequestBody) {
	r.requestBody = body
}

func (r *PostInvoiceRequest) NewResponseBody() *PostInvoiceResponseBody {
	return &PostInvoiceResponseBody{}
}

type PostInvoiceResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Invoice         Invoice
}

func (r *PostInvoiceRequest) URL() url.URL {
	return r.client.GetEndpointURL("/invoices", r.PathParams())
}

func (r *PostInvoiceRequest) Do(ctx context.Context) (PostInvoiceResponseBody, error) {
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
