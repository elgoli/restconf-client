package restconf

import (
	"net/http"
	"time"
)

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
