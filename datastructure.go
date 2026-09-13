package main

type TradingInfo struct {
	LongSymbol  string
	ShortSymbol string
	Quantity    int
	OrderID     string
}

type ControlBlock struct {
	IsPositionOpened bool
	IsSquareOffDone  bool
}

type profitLossInfo struct {
	TradePrice   float64
	SymbolName   string
	LtpOfSymbol  float64
	PandLPerLot  float64
	PandLOverall float64
}
