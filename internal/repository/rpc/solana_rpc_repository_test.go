package rpc_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/aasumitro/asttax/internal/config"
	rpcRepo "github.com/aasumitro/asttax/internal/repository/rpc"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type solanaRPCRepositoryTestSuite struct {
	suite.Suite
}

func (s *solanaRPCRepositoryTestSuite) SetupSuite() {}

type Param struct {
	Name          string
	RequestBody   string
	ResponseBody  string
	F             func(url string) (any, error)
	ExpectedValue any
	ExpectedError error
}

func DoTest(t *testing.T, param Param) {
	// setup test server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check request body match
		body, err := io.ReadAll(req.Body)
		assert.Nil(t, err)
		assert.JSONEq(t, param.RequestBody, string(body))
		// check write response body success
		n, err := rw.Write([]byte(param.ResponseBody))
		assert.Nil(t, err)
		assert.Equal(t, len([]byte(param.ResponseBody)), n)
	}))
	// test function
	got, err := param.F(server.URL)
	assert.Equal(t, param.ExpectedValue, got)
	assert.Equal(t, param.ExpectedError, err)
	server.Close()
}

func (s *solanaRPCRepositoryTestSuite) Test_GetBalance() {
	viper.SetConfigFile("../../../.env")
	cfg := config.LoadWith(context.TODO(), config.InMemoryCache())
	acc := "CvRuXXptXE6itGCvMxPWDnc2UYfSGKszWS14wvsK8CzK"
	ctx := context.TODO()
	balance := uint64(2039280)

	// success from rpc
	DoTest(s.T(), Param{
		RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getBalance", "params":["CvRuXXptXE6itGCvMxPWDnc2UYfSGKszWS14wvsK8CzK"]}`,
		ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"apiVersion":"1.14.10","slot":187552526},"value":2039280},"id":1}`,
		F: func(url string) (any, error) {
			c := client.NewClient(url)
			srv := rpcRepo.NewSolanaRPCRepository(c, cfg.CachePool)
			b, err := srv.GetBalance(ctx, acc)
			s.Nil(err)
			s.NotZero(b)
			return c.GetBalance(ctx, acc)
		},
		ExpectedValue: balance,
		ExpectedError: nil,
	})

	// success from cache
	cacheKey := fmt.Sprintf("%s_sol_balance", acc)
	cfg.CachePool.Set(cacheKey, balance, 10*time.Second)
	srv := rpcRepo.NewSolanaRPCRepository(nil, cfg.CachePool)
	b, err := srv.GetBalance(ctx, acc)
	s.Nil(err)
	s.NotZero(b)
	cfg.CachePool.Delete(cacheKey)
}

func TestSolanaRPCRepository(t *testing.T) {
	suite.Run(t, new(solanaRPCRepositoryTestSuite))
}
