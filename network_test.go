package network_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/benitogf/network"
	"github.com/stretchr/testify/require"
)

func TestValidIP(t *testing.T) {
	require.True(t, network.IsValidIP("127.0.0.1"))
}

func TestInvalidIP(t *testing.T) {
	require.False(t, network.IsValidIP("999.999.999.999"))
}

func TestIsHostReachable_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	u, err := url.Parse(srv.URL)
	require.NoError(t, err)

	client := network.NewHttpClient()
	reachable := network.IsHostReachable(client, u.Host)
	require.True(t, reachable)
}

type errorRoundTripper struct{}

func (e errorRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errors.New("forced error")
}

func TestIsHostReachable_Error(t *testing.T) {
	client := &http.Client{Transport: errorRoundTripper{}}
	reachable := network.IsHostReachable(client, "example.invalid")
	require.False(t, reachable)
}

func TestNewHttpClient_Config(t *testing.T) {
	client := network.NewHttpClient()
	require.NotNil(t, client)
	require.Equal(t, 5*time.Second, client.Timeout)

	transport, ok := client.Transport.(*http.Transport)
	require.True(t, ok)
	require.NotNil(t, transport)
}

func TestNewFastHttpClient_Config(t *testing.T) {
	client := network.NewFastHttpClient()
	require.NotNil(t, client)
	require.Equal(t, 1*time.Second, client.Timeout)

	transport, ok := client.Transport.(*http.Transport)
	require.True(t, ok)
	require.NotNil(t, transport)
}
