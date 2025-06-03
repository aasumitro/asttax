package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/aasumitro/asttax/internal/util/cache"
	solanaRpcClient "github.com/blocto/solana-go-sdk/client"
)

type ISolanaRPCRepository interface {
	GetBalance(ctx context.Context, account string) (uint64, error)
}

type solanaRPCRepository struct {
	rpcClient *solanaRpcClient.Client
	cachePool *cache.Cache
}

func (repo *solanaRPCRepository) GetBalance(
	ctx context.Context,
	account string,
) (uint64, error) {
	// from cache
	cacheKey := fmt.Sprintf("%s_sol_balance", account)
	if cacheData, ok := repo.cachePool.Get(cacheKey); ok {
		if price, ok := cacheData.(uint64); ok {
			return price, nil
		}
	}
	// from net
	solBalance, err := repo.rpcClient.GetBalance(ctx, account)
	if err != nil {
		return 0.0, err
	}
	// cache data and return
	expiredIn := CacheDuration * time.Second
	repo.cachePool.Set(cacheKey, solBalance, expiredIn)
	return solBalance, nil
}

func NewSolanaRPCRepository(
	rpcClient *solanaRpcClient.Client,
	cachePool *cache.Cache,
) ISolanaRPCRepository {
	return &solanaRPCRepository{
		rpcClient: rpcClient,
		cachePool: cachePool,
	}
}
