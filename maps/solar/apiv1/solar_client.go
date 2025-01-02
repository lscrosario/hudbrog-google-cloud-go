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

package solar

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	solarpb "cloud.google.com/go/maps/solar/apiv1/solarpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	FindClosestBuildingInsights []gax.CallOption
	GetDataLayers               []gax.CallOption
	GetGeoTiff                  []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("solar.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("solar.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("solar.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://solar.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		FindClosestBuildingInsights: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetDataLayers: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetGeoTiff: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultRESTCallOptions() *CallOptions {
	return &CallOptions{
		FindClosestBuildingInsights: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetDataLayers: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetGeoTiff: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalClient is an interface that defines the methods available from Solar API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	FindClosestBuildingInsights(context.Context, *solarpb.FindClosestBuildingInsightsRequest, ...gax.CallOption) (*solarpb.BuildingInsights, error)
	GetDataLayers(context.Context, *solarpb.GetDataLayersRequest, ...gax.CallOption) (*solarpb.DataLayers, error)
	GetGeoTiff(context.Context, *solarpb.GetGeoTiffRequest, ...gax.CallOption) (*httpbodypb.HttpBody, error)
}

// Client is a client for interacting with Solar API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service definition for the Solar API.
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

// FindClosestBuildingInsights locates the closest building to a query point. Returns an error with
// code NOT_FOUND if there are no buildings within approximately 50m of the
// query point.
func (c *Client) FindClosestBuildingInsights(ctx context.Context, req *solarpb.FindClosestBuildingInsightsRequest, opts ...gax.CallOption) (*solarpb.BuildingInsights, error) {
	return c.internalClient.FindClosestBuildingInsights(ctx, req, opts...)
}

// GetDataLayers gets solar information for a region surrounding a location.
// Returns an error with code NOT_FOUND if the location is outside
// the coverage area.
func (c *Client) GetDataLayers(ctx context.Context, req *solarpb.GetDataLayersRequest, opts ...gax.CallOption) (*solarpb.DataLayers, error) {
	return c.internalClient.GetDataLayers(ctx, req, opts...)
}

// GetGeoTiff returns an image by its ID.
func (c *Client) GetGeoTiff(ctx context.Context, req *solarpb.GetGeoTiffRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.GetGeoTiff(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Solar API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client solarpb.SolarClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewClient creates a new solar client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service definition for the Solar API.
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
		connPool:    connPool,
		client:      solarpb.NewSolarClient(connPool),
		CallOptions: &client.CallOptions,
		logger:      internaloption.GetLogger(opts),
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

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type restClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	logger *slog.Logger
}

// NewRESTClient creates a new solar rest client.
//
// Service definition for the Solar API.
func NewRESTClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := append(defaultRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRESTCallOptions()
	c := &restClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://solar.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://solar.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://solar.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://solar.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *restClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *restClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *restClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *gRPCClient) FindClosestBuildingInsights(ctx context.Context, req *solarpb.FindClosestBuildingInsightsRequest, opts ...gax.CallOption) (*solarpb.BuildingInsights, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).FindClosestBuildingInsights[0:len((*c.CallOptions).FindClosestBuildingInsights):len((*c.CallOptions).FindClosestBuildingInsights)], opts...)
	var resp *solarpb.BuildingInsights
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.FindClosestBuildingInsights, req, settings.GRPC, c.logger, "FindClosestBuildingInsights")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetDataLayers(ctx context.Context, req *solarpb.GetDataLayersRequest, opts ...gax.CallOption) (*solarpb.DataLayers, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).GetDataLayers[0:len((*c.CallOptions).GetDataLayers):len((*c.CallOptions).GetDataLayers)], opts...)
	var resp *solarpb.DataLayers
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetDataLayers, req, settings.GRPC, c.logger, "GetDataLayers")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetGeoTiff(ctx context.Context, req *solarpb.GetGeoTiffRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).GetGeoTiff[0:len((*c.CallOptions).GetGeoTiff):len((*c.CallOptions).GetGeoTiff)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetGeoTiff, req, settings.GRPC, c.logger, "GetGeoTiff")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FindClosestBuildingInsights locates the closest building to a query point. Returns an error with
// code NOT_FOUND if there are no buildings within approximately 50m of the
// query point.
func (c *restClient) FindClosestBuildingInsights(ctx context.Context, req *solarpb.FindClosestBuildingInsightsRequest, opts ...gax.CallOption) (*solarpb.BuildingInsights, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/buildingInsights:findClosest")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetExactQualityRequired() {
		params.Add("exactQualityRequired", fmt.Sprintf("%v", req.GetExactQualityRequired()))
	}
	if req.GetLocation().GetLatitude() != 0 {
		params.Add("location.latitude", fmt.Sprintf("%v", req.GetLocation().GetLatitude()))
	}
	if req.GetLocation().GetLongitude() != 0 {
		params.Add("location.longitude", fmt.Sprintf("%v", req.GetLocation().GetLongitude()))
	}
	if req.GetRequiredQuality() != 0 {
		params.Add("requiredQuality", fmt.Sprintf("%v", req.GetRequiredQuality()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).FindClosestBuildingInsights[0:len((*c.CallOptions).FindClosestBuildingInsights):len((*c.CallOptions).FindClosestBuildingInsights)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &solarpb.BuildingInsights{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "FindClosestBuildingInsights")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetDataLayers gets solar information for a region surrounding a location.
// Returns an error with code NOT_FOUND if the location is outside
// the coverage area.
func (c *restClient) GetDataLayers(ctx context.Context, req *solarpb.GetDataLayersRequest, opts ...gax.CallOption) (*solarpb.DataLayers, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/dataLayers:get")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetExactQualityRequired() {
		params.Add("exactQualityRequired", fmt.Sprintf("%v", req.GetExactQualityRequired()))
	}
	if req.GetLocation().GetLatitude() != 0 {
		params.Add("location.latitude", fmt.Sprintf("%v", req.GetLocation().GetLatitude()))
	}
	if req.GetLocation().GetLongitude() != 0 {
		params.Add("location.longitude", fmt.Sprintf("%v", req.GetLocation().GetLongitude()))
	}
	if req.GetPixelSizeMeters() != 0 {
		params.Add("pixelSizeMeters", fmt.Sprintf("%v", req.GetPixelSizeMeters()))
	}
	params.Add("radiusMeters", fmt.Sprintf("%v", req.GetRadiusMeters()))
	if req.GetRequiredQuality() != 0 {
		params.Add("requiredQuality", fmt.Sprintf("%v", req.GetRequiredQuality()))
	}
	if req.GetView() != 0 {
		params.Add("view", fmt.Sprintf("%v", req.GetView()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetDataLayers[0:len((*c.CallOptions).GetDataLayers):len((*c.CallOptions).GetDataLayers)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &solarpb.DataLayers{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetDataLayers")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetGeoTiff returns an image by its ID.
func (c *restClient) GetGeoTiff(ctx context.Context, req *solarpb.GetGeoTiffRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/geoTiff:get")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	params.Add("id", fmt.Sprintf("%v", req.GetId()))

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetGeoTiff[0:len((*c.CallOptions).GetGeoTiff):len((*c.CallOptions).GetGeoTiff)], opts...)
	resp := &httpbodypb.HttpBody{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, httpRsp, err := executeHTTPRequestWithResponse(ctx, c.httpClient, httpReq, c.logger, nil, "GetGeoTiff")
		if err != nil {
			return err
		}

		resp.Data = buf
		if headers := httpRsp.Header; len(headers["Content-Type"]) > 0 {
			resp.ContentType = headers["Content-Type"][0]
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
