package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetOrderByNumberRequest() GetOrderByNumberRequest {
	r := GetOrderByNumberRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetOrderByNumberQueryParams()
	r.pathParams = r.NewGetOrderByNumberPathParams()
	r.requestBody = r.NewGetOrderByNumberRequestBody()
	return r
}

type GetOrderByNumberRequest struct {
	client      *Client
	queryParams *GetOrderByNumberQueryParams
	pathParams  *GetOrderByNumberPathParams
	method      string
	headers     http.Header
	requestBody GetOrderByNumberRequestBody
}

func (r GetOrderByNumberRequest) NewGetOrderByNumberQueryParams() *GetOrderByNumberQueryParams {
	return &GetOrderByNumberQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetOrderByNumberQueryParams struct {
}

func (p GetOrderByNumberQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetOrderByNumberRequest) QueryParams() *GetOrderByNumberQueryParams {
	return r.queryParams
}

func (r GetOrderByNumberRequest) NewGetOrderByNumberPathParams() *GetOrderByNumberPathParams {
	return &GetOrderByNumberPathParams{}
}

type GetOrderByNumberPathParams struct {
	DocumentNumber string `schema:"documentnumber"`
}

func (p *GetOrderByNumberPathParams) Params() map[string]string {
	return map[string]string{
		"documentnumber": p.DocumentNumber,
	}
}

func (r *GetOrderByNumberRequest) PathParams() *GetOrderByNumberPathParams {
	return r.pathParams
}

func (r *GetOrderByNumberRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetOrderByNumberRequest) Method() string {
	return r.method
}

func (r GetOrderByNumberRequest) NewGetOrderByNumberRequestBody() GetOrderByNumberRequestBody {
	return GetOrderByNumberRequestBody{}
}

type GetOrderByNumberRequestBody struct{}

func (r *GetOrderByNumberRequest) RequestBody() *GetOrderByNumberRequestBody {
	return &r.requestBody
}

func (r *GetOrderByNumberRequest) SetRequestBody(body GetOrderByNumberRequestBody) {
	r.requestBody = body
}

func (r *GetOrderByNumberRequest) NewResponseBody() *GetOrderByNumberResponseBody {
	return &GetOrderByNumberResponseBody{}
}

type GetOrderByNumberResponseBody struct {
	Order Order `json:"Order"`
}

func (r *GetOrderByNumberRequest) URL() url.URL {
	return r.client.GetEndpointURL("/orders/{{.documentnumber}}", r.PathParams())
}

func (r *GetOrderByNumberRequest) Do() (GetOrderByNumberResponseBody, error) {
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
