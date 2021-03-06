// Copyright 2019-2020 Enseada authors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by protoc-gen-twirp v5.10.0, DO NOT EDIT.
// source: auth/v1beta1/o_auth_clients_api.proto

package authv1beta1

import (
	bytes "bytes"
	strings "strings"

	context "context"

	fmt "fmt"

	ioutil "io/ioutil"

	http "net/http"

	strconv "strconv"

	jsonpb "github.com/golang/protobuf/jsonpb"

	proto "github.com/golang/protobuf/proto"

	twirp "github.com/twitchtv/twirp"

	ctxsetters "github.com/twitchtv/twirp/ctxsetters"
)

// =========================
// OAuthClientsAPI Interface
// =========================

// RPC API to manage OAuth clients.
type OAuthClientsAPI interface {
	// List all available OAuth clients.
	// OAuth scopes: oauth:client.read
	ListClients(context.Context, *ListClientsRequest) (*ListClientsResponse, error)

	// Get a single OAuth client.
	// OAuth scopes: oauth:client.read
	GetClient(context.Context, *GetClientRequest) (*GetClientResponse, error)

	// Add a new OAuth clients.
	// OAuth scopes: oauth:client.write
	CreateClient(context.Context, *CreateClientRequest) (*CreateClientResponse, error)

	// Update an OAuth client.
	// OAuth scopes: oauth:client.write
	UpdateClient(context.Context, *UpdateClientRequest) (*UpdateClientResponse, error)

	// Delete an OAuth clients.
	// OAuth scopes: oauth:client.write
	DeleteClient(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error)
}

// ===============================
// OAuthClientsAPI Protobuf Client
// ===============================

type oAuthClientsAPIProtobufClient struct {
	client HTTPClient
	urls   [5]string
	opts   twirp.ClientOptions
}

