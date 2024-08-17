package gmx

// Wrapper for Avalanche
type GmxAvalancheClient struct {
	gmxClient
}

func NewGmxAvalancheClient() (*GmxAvalancheClient, error) {
	baseClient, err := newGmxClient(
		"gmx-v2-avalanche",
		indexerGraphqlAvalanche,
		avalancheTokensUrl,
	)

	if err != nil {
		return nil, err
	}

	return &GmxAvalancheClient{*baseClient}, nil
}

type GmxArbitrumClient struct {
	gmxClient
}

func NewGmxArbitrumClient() (*GmxArbitrumClient, error) {
	baseClient, err := newGmxClient(
		"gmx-v2-arbitrum",
		indexerGraphqlArbitrum,
		arbitrumTokensUrl,
	)

	if err != nil {
		return nil, err
	}

	return &GmxArbitrumClient{*baseClient}, nil
}
