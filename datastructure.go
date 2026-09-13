package main

type TradingInfo struct {

}

type ControlBlock struct {

}

type profitLossInfo struct {
	TradePrice   float64
	SymbolName   string
	LtpOfSymbol  float64
	PandLPerLot  float64 // Per lot Profit and Loss
	PandLOverall float64 // overall profit and loss with traded quantity
}


