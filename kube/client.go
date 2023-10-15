package kube

import (
	"net/http"
	"time"

	"github.com/rafaelmgr12/kube-client/deployment"
)

type Client struct {
	url        string
	timeout    time.Duration
	httpClient *http.Client

	Deployment deployment.Service
}

func NewClient(options ...option) (*Client, error) {
	c := Client{
		url:        "http://localhost:3000",
		httpClient: &http.Client{},
	}
	for _, option := range options {
		if err := option(&c); err != nil {
			return nil, err
		}
	}

	if c.timeout != 0 {
		c.httpClient.Timeout = c.timeout
	}

	c.Deployment = deployment.NewService(c.httpClient, c.url)

	return &c, nil
}
