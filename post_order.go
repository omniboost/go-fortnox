package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewPostOrderRequest() PostOrderRequest {
	r := PostOrderRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewPostOrderQueryParams()
	r.pathParams = r.NewPostOrderPathParams()
	r.requestBody = r.NewPostOrderRequestBody()
	return r
}

type PostOrderRequest struct {
	client      *Client
	queryParams *PostOrderQueryParams
	pathParams  *PostOrderPathParams
	method      string
	headers     http.Header
	requestBody PostOrderRequestBody
}

func (r PostOrderRequest) NewPostOrderQueryParams() *PostOrderQueryParams {
	return &PostOrderQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type PostOrderQueryParams struct{}

func (p PostOrderQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostOrderRequest) QueryParams() *PostOrderQueryParams {
	return r.queryParams
}

func (r PostOrderRequest) NewPostOrderPathParams() *PostOrderPathParams {
	return &PostOrderPathParams{}
}

type PostOrderPathParams struct {
}

func (p *PostOrderPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostOrderRequest) PathParams() *PostOrderPathParams {
	return r.pathParams
}

func (r *PostOrderRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostOrderRequest) Method() string {
	return r.method
}

func (r PostOrderRequest) NewPostOrderRequestBody() PostOrderRequestBody {
	return PostOrderRequestBody{}
}

type PostOrderRequestBody struct {
	Order Order `json:"Order"`
}

func (r *PostOrderRequest) RequestBody() *PostOrderRequestBody {
	return &r.requestBody
}

func (r *PostOrderRequest) SetRequestBody(body PostOrderRequestBody) {
	r.requestBody = body
}

func (r *PostOrderRequest) NewResponseBody() *PostOrderResponseBody {
	return &PostOrderResponseBody{}
}

type PostOrderResponseBody struct {
	Order Order `json:"Order"`
}

func (r *PostOrderRequest) URL() url.URL {
	return r.client.GetEndpointURL("/orders", r.PathParams())
}

func (r *PostOrderRequest) Do() (PostOrderResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
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
