package main

import (
	"github.com/cordilleradev/bean/clients/avantis"
	"github.com/cordilleradev/bean/clients/gains"
	"github.com/cordilleradev/bean/clients/gmx"
	"github.com/cordilleradev/bean/clients/hyperliquid"
	"github.com/cordilleradev/bean/clients/kwenta"
)

func main() {
	gmxAr, _ := gmx.NewGmxArbitrumClient()
	gmxAx, _ := gmx.NewGmxAvalancheClient()
	hyp, _ := hyperliquid.NewHyperliquidClient()
	kweB, _ := kwenta.NewKwentaBaseClient()
	kweO, _ := kwenta.NewKwentaOptimismClient()
	gain, _ := gains.NewGainsClient()
	avan, _ := avantis.NewAvantisClient()

}
