package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hariprathap-hp/backend_masterclass/token"
)

func TestAuthMiddleWare(t *testing.T) {
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		checkResponse func(t *testing.T, resp *httptest.ResponseRecorder)
	}{}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, nil)
			server.router.GET()
		})
	}
}
