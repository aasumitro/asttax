package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aasumitro/asttax/internal/util/cache"
)

type ICoingeckoRepository interface {
	GetSolanaPrice(ctx context.Context) (float64, error)
}

type coingeckoRepository struct {
	apiURL    string
	cachePool *cache.Cache
}

func (repo *coingeckoRepository) GetSolanaPrice(ctx context.Context) (float64, error) {
	// get data from cache
	const cacheKey = "sol_usd_price"
	if cacheData, ok := repo.cachePool.Get(cacheKey); ok {
		if price, ok := cacheData.(float64); ok {
			return price, nil
		}
	}
	// get data from api
	url := fmt.Sprintf("%s/v3/simple/price?ids=solana&vs_currencies=usd", repo.apiURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}
	client := &http.Client{
		Timeout: ContextTimeoutDuration * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch SOL price: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected response status: %s", resp.Status)
	}
	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode response: %v", err)
	}
	priceUSD, ok := result["solana"]["usd"]
	if !ok {
		return 0, fmt.Errorf("SOL price not found in response")
	}
	// cache data and return
	expiredIn := CacheDuration * time.Second
	repo.cachePool.Set(cacheKey, priceUSD, expiredIn)
	return priceUSD, nil
}

func NewCoingeckoRepository(
	apiURL string,
	cachePool *cache.Cache,
) ICoingeckoRepository {
	return &coingeckoRepository{
		apiURL:    apiURL,
		cachePool: cachePool,
	}
}
