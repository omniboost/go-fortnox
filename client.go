package fortnox

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/omniboost/go-fortnox/utils"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-fortnox/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "api.fortnox.se",
		Path:   "/3",
	}
	requestTimestamps = make(map[string]*timestamps)
	requestsPerSecond = 4
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetClientSecret("")
	client.SetAccessToken("")
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	clientSecret string
	accessToken  string

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c Client) ClientSecret() string {
	return c.clientSecret
}

func (c *Client) SetClientSecret(clientSecret string) {
	c.clientSecret = clientSecret
}

func (c Client) AccessToken() string {
	return c.accessToken
}

func (c *Client) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
	if requestTimestamps[accessToken] == nil {
		requestTimestamps[accessToken] = &timestamps{}
	}
}

func (c Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) GetEndpointURL(path string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()
	clientURL.Path = clientURL.Path + path

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, method string, URL url.URL, body interface{}) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	err = utils.AddURLValuesToRequest(values, req, true)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// set other headers
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())

	if c.AccessToken() != "" && c.ClientSecret() != "" {
		req.Header.Add("Client-Secret", c.ClientSecret())
		req.Header.Add("Access-Token", c.AccessToken())
	}

	return req, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	c.SleepUntilRequestRate()

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	// register timestamp after request has a response
	c.RegisterRequestTimestamp(time.Now())

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if responseBody == nil {
		return httpResp, err
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	if responseBody == nil {
		return httpResp, nil
	}

	err = c.Unmarshal(httpResp.Body, &responseBody)
	if err != nil {
		return httpResp, err
	}

	// if len(errorResponse.Messages) > 0 {
	// 	return httpResp, errorResponse
	// }

	return httpResp, nil
}

func (c *Client) RegisterRequestTimestamp(t time.Time) {
	if len(*requestTimestamps[c.accessToken]) >= requestsPerSecond {
		ts := (*requestTimestamps[c.accessToken])[1:requestsPerSecond]
		requestTimestamps[c.accessToken] = &ts
	}
	ts := append(*requestTimestamps[c.accessToken], t)
	requestTimestamps[c.accessToken] = &ts
}

func (c *Client) SleepUntilRequestRate() {
	// Requestrate is 4r/1s

	// if there are less then 4 registered requests: execute the request
	// immediately
	if len(*requestTimestamps[c.accessToken]) < (requestsPerSecond - 1) {
		return
	}

	// is the first item within 1 second? If it's > 1 second the request can be
	// executed imediately
	diff := time.Now().Sub((*requestTimestamps[c.accessToken])[0])
	if diff >= time.Second {
		return
	}

	// Sleep for the time it takes for the first item to be > 1 second old
	// + 1ms to be sure :)
	time.Sleep(time.Second - diff + (1 * time.Millisecond))
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	wg := sync.WaitGroup{}
	wg.Add(len(vv))
	errs := []error{}
	writers := make([]io.Writer, len(vv))

	for i, v := range vv {
		pr, pw := io.Pipe()
		writers[i] = pw

		go func(i int, v interface{}, pr *io.PipeReader, pw *io.PipeWriter) {
			dec := json.NewDecoder(pr)
			if c.disallowUnknownFields {
				dec.DisallowUnknownFields()
			}

			err := dec.Decode(v)
			if err != nil {
				errs = append(errs, err)
			}

			// mark routine as done
			wg.Done()

			// Drain reader
			io.Copy(ioutil.Discard, pr)

			// close reader
			// pr.CloseWithError(err)
			pr.Close()
		}(i, v, pr, pw)
	}

	// copy the data in a multiwriter
	mw := io.MultiWriter(writers...)
	_, err := io.Copy(mw, r)
	if err != nil {
		return err
	}

	wg.Wait()
	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}
	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r, data)
	if err != nil {
		errorResponse.ErrorInformation = err
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert json to struct
	dest := struct {
		ErrorInformation ErrorInformation
	}{}
	err = json.Unmarshal(data, &dest)
	if err != nil {
		errorResponse.ErrorInformation = err
		return errorResponse
	}
	errorResponse.ErrorInformation = dest.ErrorInformation

	return errorResponse
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	ErrorInformation error `json:"ErrorInformation"`
}

func (r ErrorResponse) Error() string {
	if r.ErrorInformation == nil {
		return ""
	}
	return r.ErrorInformation.Error()
}

type ErrorInformation struct {
	Err     int    `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e ErrorInformation) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func checkContentType(response *http.Response, data []byte) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]

	// if the response is not JSON, fall back to raw data
	if contentType != mediaType {
		msg := strings.TrimSpace(string(data))
		if msg == "" {
			msg = response.Status
		}
		ei := ErrorInformation{
			Message: msg,
			Code:    response.StatusCode,
		}
		return ei
	}

	return nil
}

type PathParams interface {
	Params() map[string]string
}

type timestamps []time.Time
