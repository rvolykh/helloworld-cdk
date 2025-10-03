package tools

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetExternalIP_Integration(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("127.0.0.1"))
	}))
	defer testServer.Close()

	ip, err := GetExternalIP(&GetExternalIPParams{URL: testServer.URL})

	require.NoError(t, err)
	require.Equal(t, ip, "127.0.0.1")
}

func TestMustGetExternalIP_Integration(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Unit Test"))
	}))
	defer testServer.Close()

	require.Panics(t, func() {
		MustGetExternalIP(&GetExternalIPParams{URL: testServer.URL})
	})
}
