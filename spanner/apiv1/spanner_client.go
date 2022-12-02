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

package spanner

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	spannerpb "cloud.google.com/go/spanner/apiv1/spannerpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateSession       []gax.CallOption
	BatchCreateSessions []gax.CallOption
	GetSession          []gax.CallOption
	ListSessions        []gax.CallOption
	DeleteSession       []gax.CallOption
	ExecuteSql          []gax.CallOption
	ExecuteStreamingSql []gax.CallOption
	ExecuteBatchDml     []gax.CallOption
	Read                []gax.CallOption
	StreamingRead       []gax.CallOption
	BeginTransaction    []gax.CallOption
	Commit              []gax.CallOption
	Rollback            []gax.CallOption
	PartitionQuery      []gax.CallOption
	PartitionRead       []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("spanner.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("spanner.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://spanner.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateSession: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchCreateSessions: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetSession: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListSessions: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteSession: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ExecuteSql: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ExecuteStreamingSql: []gax.CallOption{},
		ExecuteBatchDml: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		Read: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		StreamingRead: []gax.CallOption{},
		BeginTransaction: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		Commit: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		Rollback: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		PartitionQuery: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		PartitionRead: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    250 * time.Millisecond,
					Max:        32000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalClient is an interface that defines the methods available from Cloud Spanner API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateSession(context.Context, *spannerpb.CreateSessionRequest, ...gax.CallOption) (*spannerpb.Session, error)
	BatchCreateSessions(context.Context, *spannerpb.BatchCreateSessionsRequest, ...gax.CallOption) (*spannerpb.BatchCreateSessionsResponse, error)
	GetSession(context.Context, *spannerpb.GetSessionRequest, ...gax.CallOption) (*spannerpb.Session, error)
	ListSessions(context.Context, *spannerpb.ListSessionsRequest, ...gax.CallOption) *SessionIterator
	DeleteSession(context.Context, *spannerpb.DeleteSessionRequest, ...gax.CallOption) error
	ExecuteSql(context.Context, *spannerpb.ExecuteSqlRequest, ...gax.CallOption) (*spannerpb.ResultSet, error)
	ExecuteStreamingSql(context.Context, *spannerpb.ExecuteSqlRequest, ...gax.CallOption) (spannerpb.Spanner_ExecuteStreamingSqlClient, error)
	ExecuteBatchDml(context.Context, *spannerpb.ExecuteBatchDmlRequest, ...gax.CallOption) (*spannerpb.ExecuteBatchDmlResponse, error)
	Read(context.Context, *spannerpb.ReadRequest, ...gax.CallOption) (*spannerpb.ResultSet, error)
	StreamingRead(context.Context, *spannerpb.ReadRequest, ...gax.CallOption) (spannerpb.Spanner_StreamingReadClient, error)
	BeginTransaction(context.Context, *spannerpb.BeginTransactionRequest, ...gax.CallOption) (*spannerpb.Transaction, error)
	Commit(context.Context, *spannerpb.CommitRequest, ...gax.CallOption) (*spannerpb.CommitResponse, error)
	Rollback(context.Context, *spannerpb.RollbackRequest, ...gax.CallOption) error
	PartitionQuery(context.Context, *spannerpb.PartitionQueryRequest, ...gax.CallOption) (*spannerpb.PartitionResponse, error)
	PartitionRead(context.Context, *spannerpb.PartitionReadRequest, ...gax.CallOption) (*spannerpb.PartitionResponse, error)
}

// Client is a client for interacting with Cloud Spanner API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// # Cloud Spanner API
//
// The Cloud Spanner API can be used to manage sessions and execute
// transactions on data stored in Cloud Spanner databases.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateSession creates a new session. A session can be used to perform
// transactions that read and/or modify data in a Cloud Spanner database.
// Sessions are meant to be reused for many consecutive
// transactions.
//
// Sessions can only execute one transaction at a time. To execute
// multiple concurrent read-write/write-only transactions, create
// multiple sessions. Note that standalone reads and queries use a
// transaction internally, and count toward the one transaction
// limit.
//
// Active sessions use additional server resources, so it is a good idea to
// delete idle and unneeded sessions.
// Aside from explicit deletes, Cloud Spanner may delete sessions for which no
// operations are sent for more than an hour. If a session is deleted,
// requests to it return NOT_FOUND.
//
// Idle sessions can be kept alive by sending a trivial SQL query
// periodically, e.g., "SELECT 1".
func (c *Client) CreateSession(ctx context.Context, req *spannerpb.CreateSessionRequest, opts ...gax.CallOption) (*spannerpb.Session, error) {
	return c.internalClient.CreateSession(ctx, req, opts...)
}

