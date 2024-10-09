package main

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/cordilleradev/bean/api"
	"github.com/cordilleradev/bean/clients/gmx"
	"github.com/cordilleradev/bean/clients/hyperliquid"
	"github.com/cordilleradev/bean/common"
	"github.com/cordilleradev/bean/common/utils"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatal("Atleast two commands ('local', 'prod') & ('api','worker','cron') required.")
		return
	}

	if !utils.IsStringListSet(args) {
		log.Fatal("Duplicate commands detected")
	}

	isProd := utils.ContainsString(args, "prod")
	if isProd {
		log.Println("Running production settings")
	} else {
		log.Println("Running local/staging settings")
	}

	if !isProd {
		godotenv.Load()
	}

	var initErr error
	var gmxArbClient common.FuturesClient
	var gmxAvaxClient common.FuturesClient
	var hyperliquidClient common.FuturesClient

	arbRpcs := strings.Split(
		os.Getenv("ARBITRUM_RPCS"),
		",",
	)

	avaxRpcs := strings.Split(
		os.Getenv("AVAX_RPCS"),
		",",
	)

	if os.Getenv("ENABLE_GMX_ARB") == "true" {
		if len(arbRpcs) > 0 {
			gmxArbClient, initErr = gmx.NewGmxArbitrumClient(arbRpcs)
			if initErr != nil {
				log.Fatalf("Failed to Init GMX_ARB: %s\n", initErr.Error())
			}
			log.Println("Successfully initialized GMX_ARB")
		} else {
			log.Fatal("Failed to Init GMX_ARB: No RPCs")
		}
	} else {
		log.Println("Skipping GMX_ARB, ENABLE_GMX_ARB is not true")
	}

	if os.Getenv("ENABLE_GMX_AVAX") == "true" {
		if len(avaxRpcs) > 0 {
			gmxAvaxClient, initErr = gmx.NewGmxAvalancheClient(avaxRpcs)
			if initErr != nil {
				log.Fatalf("Failed to Init GMX_AVAX: %s\n", initErr.Error())
			}
			log.Println("Successfully initialized GMX_AVAX")
		} else {
			log.Fatal("Failed to Init GMX_AVAX: No RPCs")
		}
	} else {
		log.Println("Skipping GMX_ARB, ENABLE_GMX_ARB is not true")
	}

	if os.Getenv("ENABLE_HYPERLIQUID") == "true" {
		cacheLeaderboard := os.Getenv("ENABLE_LEADERBOARD_CACHING") == "true" || utils.ContainsString(args, "worker")
		hyperliquidClient, initErr = hyperliquid.NewHyperLiquidClient(cacheLeaderboard)
		if initErr != nil {
			log.Fatalf("Failed to Init HYPERLIQUID: %s\n", initErr.Error())
		}
		log.Println("Successfully initialized HYPERLIQUID")
		log.Printf("HYPERLIQUID leaderboard caching enabled: %t\n", cacheLeaderboard)
	} else {
		log.Println("Skipping HYPERLIQUID, ENABLE_HYPERLIQUID is not true")
	}

	clientMap := make(map[string]common.FuturesClient)

	for _, client := range []common.FuturesClient{
		gmxArbClient,
		gmxAvaxClient,
		hyperliquidClient,
	} {
		if client != nil {
			clientMap[client.ExchangeName()] = client
		}
	}

	if len(clientMap) > 0 {
		var wg sync.WaitGroup
		for _, command := range args {
			switch command {
			case "api":
				log.Println("Starting API...")
				wg.Add(1)
				go api.NewApiInstance(&clientMap).Run(isProd)

			case "worker":
				wg.Add(1)
				log.Println("Starting Worker...")

			case "cron":
				wg.Add(1)
				log.Println("Starting Cron...")
			}
		}
		wg.Wait()
	} else {
		log.Fatal("No clients initialized, shutting down")
	}
}
