// Package authclient provides support to access the auth service.
package authclient

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

// This provides a default client configuration, but it's recommended
// this is replaced by the user with application specific settings using
// the WithClient function at the time a AuthAPI is constructed.
var defaultClient = http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// Logger represents a function that has user logging context.
type Logger func(ctx context.Context, msg string, v ...any)

// Client represents a client that can talk to the auth service.
type Client struct {
	url  string
	log  Logger
	http *http.Client
}

// New constructs an Auth that can be used to talk with the auth service.
func New(url string, log Logger, options ...func(cln *Client)) *Client {
	cln := Client{
		url:  url,
		log:  log,
		http: &defaultClient,
	}

	for _, option := range options {
		option(&cln)
	}

	return &cln
}

// WithClient adds a custom client for processing requests. It's recommend
// to not use the default client and provide your own.
func WithClient(http *http.Client) func(cln *Client) {
	return func(cln *Client) {
		cln.http = http
	}
}

// Authenticate calls the auth service to authenticate the user.
func (cln *Client) Authenticate(ctx context.Context, authorization string) (AuthenticateResp, error) {
	endpoint := fmt.Sprintf("%s/auth/authenticate", cln.url)

	headers := map[string]string{
		"authorization": authorization,
	}

	var resp AuthenticateResp
	if err := cln.rawRequest(ctx, http.MethodGet, endpoint, headers, nil, &resp); err != nil {
		return AuthenticateResp{}, err
	}

	return resp, nil
}
