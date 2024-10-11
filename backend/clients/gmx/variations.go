package gmx

// Wrapper for Avalanche
type GmxAvalancheClient struct {
	*gmxClient
}

func NewGmxAvalancheClient(rpcs []string) (*GmxAvalancheClient, error) {
	baseClient, err := newGmxClient(
		"gmx-v2-avalanche",
		indexerGraphqlAvalanche,
		avalancheTokensUrl,
		avalanchePricesUrl,
		rpcs,
		avalancheReaderAddress,
		avalancheDataStoreAddress,
	)

	if err != nil {
		return nil, err
	}

	return &GmxAvalancheClient{baseClient}, nil
}

type GmxArbitrumClient struct {
	*gmxClient
}

func NewGmxArbitrumClient(rpcs []string) (*GmxArbitrumClient, error) {
	baseClient, err := newGmxClient(
		"gmx-v2-arbitrum",
		indexerGraphqlArbitrum,
		arbitrumTokensUrl,
		arbitrumPricesUrl,
		rpcs,
		arbitrumReaderAddress,
		arbitrumDataStoreAddress,
	)

	if err != nil {
		return nil, err
	}

	return &GmxArbitrumClient{baseClient}, nil
}
