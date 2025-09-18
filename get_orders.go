package fortnox

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetOrdersRequest() GetOrdersRequest {
	r := GetOrdersRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetOrdersQueryParams()
	r.pathParams = r.NewGetOrdersPathParams()
	r.requestBody = r.NewGetOrdersRequestBody()
	return r
}

type GetOrdersRequest struct {
	client      *Client
	queryParams *GetOrdersQueryParams
	pathParams  *GetOrdersPathParams
	method      string
	headers     http.Header
	requestBody GetOrdersRequestBody
}

func (r GetOrdersRequest) NewGetOrdersQueryParams() *GetOrdersQueryParams {
	return &GetOrdersQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetOrdersQueryParams struct {
	Filter                    string `schema:"filter,omitempty"`
	CustomerName              string `schema:"customername,omitempty"`
	CustomerNumber            string `schema:"customernumber,omitempty"`
	Label                     string `schema:"label,omitempty"`
	DocumentNumber            string `schema:"documentnumber,omitempty"`
	ExternalInvoiceReference1 string `schema:"externalinvoicereference1,omitempty"`
	ExternalInvoiceReference2 string `schema:"externalinvoicereference2,omitempty"`
	FromDate                  string `schema:"fromdate,omitempty"`
	ToDate                    string `schema:"todate,omitempty"`
	CostCenter                string `schema:"costcenter,omitempty"`
	Project                   string `schema:"project,omitempty"`
	Sent                      bool   `schema:"sent,omitempty"`
	NotCompleted              bool   `schema:"notcompleted,omitempty"`
	OurReference              string `schema:"ourreference,omitempty"`
	YourReference             string `schema:"yourreference,omitempty"`
	LastModified              string `schema:"lastmodified,omitempty"`
	OrderType                 string `schema:"ordertype,omitempty"`
	SortBy                    string `schema:"sortby,omitempty"`
}

func (p GetOrdersQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetOrdersRequest) QueryParams() *GetOrdersQueryParams {
	return r.queryParams
}

func (r GetOrdersRequest) NewGetOrdersPathParams() *GetOrdersPathParams {
	return &GetOrdersPathParams{}
}

type GetOrdersPathParams struct {
}

func (p *GetOrdersPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetOrdersRequest) PathParams() *GetOrdersPathParams {
	return r.pathParams
}

func (r *GetOrdersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetOrdersRequest) Method() string {
	return r.method
}

func (r GetOrdersRequest) NewGetOrdersRequestBody() GetOrdersRequestBody {
	return GetOrdersRequestBody{}
}

type GetOrdersRequestBody struct{}

func (r *GetOrdersRequest) RequestBody() *GetOrdersRequestBody {
	return &r.requestBody
}

func (r *GetOrdersRequest) SetRequestBody(body GetOrdersRequestBody) {
	r.requestBody = body
}

func (r *GetOrdersRequest) NewResponseBody() *GetOrdersResponseBody {
	return &GetOrdersResponseBody{}
}

type GetOrdersResponseBody struct {
	MetaInformation `json:"MetaInformation"`
	Orders          OrderRows `json:"Orders"`
}

func (r *GetOrdersRequest) URL() url.URL {
	return r.client.GetEndpointURL("/orders", r.PathParams())
}

func (r *GetOrdersRequest) Do() (GetOrdersResponseBody, error) {
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
