package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cordilleradev/bean/clients/gmx"
	"github.com/cordilleradev/bean/clients/hyperliquid"
	"github.com/cordilleradev/bean/common"
	"github.com/cordilleradev/bean/common/types"
)

func main() {
	arbRpcs := []string{
		"https://arbitrum.llamarpc.com",                        // Block Height: 261359389
		"https://arbitrum.rpc.subquery.network/public",         // Block Height: 261359394
		"https://arbitrum.blockpi.network/v1/rpc/public",       // Block Height: 261359395
		"https://arb-mainnet.g.alchemy.com/v2/demo",            // Block Height: 261359396
		"https://arbitrum.meowrpc.com",                         // Block Height: 261359396
		"https://arb-pokt.nodies.app",                          // Block Height: 261359398
		"https://arbitrum-one.publicnode.com",                  // Block Height: 261359398
		"https://arbitrum.gateway.tenderly.co",                 // Block Height: 261359398
		"https://arbitrum-one.public.blastapi.io",              // Block Height: 261359400
		"https://arb-mainnet-public.unifra.io",                 // Block Height: 261359401
		"https://arbitrum.drpc.org",                            // Block Height: 261359401
		"https://arbitrum-one-rpc.publicnode.com",              // Block Height: 261359402
		"https://arb1.arbitrum.io/rpc",                         // Block Height: 261359400
		"https://1rpc.io/arb",                                  // Block Height: 261359403
		"https://rpc.ankr.com/arbitrum",                        // Block Height: 261359398
		"https://api.stateless.solutions/arbitrum-one/v1/demo", // Block Height: 257061350
	}

	avaxRpcs := []string{
		"https://ava-mainnet.public.blastapi.io/ext/bc/C/rpc",   // Block Height: 51481553
		"https://rpc.ankr.com/avalanche",                        // Block Height: 51481552
		"https://avax-pokt.nodies.app/ext/bc/C/rpc",             // Block Height: 51481553
		"https://avalanche.blockpi.network/v1/rpc/public",       // Block Height: 51481553
		"https://avalanche.drpc.org",                            // Block Height: 51481554
		"https://1rpc.io/avax/c",                                // Block Height: 51481553
		"https://avalanche-mainnet.gateway.tenderly.co",         // Block Height: 51481553
		"https://avax.meowrpc.com",                              // Block Height: 51481554
		"https://avalanche-c-chain-rpc.publicnode.com",          // Block Height: 51481555
		"https://endpoints.omniatech.io/v1/avax/mainnet/public", // Block Height: 51481555
		"https://api.avax.network/ext/bc/C/rpc",                 // Block Height: 51481554
		"https://avalanche.public-rpc.com",                      // Block Height: 51481554
	}

	hlClient, err := hyperliquid.NewHyperLiquidClient(false)
	if err != nil {
		panic(err)
	}

	gmxArbClient, err := gmx.NewGmxArbitrumClient(arbRpcs, 10)
	if err != nil {
		panic(err)
	}

	gmxAvaxClient, err := gmx.NewGmxAvalancheClient(avaxRpcs, 10)
	if err != nil {
		panic(err)
	}

	type response struct {
		MarginTypes                    []string             `json:"marginTypes"`
		FixedLeaderboardPeriods        []string             `json:"leaderboardPeriods"`
		ConfigurableLeaderboardPeriods *types.CustomPeriods `json:"configurableLeaderboardPeriods,omitempty"`
		LeaderboardFields              []string             `json:"leaderboardFields"`
	}

	infoResponse := make(map[string]response)

	startTime := time.Now()

	for _, i := range []common.FuturesClient{hlClient, gmxArbClient, gmxAvaxClient} {
		marginTypes := i.GetSupportedMarginTypes()
		leaderboardFields := i.GetSupportedLeaderboardFields()

		marginTypesStr := make([]string, len(marginTypes))
		for j, mt := range marginTypes {
			marginTypesStr[j] = string(mt)
		}

		leaderboardFieldsStr := make([]string, len(leaderboardFields))
		for j, lf := range leaderboardFields {
			leaderboardFieldsStr[j] = string(lf)
		}

		leaderboardPeriods := i.GetLeaderboardPeriods()
		infoResponse[i.ExchangeName()] = response{
			MarginTypes:                    marginTypesStr,
			FixedLeaderboardPeriods:        leaderboardPeriods.FixedPeriods,
			ConfigurableLeaderboardPeriods: leaderboardPeriods.CustomPeriods,
			LeaderboardFields:              leaderboardFieldsStr,
		}
	}

	duration := time.Since(startTime)

	jsonData, err := json.MarshalIndent(infoResponse, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
	fmt.Printf("Time taken: %v\n", duration)
}
