package gmx

// Wrapper for Avalanche
type GmxAvalancheClient struct {
	gmxClient
}

func NewGmxAvalancheClient(rpcs []string, priceUpdatePeriod float64) (*GmxAvalancheClient, error) {
	baseClient, err := newGmxClient(
		"gmx-v2-avalanche",
		indexerGraphqlAvalanche,
		avalancheTokensUrl,
		avalanchePricesUrl,
		priceUpdatePeriod,
		rpcs,
		avalancheReaderAddress,
		avalancheDataStoreAddress,
	)

	if err != nil {
		return nil, err
	}

	return &GmxAvalancheClient{*baseClient}, nil
}

type GmxArbitrumClient struct {
	gmxClient
}

func NewGmxArbitrumClient(rpcs []string, priceUpdatePeriod float64) (*GmxArbitrumClient, error) {
	baseClient, err := newGmxClient(
		"gmx-v2-arbitrum",
		indexerGraphqlArbitrum,
		arbitrumTokensUrl,
		arbitrumPricesUrl,
		priceUpdatePeriod,
		rpcs,
		arbitrumReaderAddress,
		arbitrumDataStoreAddress,
	)

	if err != nil {
		return nil, err
	}

	return &GmxArbitrumClient{*baseClient}, nil
}
