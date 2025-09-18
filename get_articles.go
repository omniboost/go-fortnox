package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetArticlesRequest() GetArticlesRequest {
	r := GetArticlesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetArticlesQueryParams()
	r.pathParams = r.NewGetArticlesPathParams()
	r.requestBody = r.NewGetArticlesRequestBody()
	return r
}

type GetArticlesRequest struct {
	client      *Client
	queryParams *GetArticlesQueryParams
	pathParams  *GetArticlesPathParams
	method      string
	headers     http.Header
	requestBody GetArticlesRequestBody
}

func (r GetArticlesRequest) NewGetArticlesQueryParams() *GetArticlesQueryParams {
	return &GetArticlesQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetArticlesQueryParams struct {
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

func (p GetArticlesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetArticlesRequest) QueryParams() *GetArticlesQueryParams {
	return r.queryParams
}

func (r GetArticlesRequest) NewGetArticlesPathParams() *GetArticlesPathParams {
	return &GetArticlesPathParams{}
}

type GetArticlesPathParams struct {
}

func (p *GetArticlesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetArticlesRequest) PathParams() *GetArticlesPathParams {
	return r.pathParams
}

func (r *GetArticlesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetArticlesRequest) Method() string {
	return r.method
}

func (r GetArticlesRequest) NewGetArticlesRequestBody() GetArticlesRequestBody {
	return GetArticlesRequestBody{}
}

type GetArticlesRequestBody struct{}

func (r *GetArticlesRequest) RequestBody() *GetArticlesRequestBody {
	return &r.requestBody
}

func (r *GetArticlesRequest) SetRequestBody(body GetArticlesRequestBody) {
	r.requestBody = body
}

func (r *GetArticlesRequest) NewResponseBody() *GetArticlesResponseBody {
	return &GetArticlesResponseBody{}
}

type GetArticlesResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Articles        ArticleRows `json:"Articles"`
}

func (r *GetArticlesRequest) URL() url.URL {
	return r.client.GetEndpointURL("/articles", r.PathParams())
}

func (r *GetArticlesRequest) Do() (GetArticlesResponseBody, error) {
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
