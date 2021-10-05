package proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/adamjq/http-reverse-proxy/internal/config"
)

func getMockServer() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {},
		),
	)
}

func TestProxy_RoutingSucceeds(t *testing.T) {
	assert := require.New(t)

	testServerA := getMockServer()
	defer testServerA.Close()

	testServerB := getMockServer()
	defer testServerB.Close()

	ps := NewProxyServer(
		&config.Config{
			Port: "5000",
			ServiceAUrl: testServerA.URL,
			ServiceBUrl: testServerB.URL,
		},
	)
	handler := ps.proxyHandler()
	w := httptest.NewRecorder()

	url := "http://localhost:5000/service_a/"

	r, err := http.NewRequest("GET", url, nil)
	assert.NoError(err)

	handler.ServeHTTP(w, r)
	assert.Equal(http.StatusOK, w.Code)
}