// BatchCreateSessions creates multiple new sessions.
//
// This API can be used to initialize a session cache on the clients.
// See https://goo.gl/TgSFN2 (at https://goo.gl/TgSFN2) for best practices on session cache management.
func (c *Client) BatchCreateSessions(ctx context.Context, req *spannerpb.BatchCreateSessionsRequest, opts ...gax.CallOption) (*spannerpb.BatchCreateSessionsResponse, error) {
	return c.internalClient.BatchCreateSessions(ctx, req, opts...)
}

// GetSession gets a session. Returns NOT_FOUND if the session does not exist.
// This is mainly useful for determining whether a session is still
// alive.
func (c *Client) GetSession(ctx context.Context, req *spannerpb.GetSessionRequest, opts ...gax.CallOption) (*spannerpb.Session, error) {
	return c.internalClient.GetSession(ctx, req, opts...)
}

// ListSessions lists all sessions in a given database.
func (c *Client) ListSessions(ctx context.Context, req *spannerpb.ListSessionsRequest, opts ...gax.CallOption) *SessionIterator {
	return c.internalClient.ListSessions(ctx, req, opts...)
}

// DeleteSession ends a session, releasing server resources associated with it. This will
// asynchronously trigger cancellation of any operations that are running with
// this session.
func (c *Client) DeleteSession(ctx context.Context, req *spannerpb.DeleteSessionRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteSession(ctx, req, opts...)
}

// ExecuteSql executes an SQL statement, returning all results in a single reply. This
// method cannot be used to return a result set larger than 10 MiB;
// if the query yields more data than that, the query fails with
// a FAILED_PRECONDITION error.
//
// Operations inside read-write transactions might return ABORTED. If
// this occurs, the application should restart the transaction from
// the beginning. See Transaction for more details.
//
// Larger result sets can be fetched in streaming fashion by calling
// ExecuteStreamingSql instead.
func (c *Client) ExecuteSql(ctx context.Context, req *spannerpb.ExecuteSqlRequest, opts ...gax.CallOption) (*spannerpb.ResultSet, error) {
	return c.internalClient.ExecuteSql(ctx, req, opts...)
}

// ExecuteStreamingSql like ExecuteSql, except returns the result
// set as a stream. Unlike ExecuteSql, there
// is no limit on the size of the returned result set. However, no
// individual row in the result set can exceed 100 MiB, and no
// column value can exceed 10 MiB.
func (c *Client) ExecuteStreamingSql(ctx context.Context, req *spannerpb.ExecuteSqlRequest, opts ...gax.CallOption) (spannerpb.Spanner_ExecuteStreamingSqlClient, error) {
	return c.internalClient.ExecuteStreamingSql(ctx, req, opts...)
}

// ExecuteBatchDml executes a batch of SQL DML statements. This method allows many statements
// to be run with lower latency than submitting them sequentially with
// ExecuteSql.
//
// Statements are executed in sequential order. A request can succeed even if
// a statement fails. The ExecuteBatchDmlResponse.status field in the
// response provides information about the statement that failed. Clients must
// inspect this field to determine whether an error occurred.
//
// Execution stops after the first failed statement; the remaining statements
// are not executed.
func (c *Client) ExecuteBatchDml(ctx context.Context, req *spannerpb.ExecuteBatchDmlRequest, opts ...gax.CallOption) (*spannerpb.ExecuteBatchDmlResponse, error) {
	return c.internalClient.ExecuteBatchDml(ctx, req, opts...)
}

// Read reads rows from the database using key lookups and scans, as a
// simple key/value style alternative to
// ExecuteSql.  This method cannot be used to
// return a result set larger than 10 MiB; if the read matches more
// data than that, the read fails with a FAILED_PRECONDITION
// error.
//
// Reads inside read-write transactions might return ABORTED. If
// this occurs, the application should restart the transaction from
// the beginning. See Transaction for more details.
//
// Larger result sets can be yielded in streaming fashion by calling
// StreamingRead instead.
func (c *Client) Read(ctx context.Context, req *spannerpb.ReadRequest, opts ...gax.CallOption) (*spannerpb.ResultSet, error) {
	return c.internalClient.Read(ctx, req, opts...)
}

// StreamingRead like Read, except returns the result set as a
// stream. Unlike Read, there is no limit on the
// size of the returned result set. However, no individual row in
// the result set can exceed 100 MiB, and no column value can exceed
// 10 MiB.
func (c *Client) StreamingRead(ctx context.Context, req *spannerpb.ReadRequest, opts ...gax.CallOption) (spannerpb.Spanner_StreamingReadClient, error) {
	return c.internalClient.StreamingRead(ctx, req, opts...)
}

