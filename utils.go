package main

import (
	"log"
	"strconv"

	zerodha "github.com/sainipankaj15/All-In-One-Broker/Zerodha"
)

func placeMarketOrderForBuy(symbol string) error {
	resp, err := zerodha.PlaceMarketOrder(
		zerodha.Exchange.NSE,
		symbol,
		strconv.Itoa(Quantity),
		zerodha.OrderType.MARKET,
		zerodha.TransactionSide.BUY,
		zerodha.ProductType.INTRADAY,
		ZerodhaUserID,
	)
	if err != nil {
		return err
	}

	log.Printf("BUY order placed successfully for %s with quantity %d. Response: %+v", symbol, Quantity, resp)
	return nil
}

func placeMarketOrderForSell(symbol string) error {
	resp, err := zerodha.PlaceMarketOrder(
		zerodha.Exchange.NSE,
		symbol,
		strconv.Itoa(Quantity),
		zerodha.OrderType.MARKET,
		zerodha.TransactionSide.SELL,
		zerodha.ProductType.INTRADAY,
		ZerodhaUserID,
	)
	if err != nil {
		return err
	}

	log.Printf("SELL order placed successfully for %s with quantity %d. Response: %+v", symbol, Quantity, resp)
	return nil
}

func squareOffPosition(symbol string) error {
	// Replace this placeholder with the real square-off / close-position API call for your broker.
	log.Printf("Square-off placeholder executed for %s", symbol)
	return nil
}
