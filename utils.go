package main

import (
	"log"
	"time"

	"github.com/markcheno/go-talib"
)

func roundToNearest15Minutes(t time.Time) time.Time {
	rounded := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), (t.Minute()/15)*15, 0, 0, t.Location())
	return rounded
}

func roundToNearest15MinutesFromCurrentTime() int64 {

	// Set the location to Indian Standard Time (IST)
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Println("Error loading IST location:", err)
		return -1
	}

	// Get the current time in IST
	currentTime := time.Now().In(ist)

	// Round down to the nearest 15-minute interval
	roundedTime := roundToNearest15Minutes(currentTime)

	epochTime := roundedTime.Unix()
	return epochTime
}

func rsi(px []float64, lookback int) float64 {
	rsi := talib.Rsi(px, lookback)
	return rsi[len(rsi)-1]
}

func ema(px []float64, lookback int) float64 {
	ema := talib.Ema(px, lookback)
	return ema[len(ema)-1]
}

func placeMarketOrderForBuy(symbol string) error {
	// Replace this placeholder with the actual Zerodha API call once the Zerodha package is available.
	// Example:
	// resp, err := zerodha.PlaceMarketOrder(
	// 	zerodha.Exchange.NSE,
	// 	symbol,
	// 	"1",
	// 	zerodha.OrderType.MARKET,
	// 	zerodha.TransactionSide.BUY,
	// 	zerodha.ProductType.INTRADAY,
	// 	"FC8173",
	// )
	// if err != nil {
	// 	return err
	// }
	// log.Println("Buy order placed successfully for", symbol, "Response:", resp)

	log.Printf("BUY placeholder executed for %s with quantity %d", symbol, Quantity)
	return nil
}

func placeMarketOrderForSell(symbol string) error {
	// Replace this placeholder with the actual Zerodha API call once the Zerodha package is available.
	// Example:
	// resp, err := zerodha.PlaceMarketOrder(
	// 	zerodha.Exchange.NSE,
	// 	symbol,
	// 	"1",
	// 	zerodha.OrderType.MARKET,
	// 	zerodha.TransactionSide.SELL,
	// 	zerodha.ProductType.INTRADAY,
	// 	"FC8173",
	// )
	// if err != nil {
	// 	return err
	// }
	// log.Println("Sell order placed successfully for", symbol, "Response:", resp)

	log.Printf("SELL placeholder executed for %s with quantity %d", symbol, Quantity)
	return nil
}

func squareOffPosition(symbol string) error {
	// Replace this placeholder with the real square-off / close-position API call for your broker.
	log.Printf("Square-off placeholder executed for %s", symbol)
	return nil
}
