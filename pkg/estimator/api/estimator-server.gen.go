// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Send a power consumption estimate request.
	// (POST /namespaces/{namespace}/estimators/{estimator}/resources/powerconsumption)
	PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption(w http.ResponseWriter, r *http.Request, namespace string, estimator string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption operation middleware
func (siw *ServerInterfaceWrapper) PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "namespace" -------------
	var namespace string

	err = runtime.BindStyledParameterWithLocation("simple", false, "namespace", runtime.ParamLocationPath, chi.URLParam(r, "namespace"), &namespace)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "namespace", Err: err})
		return
	}

	// ------------- Path parameter "estimator" -------------
	var estimator string

	err = runtime.BindStyledParameterWithLocation("simple", false, "estimator", runtime.ParamLocationPath, chi.URLParam(r, "estimator"), &estimator)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "estimator", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ApiKeyAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption(w, r, namespace, estimator)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/namespaces/{namespace}/estimators/{estimator}/resources/powerconsumption", wrapper.PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption)
	})

	return r
}

type PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionRequestObject struct {
	Namespace string `json:"namespace"`
	Estimator string `json:"estimator"`
	Body      *PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionJSONRequestBody
}

type PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionResponseObject interface {
	VisitPostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionResponse(w http.ResponseWriter) error
}

type PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption200JSONResponse PowerConsumption

func (response PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption200JSONResponse) VisitPostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption400Response struct {
}

func (response PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption400Response) VisitPostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption401Response struct {
}

func (response PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption401Response) VisitPostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption404Response struct {
}

func (response PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption404Response) VisitPostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption500Response struct {
}

func (response PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption500Response) VisitPostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Send a power consumption estimate request.
	// (POST /namespaces/{namespace}/estimators/{estimator}/resources/powerconsumption)
	PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption(ctx context.Context, request PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionRequestObject) (PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionResponseObject, error)
}

type StrictHandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption operation middleware
func (sh *strictHandler) PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption(w http.ResponseWriter, r *http.Request, namespace string, estimator string) {
	var request PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionRequestObject

	request.Namespace = namespace
	request.Estimator = estimator

	var body PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption(ctx, request.(PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumption")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionResponseObject); ok {
		if err := validResponse.VisitPostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}