// BeginTransaction begins a new transaction. This step can often be skipped:
// Read, ExecuteSql and
// Commit can begin a new transaction as a
// side-effect.
func (c *Client) BeginTransaction(ctx context.Context, req *spannerpb.BeginTransactionRequest, opts ...gax.CallOption) (*spannerpb.Transaction, error) {
	return c.internalClient.BeginTransaction(ctx, req, opts...)
}

// Commit commits a transaction. The request includes the mutations to be
// applied to rows in the database.
//
// Commit might return an ABORTED error. This can occur at any time;
// commonly, the cause is conflicts with concurrent
// transactions. However, it can also happen for a variety of other
// reasons. If Commit returns ABORTED, the caller should re-attempt
// the transaction from the beginning, re-using the same session.
//
// On very rare occasions, Commit might return UNKNOWN. This can happen,
// for example, if the client job experiences a 1+ hour networking failure.
// At that point, Cloud Spanner has lost track of the transaction outcome and
// we recommend that you perform another read from the database to see the
// state of things as they are now.
func (c *Client) Commit(ctx context.Context, req *spannerpb.CommitRequest, opts ...gax.CallOption) (*spannerpb.CommitResponse, error) {
	return c.internalClient.Commit(ctx, req, opts...)
}

// Rollback rolls back a transaction, releasing any locks it holds. It is a good
// idea to call this for any transaction that includes one or more
// Read or ExecuteSql requests and
// ultimately decides not to commit.
//
// Rollback returns OK if it successfully aborts the transaction, the
// transaction was already aborted, or the transaction is not
// found. Rollback never returns ABORTED.
func (c *Client) Rollback(ctx context.Context, req *spannerpb.RollbackRequest, opts ...gax.CallOption) error {
	return c.internalClient.Rollback(ctx, req, opts...)
}

// PartitionQuery creates a set of partition tokens that can be used to execute a query
// operation in parallel.  Each of the returned partition tokens can be used
// by ExecuteStreamingSql to specify a subset
// of the query result to read.  The same session and read-only transaction
// must be used by the PartitionQueryRequest used to create the
// partition tokens and the ExecuteSqlRequests that use the partition tokens.
//
// Partition tokens become invalid when the session used to create them
// is deleted, is idle for too long, begins a new transaction, or becomes too
// old.  When any of these happen, it is not possible to resume the query, and
// the whole operation must be restarted from the beginning.
func (c *Client) PartitionQuery(ctx context.Context, req *spannerpb.PartitionQueryRequest, opts ...gax.CallOption) (*spannerpb.PartitionResponse, error) {
	return c.internalClient.PartitionQuery(ctx, req, opts...)
}

