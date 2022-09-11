// Package restconf contains methods for configuring Routers via RESTCONF API.
package restconf

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	IetfInterfacesPath = "restconf/data/ietf-interfaces:interfaces"
	IetfInterfaceType  = "iana-if-type:softwareLoopback"
	contentType        = "application/yang-config+json"
)

type Client struct {
	httpClient *http.Client
}

type Server struct {
	URL      string
	username string
	password string
}

// New creates a HTTP client with timeout.
func New(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
	}
}

// NewLoopbackInterface creates a loopback interface on a Router via RESTCONF API.
func (c *Client) NewLoopbackInterface(ctx context.Context, config IetfInterfaceRequest, server Server) (*http.Response, error) {
	payload, err := json.Marshal(config)
	if err != nil {
		return &http.Response{}, fmt.Errorf("failed to encode config %w", err)
	}

	request, err := c.formRequest(ctx, http.MethodPost, IetfInterfacesPath, &payload, &server)
	if err != nil {
		return &http.Response{}, err
	}

	response, err := c.sendRequest(request)
	if err != nil {
		return &http.Response{}, err
	}
	return response, nil
}

func (c *Client) formRequest(ctx context.Context, method string, path string, config *[]byte, server *Server) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", server.URL, path), bytes.NewBuffer(*config))
	if err != nil {
		return nil, fmt.Errorf("failed to form RESTCONF request %w", err)
	}
	request.SetBasicAuth(server.username, server.password)
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", contentType)
	return request, nil
}

func (c *Client) sendRequest(req *http.Request) (*http.Response, error) {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send RESTCONF request %w", err)
	}
	return response, nil
}
