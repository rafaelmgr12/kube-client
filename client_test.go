package kube

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultClient(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Errorf("should not fail to create client: %s", err)
	}

	if c.httpClient == nil {
		t.Error("expected httpClient to be set")
	}

}

func TestWithURL(t *testing.T) {
	c, err := NewClient(WithURL("http://localhost:8080"))
	if err != nil {
		t.Errorf("should not fail to create client: %s", err)
	}
	if c.url != "http://localhost:8080" {
		t.Errorf("expected url to be http://localhost:8080, got %s", c.url)
	}
}

func TestNewClient(t *testing.T) {
	// Cria um servidor HTTP de teste
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Testa a função NewClient
	client, err := NewClient(
		WithURL(server.URL),
		WithTimeout(10*time.Second),
	)

	if err != nil {
		t.Errorf("should not fail to create client: %s", err)
	}

	assert.NotNil(t, client, "Expected client to be not nil")
	assert.Equal(t, server.URL, client.url, "Expected client URL to be set to test server URL")
	assert.Equal(t, 10*time.Second, client.timeout, "Expected client timeout to be set to 10 seconds")
}

func TestClientOptions(t *testing.T) {
	client := &Client{}

	// Testa a opção WithURL
	err := WithURL("https://example.com")(client)
	assert.NoError(t, err, "Expected no error from WithURL")
	assert.Equal(t, "https://example.com", client.url, "Expected URL to be set by WithURL")

	// Testa a opção WithTimeout
	err = WithTimeout(5 * time.Second)(client)
	assert.NoError(t, err, "Expected no error from WithTimeout")
	assert.Equal(t, 5*time.Second, client.timeout, "Expected timeout to be set by WithTimeout")

	// Testa a opção WithHTTPClient
	httpClient := &http.Client{}
	err = WithHTTPClient(httpClient)(client)
	assert.NoError(t, err, "Expected no error from WithHTTPClient")
	assert.Equal(t, httpClient, client.httpClient, "Expected HttpClient to be set by WithHTTPClient")
}
