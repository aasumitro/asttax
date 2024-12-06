package rest_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/aasumitro/asttax/internal/config"
	"github.com/aasumitro/asttax/internal/repository/rest"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type testcase struct {
	name        string
	response    string
	status      int
	expected    map[string]interface{}
	method      string
	wantErr     bool
	fromCache   bool
	setupServer func(*httptest.Server, string)
}

type coingeckoRESTRepositoryTestSuite struct {
	suite.Suite
	testcases []testcase
}

func (s *coingeckoRESTRepositoryTestSuite) SetupSuite() {
	s.testcases = []testcase{
		{
			name:     "valid request",
			response: `{"solana": {"usd": 239.47}}`,
			status:   200,
			method:   http.MethodGet,
			expected: map[string]interface{}{"solana": []map[string]interface{}{{"usd": 239.47}}},
			wantErr:  false,
		},
		{
			expected:  map[string]interface{}{"solana": []map[string]interface{}{{"usd": 239.47}}},
			fromCache: true,
			wantErr:   false,
		},
		{
			name:     "no price found",
			response: `{"solana": {}}`,
			status:   200,
			method:   http.MethodGet,
			expected: map[string]interface{}{"solana": []map[string]interface{}{{"usd": 239.47}}},
			wantErr:  true,
		},
		{
			method:   http.MethodGet,
			name:     "invalid JSON response",
			response: `12345`,
			status:   200,
			wantErr:  true,
		},
		{
			name:     "invalid request",
			response: `{"error": true}`,
			status:   500,
			method:   http.MethodGet,
			expected: nil,
			wantErr:  true,
		},
		{
			method:   "lorem_ipsum",
			name:     "invalid method",
			response: `{"foo": "bar`,
			wantErr:  true,
		},
	}
}

func (s *coingeckoRESTRepositoryTestSuite) Test_GetSolanaPrice() {
	viper.SetConfigFile("../../../.env")
	cfg := config.LoadWith(context.TODO(), config.InMemoryCache())
	for _, ts := range s.testcases {
		s.T().Run(ts.name, func(t *testing.T) {
			var server *httptest.Server
			if ts.setupServer != nil {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					ts.setupServer(server, ts.response)
				}))
			} else {
				server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(ts.status)
					_, _ = w.Write([]byte(ts.response))
				}))
			}
			defer server.Close()

			if ts.fromCache {
				cfg.CachePool.Set("sol_usd_price", 239.47, 10*time.Second)
			}

			api := rest.NewCoingeckoRepository(server.URL, cfg.CachePool)
			data, err := api.GetSolanaPrice(context.Background())
			if (err != nil) != ts.wantErr {
				t.Errorf("unexpected error: %v", err)
			}
			if err == nil {
				if !reflect.DeepEqual(data, 239.47) {
					t.Errorf("expected list %v, got %v", 239.47, data)
				}
			}

			if ts.fromCache {
				cfg.CachePool.Delete("sol_usd_price")
			}
		})
	}
}

func TestCoingeckoRESTRepository(t *testing.T) {
	suite.Run(t, new(coingeckoRESTRepositoryTestSuite))
}
