package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetArticleByNumberRequest() GetArticleByNumberRequest {
	r := GetArticleByNumberRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetArticleByNumberQueryParams()
	r.pathParams = r.NewGetArticleByNumberPathParams()
	r.requestBody = r.NewGetArticleByNumberRequestBody()
	return r
}

type GetArticleByNumberRequest struct {
	client      *Client
	queryParams *GetArticleByNumberQueryParams
	pathParams  *GetArticleByNumberPathParams
	method      string
	headers     http.Header
	requestBody GetArticleByNumberRequestBody
}

func (r GetArticleByNumberRequest) NewGetArticleByNumberQueryParams() *GetArticleByNumberQueryParams {
	return &GetArticleByNumberQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetArticleByNumberQueryParams struct {
}

func (p GetArticleByNumberQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetArticleByNumberRequest) QueryParams() *GetArticleByNumberQueryParams {
	return r.queryParams
}

func (r GetArticleByNumberRequest) NewGetArticleByNumberPathParams() *GetArticleByNumberPathParams {
	return &GetArticleByNumberPathParams{}
}

type GetArticleByNumberPathParams struct {
	ArticleNumber string `schema:"articlenumber"`
}

func (p *GetArticleByNumberPathParams) Params() map[string]string {
	return map[string]string{
		"articlenumber": p.ArticleNumber,
	}
}

func (r *GetArticleByNumberRequest) PathParams() *GetArticleByNumberPathParams {
	return r.pathParams
}

func (r *GetArticleByNumberRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetArticleByNumberRequest) Method() string {
	return r.method
}

func (r GetArticleByNumberRequest) NewGetArticleByNumberRequestBody() GetArticleByNumberRequestBody {
	return GetArticleByNumberRequestBody{}
}

type GetArticleByNumberRequestBody struct{}

func (r *GetArticleByNumberRequest) RequestBody() *GetArticleByNumberRequestBody {
	return &r.requestBody
}

func (r *GetArticleByNumberRequest) SetRequestBody(body GetArticleByNumberRequestBody) {
	r.requestBody = body
}

func (r *GetArticleByNumberRequest) NewResponseBody() *GetArticleByNumberResponseBody {
	return &GetArticleByNumberResponseBody{}
}

type GetArticleByNumberResponseBody struct {
	Article Article `json:"Article"`
}

func (r *GetArticleByNumberRequest) URL() url.URL {
	return r.client.GetEndpointURL("/articles/{{.articlenumber}}", r.PathParams())
}

func (r *GetArticleByNumberRequest) Do() (GetArticleByNumberResponseBody, error) {
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
