package kube

import (
	"net/http"
	"time"
)

type ClientOption func(*Client)

func WithURL(url string) ClientOption {
	return func(c *Client) {
		c.url = url
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
