// Copyright 2025 Google LLC
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

package tpu

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	tpupb "cloud.google.com/go/tpu/apiv1/tpupb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ListNodes              []gax.CallOption
	GetNode                []gax.CallOption
	CreateNode             []gax.CallOption
	DeleteNode             []gax.CallOption
	ReimageNode            []gax.CallOption
	StopNode               []gax.CallOption
	StartNode              []gax.CallOption
	ListTensorFlowVersions []gax.CallOption
	GetTensorFlowVersion   []gax.CallOption
	ListAcceleratorTypes   []gax.CallOption
	GetAcceleratorType     []gax.CallOption
	GetLocation            []gax.CallOption
	ListLocations          []gax.CallOption
	CancelOperation        []gax.CallOption
	DeleteOperation        []gax.CallOption
	GetOperation           []gax.CallOption
	ListOperations         []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("tpu.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("tpu.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("tpu.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://tpu.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ListNodes: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetNode: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CreateNode: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteNode: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ReimageNode: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		StopNode: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		StartNode: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListTensorFlowVersions: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetTensorFlowVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListAcceleratorTypes: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetAcceleratorType: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetLocation:     []gax.CallOption{},
		ListLocations:   []gax.CallOption{},
		CancelOperation: []gax.CallOption{},
		DeleteOperation: []gax.CallOption{},
		GetOperation:    []gax.CallOption{},
		ListOperations:  []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from Cloud TPU API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListNodes(context.Context, *tpupb.ListNodesRequest, ...gax.CallOption) *NodeIterator
	GetNode(context.Context, *tpupb.GetNodeRequest, ...gax.CallOption) (*tpupb.Node, error)
	CreateNode(context.Context, *tpupb.CreateNodeRequest, ...gax.CallOption) (*CreateNodeOperation, error)
	CreateNodeOperation(name string) *CreateNodeOperation
	DeleteNode(context.Context, *tpupb.DeleteNodeRequest, ...gax.CallOption) (*DeleteNodeOperation, error)
	DeleteNodeOperation(name string) *DeleteNodeOperation
	ReimageNode(context.Context, *tpupb.ReimageNodeRequest, ...gax.CallOption) (*ReimageNodeOperation, error)
	ReimageNodeOperation(name string) *ReimageNodeOperation
	StopNode(context.Context, *tpupb.StopNodeRequest, ...gax.CallOption) (*StopNodeOperation, error)
	StopNodeOperation(name string) *StopNodeOperation
	StartNode(context.Context, *tpupb.StartNodeRequest, ...gax.CallOption) (*StartNodeOperation, error)
	StartNodeOperation(name string) *StartNodeOperation
	ListTensorFlowVersions(context.Context, *tpupb.ListTensorFlowVersionsRequest, ...gax.CallOption) *TensorFlowVersionIterator
	GetTensorFlowVersion(context.Context, *tpupb.GetTensorFlowVersionRequest, ...gax.CallOption) (*tpupb.TensorFlowVersion, error)
	ListAcceleratorTypes(context.Context, *tpupb.ListAcceleratorTypesRequest, ...gax.CallOption) *AcceleratorTypeIterator
	GetAcceleratorType(context.Context, *tpupb.GetAcceleratorTypeRequest, ...gax.CallOption) (*tpupb.AcceleratorType, error)
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// Client is a client for interacting with Cloud TPU API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// # Manages TPU nodes and other resources
//
// TPU API v1
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
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

// ListNodes lists nodes.
func (c *Client) ListNodes(ctx context.Context, req *tpupb.ListNodesRequest, opts ...gax.CallOption) *NodeIterator {
	return c.internalClient.ListNodes(ctx, req, opts...)
}

// GetNode gets the details of a node.
func (c *Client) GetNode(ctx context.Context, req *tpupb.GetNodeRequest, opts ...gax.CallOption) (*tpupb.Node, error) {
	return c.internalClient.GetNode(ctx, req, opts...)
}

// CreateNode creates a node.
func (c *Client) CreateNode(ctx context.Context, req *tpupb.CreateNodeRequest, opts ...gax.CallOption) (*CreateNodeOperation, error) {
	return c.internalClient.CreateNode(ctx, req, opts...)
}

// CreateNodeOperation returns a new CreateNodeOperation from a given name.
// The name must be that of a previously created CreateNodeOperation, possibly from a different process.
func (c *Client) CreateNodeOperation(name string) *CreateNodeOperation {
	return c.internalClient.CreateNodeOperation(name)
}

// DeleteNode deletes a node.
func (c *Client) DeleteNode(ctx context.Context, req *tpupb.DeleteNodeRequest, opts ...gax.CallOption) (*DeleteNodeOperation, error) {
	return c.internalClient.DeleteNode(ctx, req, opts...)
}

// DeleteNodeOperation returns a new DeleteNodeOperation from a given name.
// The name must be that of a previously created DeleteNodeOperation, possibly from a different process.
func (c *Client) DeleteNodeOperation(name string) *DeleteNodeOperation {
	return c.internalClient.DeleteNodeOperation(name)
}

// ReimageNode reimages a node’s OS.
func (c *Client) ReimageNode(ctx context.Context, req *tpupb.ReimageNodeRequest, opts ...gax.CallOption) (*ReimageNodeOperation, error) {
	return c.internalClient.ReimageNode(ctx, req, opts...)
}

// ReimageNodeOperation returns a new ReimageNodeOperation from a given name.
// The name must be that of a previously created ReimageNodeOperation, possibly from a different process.
func (c *Client) ReimageNodeOperation(name string) *ReimageNodeOperation {
	return c.internalClient.ReimageNodeOperation(name)
}

// StopNode stops a node, this operation is only available with single TPU nodes.
func (c *Client) StopNode(ctx context.Context, req *tpupb.StopNodeRequest, opts ...gax.CallOption) (*StopNodeOperation, error) {
	return c.internalClient.StopNode(ctx, req, opts...)
}

// StopNodeOperation returns a new StopNodeOperation from a given name.
// The name must be that of a previously created StopNodeOperation, possibly from a different process.
func (c *Client) StopNodeOperation(name string) *StopNodeOperation {
	return c.internalClient.StopNodeOperation(name)
}

// StartNode starts a node.
func (c *Client) StartNode(ctx context.Context, req *tpupb.StartNodeRequest, opts ...gax.CallOption) (*StartNodeOperation, error) {
	return c.internalClient.StartNode(ctx, req, opts...)
}

// StartNodeOperation returns a new StartNodeOperation from a given name.
// The name must be that of a previously created StartNodeOperation, possibly from a different process.
func (c *Client) StartNodeOperation(name string) *StartNodeOperation {
	return c.internalClient.StartNodeOperation(name)
}

// ListTensorFlowVersions list TensorFlow versions supported by this API.
func (c *Client) ListTensorFlowVersions(ctx context.Context, req *tpupb.ListTensorFlowVersionsRequest, opts ...gax.CallOption) *TensorFlowVersionIterator {
	return c.internalClient.ListTensorFlowVersions(ctx, req, opts...)
}

// GetTensorFlowVersion gets TensorFlow Version.
func (c *Client) GetTensorFlowVersion(ctx context.Context, req *tpupb.GetTensorFlowVersionRequest, opts ...gax.CallOption) (*tpupb.TensorFlowVersion, error) {
	return c.internalClient.GetTensorFlowVersion(ctx, req, opts...)
}

// ListAcceleratorTypes lists accelerator types supported by this API.
func (c *Client) ListAcceleratorTypes(ctx context.Context, req *tpupb.ListAcceleratorTypesRequest, opts ...gax.CallOption) *AcceleratorTypeIterator {
	return c.internalClient.ListAcceleratorTypes(ctx, req, opts...)
}

// GetAcceleratorType gets AcceleratorType.
func (c *Client) GetAcceleratorType(ctx context.Context, req *tpupb.GetAcceleratorTypeRequest, opts ...gax.CallOption) (*tpupb.AcceleratorType, error) {
	return c.internalClient.GetAcceleratorType(ctx, req, opts...)
}

// GetLocation gets information about a location.
func (c *Client) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *Client) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *Client) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *Client) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *Client) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *Client) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Cloud TPU API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client tpupb.TpuClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewClient creates a new tpu client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// # Manages TPU nodes and other resources
//
// TPU API v1
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		client:           tpupb.NewTpuClient(connPool),
		CallOptions:      &client.CallOptions,
		logger:           internaloption.GetLogger(opts),
		operationsClient: longrunningpb.NewOperationsClient(connPool),
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
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) ListNodes(ctx context.Context, req *tpupb.ListNodesRequest, opts ...gax.CallOption) *NodeIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListNodes[0:len((*c.CallOptions).ListNodes):len((*c.CallOptions).ListNodes)], opts...)
	it := &NodeIterator{}
	req = proto.Clone(req).(*tpupb.ListNodesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*tpupb.Node, string, error) {
		resp := &tpupb.ListNodesResponse{}
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
			resp, err = executeRPC(ctx, c.client.ListNodes, req, settings.GRPC, c.logger, "ListNodes")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetNodes(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetNode(ctx context.Context, req *tpupb.GetNodeRequest, opts ...gax.CallOption) (*tpupb.Node, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetNode[0:len((*c.CallOptions).GetNode):len((*c.CallOptions).GetNode)], opts...)
	var resp *tpupb.Node
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetNode, req, settings.GRPC, c.logger, "GetNode")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateNode(ctx context.Context, req *tpupb.CreateNodeRequest, opts ...gax.CallOption) (*CreateNodeOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateNode[0:len((*c.CallOptions).CreateNode):len((*c.CallOptions).CreateNode)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.CreateNode, req, settings.GRPC, c.logger, "CreateNode")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateNodeOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) DeleteNode(ctx context.Context, req *tpupb.DeleteNodeRequest, opts ...gax.CallOption) (*DeleteNodeOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteNode[0:len((*c.CallOptions).DeleteNode):len((*c.CallOptions).DeleteNode)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.DeleteNode, req, settings.GRPC, c.logger, "DeleteNode")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteNodeOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) ReimageNode(ctx context.Context, req *tpupb.ReimageNodeRequest, opts ...gax.CallOption) (*ReimageNodeOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ReimageNode[0:len((*c.CallOptions).ReimageNode):len((*c.CallOptions).ReimageNode)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.ReimageNode, req, settings.GRPC, c.logger, "ReimageNode")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ReimageNodeOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) StopNode(ctx context.Context, req *tpupb.StopNodeRequest, opts ...gax.CallOption) (*StopNodeOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StopNode[0:len((*c.CallOptions).StopNode):len((*c.CallOptions).StopNode)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.StopNode, req, settings.GRPC, c.logger, "StopNode")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &StopNodeOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) StartNode(ctx context.Context, req *tpupb.StartNodeRequest, opts ...gax.CallOption) (*StartNodeOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StartNode[0:len((*c.CallOptions).StartNode):len((*c.CallOptions).StartNode)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.StartNode, req, settings.GRPC, c.logger, "StartNode")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &StartNodeOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) ListTensorFlowVersions(ctx context.Context, req *tpupb.ListTensorFlowVersionsRequest, opts ...gax.CallOption) *TensorFlowVersionIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListTensorFlowVersions[0:len((*c.CallOptions).ListTensorFlowVersions):len((*c.CallOptions).ListTensorFlowVersions)], opts...)
	it := &TensorFlowVersionIterator{}
	req = proto.Clone(req).(*tpupb.ListTensorFlowVersionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*tpupb.TensorFlowVersion, string, error) {
		resp := &tpupb.ListTensorFlowVersionsResponse{}
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
			resp, err = executeRPC(ctx, c.client.ListTensorFlowVersions, req, settings.GRPC, c.logger, "ListTensorFlowVersions")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTensorflowVersions(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetTensorFlowVersion(ctx context.Context, req *tpupb.GetTensorFlowVersionRequest, opts ...gax.CallOption) (*tpupb.TensorFlowVersion, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetTensorFlowVersion[0:len((*c.CallOptions).GetTensorFlowVersion):len((*c.CallOptions).GetTensorFlowVersion)], opts...)
	var resp *tpupb.TensorFlowVersion
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetTensorFlowVersion, req, settings.GRPC, c.logger, "GetTensorFlowVersion")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListAcceleratorTypes(ctx context.Context, req *tpupb.ListAcceleratorTypesRequest, opts ...gax.CallOption) *AcceleratorTypeIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListAcceleratorTypes[0:len((*c.CallOptions).ListAcceleratorTypes):len((*c.CallOptions).ListAcceleratorTypes)], opts...)
	it := &AcceleratorTypeIterator{}
	req = proto.Clone(req).(*tpupb.ListAcceleratorTypesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*tpupb.AcceleratorType, string, error) {
		resp := &tpupb.ListAcceleratorTypesResponse{}
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
			resp, err = executeRPC(ctx, c.client.ListAcceleratorTypes, req, settings.GRPC, c.logger, "ListAcceleratorTypes")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAcceleratorTypes(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetAcceleratorType(ctx context.Context, req *tpupb.GetAcceleratorTypeRequest, opts ...gax.CallOption) (*tpupb.AcceleratorType, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetAcceleratorType[0:len((*c.CallOptions).GetAcceleratorType):len((*c.CallOptions).GetAcceleratorType)], opts...)
	var resp *tpupb.AcceleratorType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetAcceleratorType, req, settings.GRPC, c.logger, "GetAcceleratorType")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.locationsClient.GetLocation, req, settings.GRPC, c.logger, "GetLocation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
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
			resp, err = executeRPC(ctx, c.locationsClient.ListLocations, req, settings.GRPC, c.logger, "ListLocations")
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

func (c *gRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.operationsClient.CancelOperation, req, settings.GRPC, c.logger, "CancelOperation")
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.operationsClient.DeleteOperation, req, settings.GRPC, c.logger, "DeleteOperation")
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.operationsClient.GetOperation, req, settings.GRPC, c.logger, "GetOperation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
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
			resp, err = executeRPC(ctx, c.operationsClient.ListOperations, req, settings.GRPC, c.logger, "ListOperations")
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

