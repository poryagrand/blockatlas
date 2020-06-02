package ens

import (
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	AddressEncoder "github.com/trustwallet/ens-coincodec"
)

type NamingProvider struct {
	client RpcClient
}

func Init(rpc string) *NamingProvider {
	return &NamingProvider{
		client: RpcClient{Request: blockatlas.InitJSONClient(rpc)},
	}
}

func (p *NamingProvider) Lookup(coins []uint64, name string) ([]blockatlas.Resolved, error) {
	var result []blockatlas.Resolved
	node, err := NameHash(name)
	if err != nil {
		return result, errors.E(err, "name hash failed")
	}
	for _, coin := range coins {
		resolver, err := p.client.Resolver(node[:])
		if err != nil {
			return result, errors.E(err, "query resolver failed")
		}
		// try to get multi coin address
		address, err := p.addressForCoin("0x"+resolver, node[:], coin)
		if err != nil {
			logger.Error(errors.E(err, errors.Params{"coin": coin, "name": name}))
			continue
		}
		result = append(result, blockatlas.Resolved{Coin: coin, Result: address})
	}

	return result, nil
}

func (p *NamingProvider) addressForCoin(resovler string, node []byte, coin uint64) (string, error) {
	result, err := p.client.Addr(resovler, node, coin)
	if err != nil {
		if coin == CoinType.ETH {
			// user may not set multi coin address
			result, err := p.lookupLegacyETH(resovler, node)
			if err != nil {
				return "", errors.E(err, "query legacy address failed")
			}
			return result, nil
		}
		return "", errors.E(err, "query multi coin address failed")
	}
	encoded, err := AddressEncoder.ToString(result, uint32(coin))
	if err != nil {
		return "", errors.E(err, "encode to address failed")
	}
	return encoded, nil
}

func (p *NamingProvider) lookupLegacyETH(resolver string, node []byte) (string, error) {
	return p.client.LegacyAddr(resolver, node)
}
