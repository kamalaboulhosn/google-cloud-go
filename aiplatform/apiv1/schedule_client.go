// Copyright 2024 Google LLC
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

package aiplatform

import (
	"context"
	"fmt"
	"math"
	"net/url"

	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var newScheduleClientHook clientHook

// ScheduleCallOptions contains the retry settings for each method of ScheduleClient.
type ScheduleCallOptions struct {
	CreateSchedule     []gax.CallOption
	DeleteSchedule     []gax.CallOption
	GetSchedule        []gax.CallOption
	ListSchedules      []gax.CallOption
	PauseSchedule      []gax.CallOption
	ResumeSchedule     []gax.CallOption
	UpdateSchedule     []gax.CallOption
	GetLocation        []gax.CallOption
	ListLocations      []gax.CallOption
	GetIamPolicy       []gax.CallOption
	SetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
	CancelOperation    []gax.CallOption
	DeleteOperation    []gax.CallOption
	GetOperation       []gax.CallOption
	ListOperations     []gax.CallOption
	WaitOperation      []gax.CallOption
}

func defaultScheduleGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("aiplatform.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultScheduleCallOptions() *ScheduleCallOptions {
	return &ScheduleCallOptions{
		CreateSchedule:     []gax.CallOption{},
		DeleteSchedule:     []gax.CallOption{},
		GetSchedule:        []gax.CallOption{},
		ListSchedules:      []gax.CallOption{},
		PauseSchedule:      []gax.CallOption{},
		ResumeSchedule:     []gax.CallOption{},
		UpdateSchedule:     []gax.CallOption{},
		GetLocation:        []gax.CallOption{},
		ListLocations:      []gax.CallOption{},
		GetIamPolicy:       []gax.CallOption{},
		SetIamPolicy:       []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{},
		CancelOperation:    []gax.CallOption{},
		DeleteOperation:    []gax.CallOption{},
		GetOperation:       []gax.CallOption{},
		ListOperations:     []gax.CallOption{},
		WaitOperation:      []gax.CallOption{},
	}
}

// internalScheduleClient is an interface that defines the methods available from Vertex AI API.
type internalScheduleClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateSchedule(context.Context, *aiplatformpb.CreateScheduleRequest, ...gax.CallOption) (*aiplatformpb.Schedule, error)
	DeleteSchedule(context.Context, *aiplatformpb.DeleteScheduleRequest, ...gax.CallOption) (*DeleteScheduleOperation, error)
	DeleteScheduleOperation(name string) *DeleteScheduleOperation
	GetSchedule(context.Context, *aiplatformpb.GetScheduleRequest, ...gax.CallOption) (*aiplatformpb.Schedule, error)
	ListSchedules(context.Context, *aiplatformpb.ListSchedulesRequest, ...gax.CallOption) *ScheduleIterator
	PauseSchedule(context.Context, *aiplatformpb.PauseScheduleRequest, ...gax.CallOption) error
	ResumeSchedule(context.Context, *aiplatformpb.ResumeScheduleRequest, ...gax.CallOption) error
	UpdateSchedule(context.Context, *aiplatformpb.UpdateScheduleRequest, ...gax.CallOption) (*aiplatformpb.Schedule, error)
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
	WaitOperation(context.Context, *longrunningpb.WaitOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// ScheduleClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for creating and managing Vertex AI’s Schedule resources to
// periodically launch shceudled runs to make API calls.
type ScheduleClient struct {
	// The internal transport-dependent client.
	internalClient internalScheduleClient

	// The call options for this service.
	CallOptions *ScheduleCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ScheduleClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ScheduleClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ScheduleClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateSchedule creates a Schedule.
func (c *ScheduleClient) CreateSchedule(ctx context.Context, req *aiplatformpb.CreateScheduleRequest, opts ...gax.CallOption) (*aiplatformpb.Schedule, error) {
	return c.internalClient.CreateSchedule(ctx, req, opts...)
}

// DeleteSchedule deletes a Schedule.
func (c *ScheduleClient) DeleteSchedule(ctx context.Context, req *aiplatformpb.DeleteScheduleRequest, opts ...gax.CallOption) (*DeleteScheduleOperation, error) {
	return c.internalClient.DeleteSchedule(ctx, req, opts...)
}

// DeleteScheduleOperation returns a new DeleteScheduleOperation from a given name.
// The name must be that of a previously created DeleteScheduleOperation, possibly from a different process.
func (c *ScheduleClient) DeleteScheduleOperation(name string) *DeleteScheduleOperation {
	return c.internalClient.DeleteScheduleOperation(name)
}

// GetSchedule gets a Schedule.
func (c *ScheduleClient) GetSchedule(ctx context.Context, req *aiplatformpb.GetScheduleRequest, opts ...gax.CallOption) (*aiplatformpb.Schedule, error) {
	return c.internalClient.GetSchedule(ctx, req, opts...)
}

// ListSchedules lists Schedules in a Location.
func (c *ScheduleClient) ListSchedules(ctx context.Context, req *aiplatformpb.ListSchedulesRequest, opts ...gax.CallOption) *ScheduleIterator {
	return c.internalClient.ListSchedules(ctx, req, opts...)
}

// PauseSchedule pauses a Schedule. Will mark
// Schedule.state to ‘PAUSED’. If
// the schedule is paused, no new runs will be created. Already created runs
// will NOT be paused or canceled.
func (c *ScheduleClient) PauseSchedule(ctx context.Context, req *aiplatformpb.PauseScheduleRequest, opts ...gax.CallOption) error {
	return c.internalClient.PauseSchedule(ctx, req, opts...)
}

// ResumeSchedule resumes a paused Schedule to start scheduling new runs. Will mark
// Schedule.state to ‘ACTIVE’.
// Only paused Schedule can be resumed.
//
// When the Schedule is resumed, new runs will be scheduled starting from the
// next execution time after the current time based on the time_specification
// in the Schedule. If Schedule.catchUp is set up true, all
// missed runs will be scheduled for backfill first.
func (c *ScheduleClient) ResumeSchedule(ctx context.Context, req *aiplatformpb.ResumeScheduleRequest, opts ...gax.CallOption) error {
	return c.internalClient.ResumeSchedule(ctx, req, opts...)
}

// UpdateSchedule updates an active or paused Schedule.
//
// When the Schedule is updated, new runs will be scheduled starting from the
// updated next execution time after the update time based on the
// time_specification in the updated Schedule. All unstarted runs before the
// update time will be skipped while already created runs will NOT be paused
// or canceled.
func (c *ScheduleClient) UpdateSchedule(ctx context.Context, req *aiplatformpb.UpdateScheduleRequest, opts ...gax.CallOption) (*aiplatformpb.Schedule, error) {
	return c.internalClient.UpdateSchedule(ctx, req, opts...)
}

// GetLocation gets information about a location.
func (c *ScheduleClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *ScheduleClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. Returns an empty policy
// if the resource exists and does not have a policy set.
func (c *ScheduleClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces
// any existing policy.
//
// Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED
// errors.
func (c *ScheduleClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource. If the
// resource does not exist, this will return an empty set of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware UIs and command-line tools, not for authorization
// checking. This operation may “fail open” without warning.
func (c *ScheduleClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *ScheduleClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *ScheduleClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *ScheduleClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *ScheduleClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// WaitOperation is a utility method from google.longrunning.Operations.
func (c *ScheduleClient) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.WaitOperation(ctx, req, opts...)
}

// scheduleGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type scheduleGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ScheduleClient
	CallOptions **ScheduleCallOptions

	// The gRPC API client.
	scheduleClient aiplatformpb.ScheduleServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	iamPolicyClient iampb.IAMPolicyClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewScheduleClient creates a new schedule service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for creating and managing Vertex AI’s Schedule resources to
// periodically launch shceudled runs to make API calls.
func NewScheduleClient(ctx context.Context, opts ...option.ClientOption) (*ScheduleClient, error) {
	clientOpts := defaultScheduleGRPCClientOptions()
	if newScheduleClientHook != nil {
		hookOpts, err := newScheduleClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ScheduleClient{CallOptions: defaultScheduleCallOptions()}

	c := &scheduleGRPCClient{
		connPool:         connPool,
		scheduleClient:   aiplatformpb.NewScheduleServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
		iamPolicyClient:  iampb.NewIAMPolicyClient(connPool),
		locationsClient:  locationpb.NewLocationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *scheduleGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *scheduleGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *scheduleGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *scheduleGRPCClient) CreateSchedule(ctx context.Context, req *aiplatformpb.CreateScheduleRequest, opts ...gax.CallOption) (*aiplatformpb.Schedule, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateSchedule[0:len((*c.CallOptions).CreateSchedule):len((*c.CallOptions).CreateSchedule)], opts...)
	var resp *aiplatformpb.Schedule
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.scheduleClient.CreateSchedule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *scheduleGRPCClient) DeleteSchedule(ctx context.Context, req *aiplatformpb.DeleteScheduleRequest, opts ...gax.CallOption) (*DeleteScheduleOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteSchedule[0:len((*c.CallOptions).DeleteSchedule):len((*c.CallOptions).DeleteSchedule)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.scheduleClient.DeleteSchedule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteScheduleOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *scheduleGRPCClient) GetSchedule(ctx context.Context, req *aiplatformpb.GetScheduleRequest, opts ...gax.CallOption) (*aiplatformpb.Schedule, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetSchedule[0:len((*c.CallOptions).GetSchedule):len((*c.CallOptions).GetSchedule)], opts...)
	var resp *aiplatformpb.Schedule
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.scheduleClient.GetSchedule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *scheduleGRPCClient) ListSchedules(ctx context.Context, req *aiplatformpb.ListSchedulesRequest, opts ...gax.CallOption) *ScheduleIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListSchedules[0:len((*c.CallOptions).ListSchedules):len((*c.CallOptions).ListSchedules)], opts...)
	it := &ScheduleIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListSchedulesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.Schedule, string, error) {
		resp := &aiplatformpb.ListSchedulesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.scheduleClient.ListSchedules(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSchedules(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *scheduleGRPCClient) PauseSchedule(ctx context.Context, req *aiplatformpb.PauseScheduleRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PauseSchedule[0:len((*c.CallOptions).PauseSchedule):len((*c.CallOptions).PauseSchedule)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.scheduleClient.PauseSchedule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *scheduleGRPCClient) ResumeSchedule(ctx context.Context, req *aiplatformpb.ResumeScheduleRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ResumeSchedule[0:len((*c.CallOptions).ResumeSchedule):len((*c.CallOptions).ResumeSchedule)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.scheduleClient.ResumeSchedule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *scheduleGRPCClient) UpdateSchedule(ctx context.Context, req *aiplatformpb.UpdateScheduleRequest, opts ...gax.CallOption) (*aiplatformpb.Schedule, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "schedule.name", url.QueryEscape(req.GetSchedule().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateSchedule[0:len((*c.CallOptions).UpdateSchedule):len((*c.CallOptions).UpdateSchedule)], opts...)
	var resp *aiplatformpb.Schedule
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.scheduleClient.UpdateSchedule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *scheduleGRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.locationsClient.GetLocation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *scheduleGRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.locationsClient.ListLocations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *scheduleGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *scheduleGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *scheduleGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *scheduleGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *scheduleGRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.DeleteOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *scheduleGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *scheduleGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *scheduleGRPCClient) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).WaitOperation[0:len((*c.CallOptions).WaitOperation):len((*c.CallOptions).WaitOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.WaitOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteScheduleOperation returns a new DeleteScheduleOperation from a given name.
// The name must be that of a previously created DeleteScheduleOperation, possibly from a different process.
func (c *scheduleGRPCClient) DeleteScheduleOperation(name string) *DeleteScheduleOperation {
	return &DeleteScheduleOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}
