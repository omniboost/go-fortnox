package fortnox

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewPostVoucherRequest() PostVoucherRequest {
	r := PostVoucherRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewPostVoucherQueryParams()
	r.pathParams = r.NewPostVoucherPathParams()
	r.requestBody = r.NewPostVoucherRequestBody()
	return r
}

type PostVoucherRequest struct {
	client      *Client
	queryParams *PostVoucherQueryParams
	pathParams  *PostVoucherPathParams
	method      string
	headers     http.Header
	requestBody PostVoucherRequestBody
}

func (r PostVoucherRequest) NewPostVoucherQueryParams() *PostVoucherQueryParams {
	return &PostVoucherQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type PostVoucherQueryParams struct {
}

func (p PostVoucherQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostVoucherRequest) QueryParams() *PostVoucherQueryParams {
	return r.queryParams
}

func (r PostVoucherRequest) NewPostVoucherPathParams() *PostVoucherPathParams {
	return &PostVoucherPathParams{}
}

type PostVoucherPathParams struct {
}

func (p *PostVoucherPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostVoucherRequest) PathParams() *PostVoucherPathParams {
	return r.pathParams
}

func (r *PostVoucherRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostVoucherRequest) Method() string {
	return r.method
}

func (r PostVoucherRequest) NewPostVoucherRequestBody() PostVoucherRequestBody {
	return PostVoucherRequestBody{}
}

type PostVoucherRequestBody struct {
	Voucher Voucher
}

func (r *PostVoucherRequest) RequestBody() *PostVoucherRequestBody {
	return &r.requestBody
}

func (r *PostVoucherRequest) SetRequestBody(body PostVoucherRequestBody) {
	r.requestBody = body
}

func (r *PostVoucherRequest) NewResponseBody() *PostVoucherResponseBody {
	return &PostVoucherResponseBody{}
}

type PostVoucherResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Voucher         Voucher
}

func (r *PostVoucherRequest) URL() url.URL {
	return r.client.GetEndpointURL("/vouchers", r.PathParams())
}

func (r *PostVoucherRequest) Do(ctx context.Context) (PostVoucherResponseBody, error) {
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
