package restconf

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	agent := newRestconfAgent()
	defer agent.Close()

	expectedClient := &Client{
		httpClient:  &http.Client{Timeout: 30 * time.Second},
		restconfURL: agent.URL,
		username:    "username",
		password:    "password",
	}
	actualClient := New(30*time.Second, agent.URL, expectedClient.username, expectedClient.password)

	require.Equal(t, expectedClient, actualClient)
}

func newRestconfAgent() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
}