// CreateNodeOperation returns a new CreateNodeOperation from a given name.
// The name must be that of a previously created CreateNodeOperation, possibly from a different process.
func (c *gRPCClient) CreateNodeOperation(name string) *CreateNodeOperation {
	return &CreateNodeOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeleteNodeOperation returns a new DeleteNodeOperation from a given name.
// The name must be that of a previously created DeleteNodeOperation, possibly from a different process.
func (c *gRPCClient) DeleteNodeOperation(name string) *DeleteNodeOperation {
	return &DeleteNodeOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// ReimageNodeOperation returns a new ReimageNodeOperation from a given name.
// The name must be that of a previously created ReimageNodeOperation, possibly from a different process.
func (c *gRPCClient) ReimageNodeOperation(name string) *ReimageNodeOperation {
	return &ReimageNodeOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// StartNodeOperation returns a new StartNodeOperation from a given name.
// The name must be that of a previously created StartNodeOperation, possibly from a different process.
func (c *gRPCClient) StartNodeOperation(name string) *StartNodeOperation {
	return &StartNodeOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// StopNodeOperation returns a new StopNodeOperation from a given name.
// The name must be that of a previously created StopNodeOperation, possibly from a different process.
func (c *gRPCClient) StopNodeOperation(name string) *StopNodeOperation {
	return &StopNodeOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}
