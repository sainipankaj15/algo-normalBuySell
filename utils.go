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
	var side string

	switch symbol {
	case LongSymbol:
		side = zerodha.TransactionSide.SELL
	case ShortSymbol:
		side = zerodha.TransactionSide.BUY
	}

	resp, err := zerodha.PlaceMarketOrder(
		zerodha.Exchange.NSE,
		symbol,
		strconv.Itoa(Quantity),
		zerodha.OrderType.MARKET,
		side,
		zerodha.ProductType.INTRADAY,
		ZerodhaUserID,
	)
	if err != nil {
		return err
	}

	log.Printf("Square-off executed for %s with side %s. Response: %+v", symbol, side, resp)
	return nil
}
