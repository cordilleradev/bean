package main

import (
	"encoding/json"
	"fmt"

	"github.com/cordilleradev/bean/clients/gmx"
)

func main() {

	var arbRpcs = []string{
		"https://arbitrum.llamarpc.com",
		"https://api.zan.top/node/v1/arb/one/public",
		"https://api.stateless.solutions/arbitrum-one/v1/demo",
		"https://endpoints.omniatech.io/v1/arbitrum/one/public",
		"https://1rpc.io/arb",
		"https://arbitrum.rpc.subquery.network/public",
		"https://arbitrum.blockpi.network/v1/rpc/public",
		"https://arb-pokt.nodies.app",
		"https://arbitrum.meowrpc.com",
		"https://arbitrum.drpc.org",
		"https://arbitrum-one.public.blastapi.io",
	}

	a, err := gmx.NewGmxArbitrumClient(arbRpcs, 1)
	if err != nil {
		panic(err)
	}

	address := "0x591b6F096281DD7b645767C96aC34863A4Df9a89"
	fmt.Println("Fetching positions for address:", address)
	p, fetchErr := a.FetchPositions(address)
	if fetchErr != nil {
		fmt.Printf("Error fetching positions: %v\n", fetchErr)
		panic(fetchErr)
	}

	pJson, jsonErr := json.MarshalIndent(p, "", "  ")
	if jsonErr != nil {
		fmt.Printf("Error marshalling positions to JSON: %v\n", jsonErr)
		panic(jsonErr)
	}
	fmt.Println("Positions in JSON format:", string(pJson))

	// baseRpcs := []string{
	// 	"https://base.llamarpc.com",
	// 	"https://mainnet.base.org",
	// 	"https://developer-access-mainnet.base.org",
	// 	"https://base-rpc.publicnode.com",
	// 	"https://base.blockpi.network/v1/rpc/public",
	// 	"https://base-mainnet.public.blastapi.io",
	// 	"https://base.meowrpc.com",
	// 	"https://endpoints.omniatech.io/v1/base/mainnet/public",
	// 	"https://base.gateway.tenderly.co",
	// 	"https://gateway.tenderly.co/public/base",
	// 	"https://1rpc.io/base",
	// 	"https://base.rpc.subquery.network/public",
	// 	"https://base-pokt.nodies.app",
	// 	"https://base.api.onfinality.io/public",
	// 	"https://base.drpc.org",
	// }

	// fmt.Println("Creating Avantis client with RPCs:", baseRpcs)
	// a, clientErr := avantis.NewAvantisClient(baseRpcs)

	// if clientErr != nil {
	// 	fmt.Printf("Error creating Avantis client: %v\n", clientErr)
	// 	panic(clientErr)
	// }

	// address := "0x8E1c4e0a7e85b2490f6d811824515D6FAD3115A6"
	// fmt.Println("Fetching positions for address:", address)
	// p, fetchErr := a.FetchPositions(address)
	// if fetchErr != nil {
	// 	fmt.Printf("Error fetching positions: %v\n", fetchErr)
	// 	panic(fetchErr)
	// }

	// pJson, jsonErr := json.MarshalIndent(p, "", "  ")
	// if jsonErr != nil {
	// 	fmt.Printf("Error marshalling positions to JSON: %v\n", jsonErr)
	// 	panic(jsonErr)
	// }
	// fmt.Println("Positions in JSON format:", string(pJson))
}
