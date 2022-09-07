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
	IetfInterfacePath = "restconf/data/ietf-interfaces:interfaces"
	IetfInterfaceType = "iana-if-type:softwareLoopback"
	contentType       = "application/yang-config+json"
)

type IetfInterfaceRequest struct {
	IetfInterface IetfInterface `json:"ietf-interfaces:interface"`
}

type IetfInterface struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Enabled     bool   `json:"enabled"`
	IPv4        Ipv4   `json:"ietf-ip:ipv4"`
}

type Ipv4 struct {
	Address []Address `json:"address"`
}

type Address struct {
	Ip      string `json:"ip"`
	Netmask string `json:"netmask"`
}

type Client struct {
	httpClient  *http.Client
	restconfURL string
	username    string
	password    string
}

func New(timeout time.Duration, restconfURL string, username string, password string) *Client {
	return &Client{
		httpClient:  &http.Client{Timeout: timeout},
		restconfURL: restconfURL,
		username:    username,
		password:    password,
	}
}

func (c *Client) NewLoopbackInterface(ctx context.Context, req IetfInterfaceRequest) (*http.Response, error) {
	config, err := json.Marshal(req)
	if err != nil {
		return &http.Response{}, fmt.Errorf("failed to encode config %w", err)
	}

	request, err := c.formRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s", c.restconfURL, IetfInterfacePath), config)
	if err != nil {
		return &http.Response{}, err
	}

	response, err := c.sendRequest(request)
	if err != nil {
		return &http.Response{}, err
	}
	return response, nil
}

func (c *Client) formRequest(ctx context.Context, method string, URI string, config []byte) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, URI, bytes.NewBuffer(config))
	if err != nil {
		return nil, fmt.Errorf("failed to form restconf request %w", err)
	}
	request.SetBasicAuth(c.username, c.password)
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", contentType)
	return request, nil
}

func (c *Client) sendRequest(req *http.Request) (*http.Response, error) {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send restconf request %w", err)
	}
	return response, nil
}
