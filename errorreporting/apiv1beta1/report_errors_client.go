// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package errorreporting

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	errorreportingpb "cloud.google.com/go/errorreporting/apiv1beta1/errorreportingpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newReportErrorsClientHook clientHook

// ReportErrorsCallOptions contains the retry settings for each method of ReportErrorsClient.
type ReportErrorsCallOptions struct {
	ReportErrorEvent []gax.CallOption
}

func defaultReportErrorsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("clouderrorreporting.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("clouderrorreporting.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://clouderrorreporting.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultReportErrorsCallOptions() *ReportErrorsCallOptions {
	return &ReportErrorsCallOptions{
		ReportErrorEvent: []gax.CallOption{},
	}
}

func defaultReportErrorsRESTCallOptions() *ReportErrorsCallOptions {
	return &ReportErrorsCallOptions{
		ReportErrorEvent: []gax.CallOption{},
	}
}

// internalReportErrorsClient is an interface that defines the methods available from Error Reporting API.
type internalReportErrorsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ReportErrorEvent(context.Context, *errorreportingpb.ReportErrorEventRequest, ...gax.CallOption) (*errorreportingpb.ReportErrorEventResponse, error)
}

// ReportErrorsClient is a client for interacting with Error Reporting API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// An API for reporting error events.
type ReportErrorsClient struct {
	// The internal transport-dependent client.
	internalClient internalReportErrorsClient

	// The call options for this service.
	CallOptions *ReportErrorsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ReportErrorsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ReportErrorsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ReportErrorsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ReportErrorEvent report an individual error event and record the event to a log.
//
// This endpoint accepts either an OAuth token,
// or an API key (at https://support.google.com/cloud/answer/6158862)
// for authentication. To use an API key, append it to the URL as the value of
// a key parameter. For example:
//
// POST https://clouderrorreporting.googleapis.com/v1beta1/{projectName}/events:report?key=123ABC456
//
// Note: Error Reporting (at /error-reporting) is a global service built
// on Cloud Logging and doesn’t analyze logs stored
// in regional log buckets or logs routed to other Google Cloud projects.
//
// For more information, see
// Using Error Reporting with regionalized
// logs (at /error-reporting/docs/regionalization).
func (c *ReportErrorsClient) ReportErrorEvent(ctx context.Context, req *errorreportingpb.ReportErrorEventRequest, opts ...gax.CallOption) (*errorreportingpb.ReportErrorEventResponse, error) {
	return c.internalClient.ReportErrorEvent(ctx, req, opts...)
}

// reportErrorsGRPCClient is a client for interacting with Error Reporting API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type reportErrorsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ReportErrorsClient
	CallOptions **ReportErrorsCallOptions

	// The gRPC API client.
	reportErrorsClient errorreportingpb.ReportErrorsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewReportErrorsClient creates a new report errors service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// An API for reporting error events.
func NewReportErrorsClient(ctx context.Context, opts ...option.ClientOption) (*ReportErrorsClient, error) {
	clientOpts := defaultReportErrorsGRPCClientOptions()
	if newReportErrorsClientHook != nil {
		hookOpts, err := newReportErrorsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ReportErrorsClient{CallOptions: defaultReportErrorsCallOptions()}

	c := &reportErrorsGRPCClient{
		connPool:           connPool,
		disableDeadlines:   disableDeadlines,
		reportErrorsClient: errorreportingpb.NewReportErrorsServiceClient(connPool),
		CallOptions:        &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *reportErrorsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *reportErrorsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *reportErrorsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type reportErrorsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing ReportErrorsClient
	CallOptions **ReportErrorsCallOptions
}

// NewReportErrorsRESTClient creates a new report errors service rest client.
//
// An API for reporting error events.
func NewReportErrorsRESTClient(ctx context.Context, opts ...option.ClientOption) (*ReportErrorsClient, error) {
	clientOpts := append(defaultReportErrorsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultReportErrorsRESTCallOptions()
	c := &reportErrorsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &ReportErrorsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultReportErrorsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://clouderrorreporting.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://clouderrorreporting.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://clouderrorreporting.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *reportErrorsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *reportErrorsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *reportErrorsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *reportErrorsGRPCClient) ReportErrorEvent(ctx context.Context, req *errorreportingpb.ReportErrorEventRequest, opts ...gax.CallOption) (*errorreportingpb.ReportErrorEventResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_name", url.QueryEscape(req.GetProjectName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ReportErrorEvent[0:len((*c.CallOptions).ReportErrorEvent):len((*c.CallOptions).ReportErrorEvent)], opts...)
	var resp *errorreportingpb.ReportErrorEventResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.reportErrorsClient.ReportErrorEvent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ReportErrorEvent report an individual error event and record the event to a log.
//
// This endpoint accepts either an OAuth token,
// or an API key (at https://support.google.com/cloud/answer/6158862)
// for authentication. To use an API key, append it to the URL as the value of
// a key parameter. For example:
//
// POST https://clouderrorreporting.googleapis.com/v1beta1/{projectName}/events:report?key=123ABC456
//
// Note: Error Reporting (at /error-reporting) is a global service built
// on Cloud Logging and doesn’t analyze logs stored
// in regional log buckets or logs routed to other Google Cloud projects.
//
// For more information, see
// Using Error Reporting with regionalized
// logs (at /error-reporting/docs/regionalization).
func (c *reportErrorsRESTClient) ReportErrorEvent(ctx context.Context, req *errorreportingpb.ReportErrorEventRequest, opts ...gax.CallOption) (*errorreportingpb.ReportErrorEventResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetEvent()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/events:report", req.GetProjectName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_name", url.QueryEscape(req.GetProjectName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).ReportErrorEvent[0:len((*c.CallOptions).ReportErrorEvent):len((*c.CallOptions).ReportErrorEvent)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &errorreportingpb.ReportErrorEventResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
