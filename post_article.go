package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewPostArticleRequest() PostArticleRequest {
	r := PostArticleRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewPostArticleQueryParams()
	r.pathParams = r.NewPostArticlePathParams()
	r.requestBody = r.NewPostArticleRequestBody()
	return r
}

type PostArticleRequest struct {
	client      *Client
	queryParams *PostArticleQueryParams
	pathParams  *PostArticlePathParams
	method      string
	headers     http.Header
	requestBody PostArticleRequestBody
}

func (r PostArticleRequest) NewPostArticleQueryParams() *PostArticleQueryParams {
	return &PostArticleQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type PostArticleQueryParams struct {
	Filter                    string `schema:"filter,omitempty"`
	ArticleNumber             string `schema:"articlenumber,omitempty"`
	Description               string `schema:"description,omitempty"`
	EAN                       string `schema:"ean,omitempty"`
	SupplierNumber            string `schema:"suppliernumber,omitempty"`
	Manufacturer              string `schema:"manufacturer,omitempty"`
	ManufacturerArticleNumber string `schema:"manufacturerarticlenumber,omitempty"`
	Webshop                   string `schema:"webshop,omitempty"`
	LastModified              string `schema:"lastmodified,omitempty"`
	SortBy                    string `schema:"sortby,omitempty"`
}

func (p PostArticleQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostArticleRequest) QueryParams() *PostArticleQueryParams {
	return r.queryParams
}

func (r PostArticleRequest) NewPostArticlePathParams() *PostArticlePathParams {
	return &PostArticlePathParams{}
}

type PostArticlePathParams struct {
}

func (p *PostArticlePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostArticleRequest) PathParams() *PostArticlePathParams {
	return r.pathParams
}

func (r *PostArticleRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostArticleRequest) Method() string {
	return r.method
}

func (r PostArticleRequest) NewPostArticleRequestBody() PostArticleRequestBody {
	return PostArticleRequestBody{}
}

type PostArticleRequestBody struct {
	Article Article `json:"Article"`
}

func (r *PostArticleRequest) RequestBody() *PostArticleRequestBody {
	return &r.requestBody
}

func (r *PostArticleRequest) SetRequestBody(body PostArticleRequestBody) {
	r.requestBody = body
}

func (r *PostArticleRequest) NewResponseBody() *PostArticleResponseBody {
	return &PostArticleResponseBody{}
}

type PostArticleResponseBody struct {
	Article Article `json:"Article"`
}

func (r *PostArticleRequest) URL() url.URL {
	return r.client.GetEndpointURL("/articles", r.PathParams())
}

func (r *PostArticleRequest) Do() (PostArticleResponseBody, error) {
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