// PartitionRead creates a set of partition tokens that can be used to execute a read
// operation in parallel.  Each of the returned partition tokens can be used
// by StreamingRead to specify a subset of the read
// result to read.  The same session and read-only transaction must be used by
// the PartitionReadRequest used to create the partition tokens and the
// ReadRequests that use the partition tokens.  There are no ordering
// guarantees on rows returned among the returned partition tokens, or even
// within each individual StreamingRead call issued with a partition_token.
//
// Partition tokens become invalid when the session used to create them
// is deleted, is idle for too long, begins a new transaction, or becomes too
// old.  When any of these happen, it is not possible to resume the read, and
// the whole operation must be restarted from the beginning.
func (c *Client) PartitionRead(ctx context.Context, req *spannerpb.PartitionReadRequest, opts ...gax.CallOption) (*spannerpb.PartitionResponse, error) {
	return c.internalClient.PartitionRead(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Cloud Spanner API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client spannerpb.SpannerClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new spanner client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// # Cloud Spanner API
//
// The Cloud Spanner API can be used to manage sessions and execute
// transactions on data stored in Cloud Spanner databases.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
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
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           spannerpb.NewSpannerClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) CreateSession(ctx context.Context, req *spannerpb.CreateSessionRequest, opts ...gax.CallOption) (*spannerpb.Session, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "database", url.QueryEscape(req.GetDatabase())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateSession[0:len((*c.CallOptions).CreateSession):len((*c.CallOptions).CreateSession)], opts...)
	var resp *spannerpb.Session
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) BatchCreateSessions(ctx context.Context, req *spannerpb.BatchCreateSessionsRequest, opts ...gax.CallOption) (*spannerpb.BatchCreateSessionsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "database", url.QueryEscape(req.GetDatabase())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchCreateSessions[0:len((*c.CallOptions).BatchCreateSessions):len((*c.CallOptions).BatchCreateSessions)], opts...)
	var resp *spannerpb.BatchCreateSessionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.BatchCreateSessions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetSession(ctx context.Context, req *spannerpb.GetSessionRequest, opts ...gax.CallOption) (*spannerpb.Session, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetSession[0:len((*c.CallOptions).GetSession):len((*c.CallOptions).GetSession)], opts...)
	var resp *spannerpb.Session
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListSessions(ctx context.Context, req *spannerpb.ListSessionsRequest, opts ...gax.CallOption) *SessionIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "database", url.QueryEscape(req.GetDatabase())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListSessions[0:len((*c.CallOptions).ListSessions):len((*c.CallOptions).ListSessions)], opts...)
	it := &SessionIterator{}
	req = proto.Clone(req).(*spannerpb.ListSessionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*spannerpb.Session, string, error) {
		resp := &spannerpb.ListSessionsResponse{}
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
			resp, err = c.client.ListSessions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSessions(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) DeleteSession(ctx context.Context, req *spannerpb.DeleteSessionRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteSession[0:len((*c.CallOptions).DeleteSession):len((*c.CallOptions).DeleteSession)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) ExecuteSql(ctx context.Context, req *spannerpb.ExecuteSqlRequest, opts ...gax.CallOption) (*spannerpb.ResultSet, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ExecuteSql[0:len((*c.CallOptions).ExecuteSql):len((*c.CallOptions).ExecuteSql)], opts...)
	var resp *spannerpb.ResultSet
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ExecuteSql(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ExecuteStreamingSql(ctx context.Context, req *spannerpb.ExecuteSqlRequest, opts ...gax.CallOption) (spannerpb.Spanner_ExecuteStreamingSqlClient, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	var resp spannerpb.Spanner_ExecuteStreamingSqlClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ExecuteStreamingSql(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ExecuteBatchDml(ctx context.Context, req *spannerpb.ExecuteBatchDmlRequest, opts ...gax.CallOption) (*spannerpb.ExecuteBatchDmlResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ExecuteBatchDml[0:len((*c.CallOptions).ExecuteBatchDml):len((*c.CallOptions).ExecuteBatchDml)], opts...)
	var resp *spannerpb.ExecuteBatchDmlResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ExecuteBatchDml(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) Read(ctx context.Context, req *spannerpb.ReadRequest, opts ...gax.CallOption) (*spannerpb.ResultSet, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).Read[0:len((*c.CallOptions).Read):len((*c.CallOptions).Read)], opts...)
	var resp *spannerpb.ResultSet
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.Read(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) StreamingRead(ctx context.Context, req *spannerpb.ReadRequest, opts ...gax.CallOption) (spannerpb.Spanner_StreamingReadClient, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	var resp spannerpb.Spanner_StreamingReadClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.StreamingRead(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) BeginTransaction(ctx context.Context, req *spannerpb.BeginTransactionRequest, opts ...gax.CallOption) (*spannerpb.Transaction, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BeginTransaction[0:len((*c.CallOptions).BeginTransaction):len((*c.CallOptions).BeginTransaction)], opts...)
	var resp *spannerpb.Transaction
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.BeginTransaction(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) Commit(ctx context.Context, req *spannerpb.CommitRequest, opts ...gax.CallOption) (*spannerpb.CommitResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).Commit[0:len((*c.CallOptions).Commit):len((*c.CallOptions).Commit)], opts...)
	var resp *spannerpb.CommitResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.Commit(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) Rollback(ctx context.Context, req *spannerpb.RollbackRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).Rollback[0:len((*c.CallOptions).Rollback):len((*c.CallOptions).Rollback)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.Rollback(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) PartitionQuery(ctx context.Context, req *spannerpb.PartitionQueryRequest, opts ...gax.CallOption) (*spannerpb.PartitionResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).PartitionQuery[0:len((*c.CallOptions).PartitionQuery):len((*c.CallOptions).PartitionQuery)], opts...)
	var resp *spannerpb.PartitionResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.PartitionQuery(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) PartitionRead(ctx context.Context, req *spannerpb.PartitionReadRequest, opts ...gax.CallOption) (*spannerpb.PartitionResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).PartitionRead[0:len((*c.CallOptions).PartitionRead):len((*c.CallOptions).PartitionRead)], opts...)
	var resp *spannerpb.PartitionResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.PartitionRead(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SessionIterator manages a stream of *spannerpb.Session.
type SessionIterator struct {
	items    []*spannerpb.Session
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*spannerpb.Session, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SessionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SessionIterator) Next() (*spannerpb.Session, error) {
	var item *spannerpb.Session
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SessionIterator) bufLen() int {
	return len(it.items)
}

func (it *SessionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