// NewOAuthClientsAPIProtobufClient creates a Protobuf client that implements the OAuthClientsAPI interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewOAuthClientsAPIProtobufClient(addr string, client HTTPClient, opts ...twirp.ClientOption) OAuthClientsAPI {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + OAuthClientsAPIPathPrefix
	urls := [5]string{
		prefix + "ListClients",
		prefix + "GetClient",
		prefix + "CreateClient",
		prefix + "UpdateClient",
		prefix + "DeleteClient",
	}

	return &oAuthClientsAPIProtobufClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *oAuthClientsAPIProtobufClient) ListClients(ctx context.Context, in *ListClientsRequest) (*ListClientsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithMethodName(ctx, "ListClients")
	out := new(ListClientsResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *oAuthClientsAPIProtobufClient) GetClient(ctx context.Context, in *GetClientRequest) (*GetClientResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithMethodName(ctx, "GetClient")
	out := new(GetClientResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *oAuthClientsAPIProtobufClient) CreateClient(ctx context.Context, in *CreateClientRequest) (*CreateClientResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithMethodName(ctx, "CreateClient")
	out := new(CreateClientResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *oAuthClientsAPIProtobufClient) UpdateClient(ctx context.Context, in *UpdateClientRequest) (*UpdateClientResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithMethodName(ctx, "UpdateClient")
	out := new(UpdateClientResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[3], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *oAuthClientsAPIProtobufClient) DeleteClient(ctx context.Context, in *DeleteClientRequest) (*DeleteClientResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithMethodName(ctx, "DeleteClient")
	out := new(DeleteClientResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[4], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===========================
// OAuthClientsAPI JSON Client
// ===========================

type oAuthClientsAPIJSONClient struct {
	client HTTPClient
	urls   [5]string
	opts   twirp.ClientOptions
}

// NewOAuthClientsAPIJSONClient creates a JSON client that implements the OAuthClientsAPI interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewOAuthClientsAPIJSONClient(addr string, client HTTPClient, opts ...twirp.ClientOption) OAuthClientsAPI {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + OAuthClientsAPIPathPrefix
	urls := [5]string{
		prefix + "ListClients",
		prefix + "GetClient",
		prefix + "CreateClient",
		prefix + "UpdateClient",
		prefix + "DeleteClient",
	}

	return &oAuthClientsAPIJSONClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *oAuthClientsAPIJSONClient) ListClients(ctx context.Context, in *ListClientsRequest) (*ListClientsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithMethodName(ctx, "ListClients")
	out := new(ListClientsResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *oAuthClientsAPIJSONClient) GetClient(ctx context.Context, in *GetClientRequest) (*GetClientResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithMethodName(ctx, "GetClient")
	out := new(GetClientResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *oAuthClientsAPIJSONClient) CreateClient(ctx context.Context, in *CreateClientRequest) (*CreateClientResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithMethodName(ctx, "CreateClient")
	out := new(CreateClientResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *oAuthClientsAPIJSONClient) UpdateClient(ctx context.Context, in *UpdateClientRequest) (*UpdateClientResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithMethodName(ctx, "UpdateClient")
	out := new(UpdateClientResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[3], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *oAuthClientsAPIJSONClient) DeleteClient(ctx context.Context, in *DeleteClientRequest) (*DeleteClientResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithMethodName(ctx, "DeleteClient")
	out := new(DeleteClientResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[4], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ==============================
// OAuthClientsAPI Server Handler
// ==============================

type oAuthClientsAPIServer struct {
	OAuthClientsAPI
	hooks *twirp.ServerHooks
}

func NewOAuthClientsAPIServer(svc OAuthClientsAPI, hooks *twirp.ServerHooks) TwirpServer {
	return &oAuthClientsAPIServer{
		OAuthClientsAPI: svc,
		hooks:           hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *oAuthClientsAPIServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// OAuthClientsAPIPathPrefix is used for all URL paths on a twirp OAuthClientsAPI server.
// Requests are always: POST OAuthClientsAPIPathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const OAuthClientsAPIPathPrefix = "/twirp/auth.v1beta1.OAuthClientsAPI/"

func (s *oAuthClientsAPIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "auth.v1beta1")
	ctx = ctxsetters.WithServiceName(ctx, "OAuthClientsAPI")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/twirp/auth.v1beta1.OAuthClientsAPI/ListClients":
		s.serveListClients(ctx, resp, req)
		return
	case "/twirp/auth.v1beta1.OAuthClientsAPI/GetClient":
		s.serveGetClient(ctx, resp, req)
		return
	case "/twirp/auth.v1beta1.OAuthClientsAPI/CreateClient":
		s.serveCreateClient(ctx, resp, req)
		return
	case "/twirp/auth.v1beta1.OAuthClientsAPI/UpdateClient":
		s.serveUpdateClient(ctx, resp, req)
		return
	case "/twirp/auth.v1beta1.OAuthClientsAPI/DeleteClient":
		s.serveDeleteClient(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *oAuthClientsAPIServer) serveListClients(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveListClientsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveListClientsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *oAuthClientsAPIServer) serveListClientsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListClients")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(ListClientsRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *ListClientsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.OAuthClientsAPI.ListClients(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListClientsResponse and nil error while calling ListClients. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *oAuthClientsAPIServer) serveListClientsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListClients")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(ListClientsRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *ListClientsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.OAuthClientsAPI.ListClients(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListClientsResponse and nil error while calling ListClients. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *oAuthClientsAPIServer) serveGetClient(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetClientJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveGetClientProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *oAuthClientsAPIServer) serveGetClientJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetClient")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(GetClientRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *GetClientResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.OAuthClientsAPI.GetClient(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *GetClientResponse and nil error while calling GetClient. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *oAuthClientsAPIServer) serveGetClientProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetClient")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(GetClientRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *GetClientResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.OAuthClientsAPI.GetClient(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *GetClientResponse and nil error while calling GetClient. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *oAuthClientsAPIServer) serveCreateClient(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateClientJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateClientProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *oAuthClientsAPIServer) serveCreateClientJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateClient")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(CreateClientRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *CreateClientResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.OAuthClientsAPI.CreateClient(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateClientResponse and nil error while calling CreateClient. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *oAuthClientsAPIServer) serveCreateClientProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateClient")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(CreateClientRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *CreateClientResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.OAuthClientsAPI.CreateClient(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateClientResponse and nil error while calling CreateClient. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *oAuthClientsAPIServer) serveUpdateClient(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveUpdateClientJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveUpdateClientProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *oAuthClientsAPIServer) serveUpdateClientJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "UpdateClient")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(UpdateClientRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *UpdateClientResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.OAuthClientsAPI.UpdateClient(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UpdateClientResponse and nil error while calling UpdateClient. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *oAuthClientsAPIServer) serveUpdateClientProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "UpdateClient")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(UpdateClientRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *UpdateClientResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.OAuthClientsAPI.UpdateClient(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UpdateClientResponse and nil error while calling UpdateClient. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *oAuthClientsAPIServer) serveDeleteClient(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveDeleteClientJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveDeleteClientProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *oAuthClientsAPIServer) serveDeleteClientJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "DeleteClient")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(DeleteClientRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *DeleteClientResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.OAuthClientsAPI.DeleteClient(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *DeleteClientResponse and nil error while calling DeleteClient. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *oAuthClientsAPIServer) serveDeleteClientProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "DeleteClient")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(DeleteClientRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *DeleteClientResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.OAuthClientsAPI.DeleteClient(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *DeleteClientResponse and nil error while calling DeleteClient. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *oAuthClientsAPIServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor1, 0
}

func (s *oAuthClientsAPIServer) ProtocGenTwirpVersion() string {
	return "v5.10.0"
}

func (s *oAuthClientsAPIServer) PathPrefix() string {
	return OAuthClientsAPIPathPrefix
}

var twirpFileDescriptor1 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x51, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x49, 0x0a, 0xbd, 0xf4, 0x34, 0xdc, 0x7b, 0x9d, 0x04, 0xad, 0x79, 0xd0, 0x38, 0x50,
	0xe8, 0x53, 0x4a, 0xda, 0x15, 0xb4, 0x15, 0xb5, 0x52, 0xb0, 0x04, 0x2a, 0x45, 0x84, 0x9a, 0x36,
	0x07, 0x1a, 0x28, 0x4d, 0xec, 0x4c, 0x5d, 0x90, 0x8f, 0xae, 0xc0, 0x35, 0xb8, 0x2a, 0x49, 0x33,
	0x6a, 0x66, 0x4c, 0x0b, 0x16, 0x1f, 0xcf, 0xc9, 0x97, 0xff, 0xff, 0xe1, 0xfc, 0x0c, 0xd4, 0x83,
	0x35, 0x9f, 0x37, 0x9f, 0xbc, 0x29, 0xf2, 0xc0, 0x6b, 0xc6, 0x93, 0x74, 0x9c, 0xcc, 0x16, 0x11,
	0x2e, 0x39, 0x9b, 0x04, 0x49, 0xe4, 0x26, 0xab, 0x98, 0xc7, 0xc4, 0x48, 0xf7, 0xae, 0xc0, 0x6c,
	0x67, 0xfb, 0x4f, 0x19, 0x4f, 0x2d, 0x20, 0x83, 0x88, 0xf1, 0x5e, 0x26, 0xe4, 0xe3, 0xe3, 0x1a,
	0x19, 0xa7, 0xd7, 0x60, 0x4a, 0x5b, 0x96, 0xc4, 0x4b, 0x86, 0xa4, 0x0d, 0x7f, 0x84, 0x63, 0x4d,
	0x73, 0x4a, 0x8d, 0x6a, 0xeb, 0xd8, 0xcd, 0xdb, 0xb9, 0x37, 0x9d, 0x35, 0x9f, 0x67, 0x3f, 0xf9,
	0x1f, 0x24, 0xa5, 0xf0, 0xff, 0x12, 0x85, 0x94, 0xd0, 0x27, 0x7f, 0x41, 0x8f, 0xc2, 0x9a, 0xe6,
	0x68, 0x8d, 0x8a, 0xaf, 0x47, 0x21, 0xbd, 0x80, 0x83, 0x1c, 0x23, 0xdc, 0x3c, 0x28, 0x67, 0x1a,
	0x1b, 0x70, 0xa7, 0x99, 0x00, 0xe9, 0x03, 0x98, 0xbd, 0x15, 0x06, 0x1c, 0x65, 0xbb, 0x9f, 0x2b,
	0x91, 0x43, 0x28, 0x33, 0x9c, 0xad, 0x90, 0xd7, 0xf4, 0x4d, 0x4a, 0x31, 0xd1, 0x3e, 0x58, 0xb2,
	0xc3, 0xfe, 0x61, 0xaf, 0xc0, 0x1c, 0x25, 0xe1, 0x2f, 0x84, 0x4d, 0x43, 0xc9, 0x4a, 0xfb, 0x87,
	0xaa, 0x83, 0x79, 0x8e, 0x0b, 0x54, 0x43, 0xa9, 0x07, 0xeb, 0x83, 0x25, 0x63, 0x7b, 0x3b, 0xb6,
	0x5e, 0x4b, 0xf0, 0x2f, 0xb7, 0x67, 0x9d, 0x61, 0x9f, 0xf8, 0x50, 0xcd, 0xf5, 0x8f, 0x38, 0xb2,
	0xca, 0xf7, 0xc2, 0xda, 0x67, 0x3b, 0x08, 0x11, 0x6d, 0x00, 0x95, 0xcf, 0x8e, 0x91, 0x13, 0x99,
	0x57, 0x0b, 0x6a, 0x9f, 0x6e, 0xfd, 0x2e, 0xd4, 0x46, 0x60, 0xe4, 0x7b, 0x40, 0x94, 0x00, 0x05,
	0x2d, 0xb4, 0xe9, 0x2e, 0xe4, 0x4b, 0x36, 0x7f, 0x49, 0x55, 0xb6, 0xa0, 0x2f, 0xaa, 0x6c, 0x61,
	0x11, 0x46, 0x60, 0xe4, 0xcf, 0xa5, 0xca, 0x16, 0x5c, 0x5c, 0x95, 0x2d, 0xba, 0x76, 0x77, 0x0a,
	0x47, 0x51, 0xec, 0xe2, 0x92, 0x61, 0x10, 0x06, 0x12, 0xdf, 0xb5, 0xa4, 0x93, 0x26, 0xd1, 0x30,
	0x7d, 0x6d, 0x86, 0xda, 0x5d, 0x35, 0xa5, 0x04, 0xf4, 0xac, 0x97, 0x3a, 0xe3, 0xf1, 0x8b, 0x6e,
	0xa4, 0xa8, 0x7b, 0xeb, 0x75, 0xd3, 0xe5, 0x5b, 0x36, 0xde, 0x8b, 0x71, 0x5a, 0xde, 0xbc, 0x53,
	0xed, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x3a, 0xd8, 0xd1, 0x00, 0x05, 0x00, 0x00,
}
