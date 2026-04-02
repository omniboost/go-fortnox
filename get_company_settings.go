package fortnox

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-fortnox/utils"
)

func (c *Client) NewGetCompanySettingsRequest() GetCompanySettingsRequest {
	r := GetCompanySettingsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetCompanySettingsQueryParams()
	r.pathParams = r.NewGetCompanySettingsPathParams()
	r.requestBody = r.NewGetCompanySettingsRequestBody()
	return r
}

type GetCompanySettingsRequest struct {
	client      *Client
	queryParams *GetCompanySettingsQueryParams
	pathParams  *GetCompanySettingsPathParams
	method      string
	headers     http.Header
	requestBody GetCompanySettingsRequestBody
}

func (r GetCompanySettingsRequest) NewGetCompanySettingsQueryParams() *GetCompanySettingsQueryParams {
	return &GetCompanySettingsQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type GetCompanySettingsQueryParams struct {
}

func (p GetCompanySettingsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCompanySettingsRequest) QueryParams() *GetCompanySettingsQueryParams {
	return r.queryParams
}

func (r GetCompanySettingsRequest) NewGetCompanySettingsPathParams() *GetCompanySettingsPathParams {
	return &GetCompanySettingsPathParams{}
}

type GetCompanySettingsPathParams struct {
}

func (p *GetCompanySettingsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCompanySettingsRequest) PathParams() *GetCompanySettingsPathParams {
	return r.pathParams
}

func (r *GetCompanySettingsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCompanySettingsRequest) Method() string {
	return r.method
}

func (r GetCompanySettingsRequest) NewGetCompanySettingsRequestBody() GetCompanySettingsRequestBody {
	return GetCompanySettingsRequestBody{}
}

type GetCompanySettingsRequestBody struct{}

func (r *GetCompanySettingsRequest) RequestBody() *GetCompanySettingsRequestBody {
	return &r.requestBody
}

func (r *GetCompanySettingsRequest) SetRequestBody(body GetCompanySettingsRequestBody) {
	r.requestBody = body
}

func (r *GetCompanySettingsRequest) NewResponseBody() *GetCompanySettingsResponseBody {
	return &GetCompanySettingsResponseBody{}
}

type GetCompanySettingsResponseBody struct {
	CompanySettings CompanySettings
}

func (r *GetCompanySettingsRequest) URL() url.URL {
	return r.client.GetEndpointURL("/settings/company", r.PathParams())
}

func (r *GetCompanySettingsRequest) Do(ctx context.Context) (GetCompanySettingsResponseBody, error) {
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

type CompanySettings struct {
	Address            string `json:"Address"`
	BG                 string `json:"BG"`
	BIC                string `json:"BIC"`
	BranchCode         string `json:"BranchCode"`
	City               string `json:"City"`
	ContactFirstName   string `json:"ContactFirstName"`
	ContactLastName    string `json:"ContactLastName"`
	Country            string `json:"Country"`
	CountryCode        string `json:"CountryCode"`
	DatabaseNumber     int    `json:"DatabaseNumber"`
	Domicile           string `json:"Domicile"`
	Email              string `json:"Email"`
	Fax                string `json:"Fax"`
	IBAN               string `json:"IBAN"`
	Name               string `json:"Name"`
	OrganizationNumber string `json:"OrganizationNumber"`
	PG                 string `json:"PG"`
	Phone1             string `json:"Phone1"`
	Phone2             string `json:"Phone2"`
	TaxEnabled         bool   `json:"TaxEnabled"`
	VATNumber          string `json:"VATNumber"`
	VisitAddress       string `json:"VisitAddress"`
	VisitCity          string `json:"VisitCity"`
	VisitCountry       string `json:"VisitCountry"`
	VisitCountryCode   string `json:"VisitCountryCode"`
	VisitName          string `json:"VisitName"`
	VisitZipCode       string `json:"VisitZipCode"`
	WWW                string `json:"WWW"`
	ZipCode            string `json:"ZipCode"`
}
