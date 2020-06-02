package naming

import (
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type (
	NamingProvider struct {
		client Client
	}

	ZNSResponse struct {
		Addresses map[string]string
	}
)

func Init(client string) *NamingProvider {
	return &NamingProvider{
		client: Client{blockatlas.InitClient(client)},
	}
}

func (p *NamingProvider) Lookup(coins []uint64, name string) ([]blockatlas.Resolved, error) {
	var result []blockatlas.Resolved
	resp, err := p.client.LookupName(name)
	if err != nil {
		return result, err
	}
	for _, coin := range coins {
		symbol := CoinType.Coins[uint(coin)].Symbol
		address := resp.Addresses[symbol]
		if len(address) == 0 {
			continue
		}
		result = append(result, blockatlas.Resolved{Coin: coin, Result: address})
	}
	return result, nil
}
