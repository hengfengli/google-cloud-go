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

package gateway

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"

	gatewaypb "cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	GetResource    []gax.CallOption
	PostResource   []gax.CallOption
	DeleteResource []gax.CallOption
	PutResource    []gax.CallOption
	PatchResource  []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("connectgateway.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("connectgateway.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://connectgateway.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		GetResource:    []gax.CallOption{},
		PostResource:   []gax.CallOption{},
		DeleteResource: []gax.CallOption{},
		PutResource:    []gax.CallOption{},
		PatchResource:  []gax.CallOption{},
	}
}

func defaultRESTCallOptions() *CallOptions {
	return &CallOptions{
		GetResource:    []gax.CallOption{},
		PostResource:   []gax.CallOption{},
		DeleteResource: []gax.CallOption{},
		PutResource:    []gax.CallOption{},
		PatchResource:  []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from Connect Gateway API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetResource(context.Context, *httpbodypb.HttpBody, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	PostResource(context.Context, *httpbodypb.HttpBody, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	DeleteResource(context.Context, *httpbodypb.HttpBody, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	PutResource(context.Context, *httpbodypb.HttpBody, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	PatchResource(context.Context, *httpbodypb.HttpBody, ...gax.CallOption) (*httpbodypb.HttpBody, error)
}

// Client is a client for interacting with Connect Gateway API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Gateway service is a public API which works as a Kubernetes resource model
// proxy between end users and registered Kubernetes clusters. Each RPC in this
// service matches with an HTTP verb. End user will initiate kubectl commands
// against the Gateway service, and Gateway service will forward user requests
// to clusters.
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

// GetResource getResource performs an HTTP GET request on the Kubernetes API Server.
func (c *Client) GetResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.GetResource(ctx, req, opts...)
}

// PostResource postResource performs an HTTP POST on the Kubernetes API Server.
func (c *Client) PostResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.PostResource(ctx, req, opts...)
}

// DeleteResource deleteResource performs an HTTP DELETE on the Kubernetes API Server.
func (c *Client) DeleteResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.DeleteResource(ctx, req, opts...)
}

// PutResource putResource performs an HTTP PUT on the Kubernetes API Server.
func (c *Client) PutResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.PutResource(ctx, req, opts...)
}

// PatchResource patchResource performs an HTTP PATCH on the Kubernetes API Server.
func (c *Client) PatchResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.PatchResource(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Connect Gateway API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client gatewaypb.GatewayServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewClient creates a new gateway service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Gateway service is a public API which works as a Kubernetes resource model
// proxy between end users and registered Kubernetes clusters. Each RPC in this
// service matches with an HTTP verb. End user will initiate kubectl commands
// against the Gateway service, and Gateway service will forward user requests
// to clusters.
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
		client:      gatewaypb.NewGatewayServiceClient(connPool),
		CallOptions: &client.CallOptions,
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
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
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
}

// NewRESTClient creates a new gateway service rest client.
//
// Gateway service is a public API which works as a Kubernetes resource model
// proxy between end users and registered Kubernetes clusters. Each RPC in this
// service matches with an HTTP verb. End user will initiate kubectl commands
// against the Gateway service, and Gateway service will forward user requests
// to clusters.
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
	}
	c.setGoogleClientInfo()

	return &Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://connectgateway.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://connectgateway.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://connectgateway.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *restClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
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
func (c *gRPCClient) GetResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).GetResource[0:len((*c.CallOptions).GetResource):len((*c.CallOptions).GetResource)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) PostResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).PostResource[0:len((*c.CallOptions).PostResource):len((*c.CallOptions).PostResource)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.PostResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).DeleteResource[0:len((*c.CallOptions).DeleteResource):len((*c.CallOptions).DeleteResource)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.DeleteResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) PutResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).PutResource[0:len((*c.CallOptions).PutResource):len((*c.CallOptions).PutResource)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.PutResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) PatchResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).PatchResource[0:len((*c.CallOptions).PatchResource):len((*c.CallOptions).PatchResource)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.PatchResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetResource getResource performs an HTTP GET request on the Kubernetes API Server.
func (c *restClient) GetResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/**")

	params := url.Values{}
	if req.GetContentType() != "" {
		params.Add("contentType", fmt.Sprintf("%v", req.GetContentType()))
	}
	if req.GetData() != nil {
		params.Add("data", fmt.Sprintf("%v", req.GetData()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetResource[0:len((*c.CallOptions).GetResource):len((*c.CallOptions).GetResource)], opts...)
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

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
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

// PostResource postResource performs an HTTP POST on the Kubernetes API Server.
func (c *restClient) PostResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/**")

	params := url.Values{}
	if req.GetContentType() != "" {
		params.Add("contentType", fmt.Sprintf("%v", req.GetContentType()))
	}
	if req.GetData() != nil {
		params.Add("data", fmt.Sprintf("%v", req.GetData()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).PostResource[0:len((*c.CallOptions).PostResource):len((*c.CallOptions).PostResource)], opts...)
	resp := &httpbodypb.HttpBody{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
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

		buf, err := io.ReadAll(httpRsp.Body)
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

// DeleteResource deleteResource performs an HTTP DELETE on the Kubernetes API Server.
func (c *restClient) DeleteResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/**")

	params := url.Values{}
	if req.GetContentType() != "" {
		params.Add("contentType", fmt.Sprintf("%v", req.GetContentType()))
	}
	if req.GetData() != nil {
		params.Add("data", fmt.Sprintf("%v", req.GetData()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).DeleteResource[0:len((*c.CallOptions).DeleteResource):len((*c.CallOptions).DeleteResource)], opts...)
	resp := &httpbodypb.HttpBody{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
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

		buf, err := io.ReadAll(httpRsp.Body)
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

// PutResource putResource performs an HTTP PUT on the Kubernetes API Server.
func (c *restClient) PutResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/**")

	params := url.Values{}
	if req.GetContentType() != "" {
		params.Add("contentType", fmt.Sprintf("%v", req.GetContentType()))
	}
	if req.GetData() != nil {
		params.Add("data", fmt.Sprintf("%v", req.GetData()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).PutResource[0:len((*c.CallOptions).PutResource):len((*c.CallOptions).PutResource)], opts...)
	resp := &httpbodypb.HttpBody{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PUT", baseUrl.String(), nil)
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

		buf, err := io.ReadAll(httpRsp.Body)
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

// PatchResource patchResource performs an HTTP PATCH on the Kubernetes API Server.
func (c *restClient) PatchResource(ctx context.Context, req *httpbodypb.HttpBody, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/**")

	params := url.Values{}
	if req.GetContentType() != "" {
		params.Add("contentType", fmt.Sprintf("%v", req.GetContentType()))
	}
	if req.GetData() != nil {
		params.Add("data", fmt.Sprintf("%v", req.GetData()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).PatchResource[0:len((*c.CallOptions).PatchResource):len((*c.CallOptions).PatchResource)], opts...)
	resp := &httpbodypb.HttpBody{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), nil)
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

		buf, err := io.ReadAll(httpRsp.Body)
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
