// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// PostNamespacesNsEstimatorsNameValuesPowerconsumption request with any body
	PostNamespacesNsEstimatorsNameValuesPowerconsumptionWithBody(ctx context.Context, ns string, name string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostNamespacesNsEstimatorsNameValuesPowerconsumption(ctx context.Context, ns string, name string, body PostNamespacesNsEstimatorsNameValuesPowerconsumptionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostNamespacesNsEstimatorsNameValuesPowerconsumptionWithBody(ctx context.Context, ns string, name string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostNamespacesNsEstimatorsNameValuesPowerconsumptionRequestWithBody(c.Server, ns, name, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostNamespacesNsEstimatorsNameValuesPowerconsumption(ctx context.Context, ns string, name string, body PostNamespacesNsEstimatorsNameValuesPowerconsumptionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostNamespacesNsEstimatorsNameValuesPowerconsumptionRequest(c.Server, ns, name, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostNamespacesNsEstimatorsNameValuesPowerconsumptionRequest calls the generic PostNamespacesNsEstimatorsNameValuesPowerconsumption builder with application/json body
func NewPostNamespacesNsEstimatorsNameValuesPowerconsumptionRequest(server string, ns string, name string, body PostNamespacesNsEstimatorsNameValuesPowerconsumptionJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostNamespacesNsEstimatorsNameValuesPowerconsumptionRequestWithBody(server, ns, name, "application/json", bodyReader)
}

// NewPostNamespacesNsEstimatorsNameValuesPowerconsumptionRequestWithBody generates requests for PostNamespacesNsEstimatorsNameValuesPowerconsumption with any type of body
func NewPostNamespacesNsEstimatorsNameValuesPowerconsumptionRequestWithBody(server string, ns string, name string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "ns", runtime.ParamLocationPath, ns)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "name", runtime.ParamLocationPath, name)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/namespaces/%s/estimators/%s/values/powerconsumption", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// PostNamespacesNsEstimatorsNameValuesPowerconsumption request with any body
	PostNamespacesNsEstimatorsNameValuesPowerconsumptionWithBodyWithResponse(ctx context.Context, ns string, name string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse, error)

	PostNamespacesNsEstimatorsNameValuesPowerconsumptionWithResponse(ctx context.Context, ns string, name string, body PostNamespacesNsEstimatorsNameValuesPowerconsumptionJSONRequestBody, reqEditors ...RequestEditorFn) (*PostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse, error)
}

type PostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PowerConsumption
}

// Status returns HTTPResponse.Status
func (r PostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostNamespacesNsEstimatorsNameValuesPowerconsumptionWithBodyWithResponse request with arbitrary body returning *PostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse
func (c *ClientWithResponses) PostNamespacesNsEstimatorsNameValuesPowerconsumptionWithBodyWithResponse(ctx context.Context, ns string, name string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse, error) {
	rsp, err := c.PostNamespacesNsEstimatorsNameValuesPowerconsumptionWithBody(ctx, ns, name, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse(rsp)
}

func (c *ClientWithResponses) PostNamespacesNsEstimatorsNameValuesPowerconsumptionWithResponse(ctx context.Context, ns string, name string, body PostNamespacesNsEstimatorsNameValuesPowerconsumptionJSONRequestBody, reqEditors ...RequestEditorFn) (*PostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse, error) {
	rsp, err := c.PostNamespacesNsEstimatorsNameValuesPowerconsumption(ctx, ns, name, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse(rsp)
}

// ParsePostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse parses an HTTP response from a PostNamespacesNsEstimatorsNameValuesPowerconsumptionWithResponse call
func ParsePostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse(rsp *http.Response) (*PostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostNamespacesNsEstimatorsNameValuesPowerconsumptionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PowerConsumption
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}