package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aasumitro/asttax/internal/config"
)

type ICoingeckoRepository interface {
	GetSolanaPrice(ctx context.Context) (float64, error)
}

type coingeckoRepository struct {
	apiURL string
}

func (repo *coingeckoRepository) GetSolanaPrice(ctx context.Context) (float64, error) {
	// get data from cache
	const cacheKey = "sol_usd_price"
	if cacheData, ok := config.CachePool.Get(cacheKey); ok {
		if price, ok := cacheData.(float64); ok {
			return price, nil
		}
	}
	// get data from api
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, repo.apiURL, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}
	timeOutDur := 10
	client := &http.Client{
		Timeout: time.Duration(timeOutDur) * time.Second, // Set a reasonable timeout
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
	cacheDurTime := 20
	expiredIn := time.Duration(cacheDurTime) * time.Second
	config.CachePool.Set(cacheKey, priceUSD, expiredIn)
	return priceUSD, nil
}

func NewCoingeckoRepository(apiURL string) ICoingeckoRepository {
	return &coingeckoRepository{apiURL: apiURL}
}
