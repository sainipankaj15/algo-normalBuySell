package main

import (
	"log"
	"os"
	"time"

	utils "github.com/sainipankaj15/All-In-One-Broker/commanUtilsAcrossBroker"
)

var AlgoName = "Algo_LongShort_5Min"

func main() {
	// Step 1 : Setting the log file name
	fileName := "Log_" + AlgoName + "_" + utils.CurrentDate() + "_" + utils.CurrentTime() + "_" + ".txt"
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("Starting algorithm:", AlgoName)
	if err := utils.TelegramSend(BOTTOKEN, CHATID, "Algo started: "+AlgoName); err != nil {
		log.Printf("Telegram send failed: %v", err)
	}

	// Step 2 : Wait for market start time
	utils.ApplicationStart(StartingHour, StartingMinutes, StartingSeconds)

	// Step 3 : Place the long and short positions
	log.Printf("Opening long position in %s and short position in %s", LongSymbol, ShortSymbol)

	if err := placeMarketOrderForBuy(LongSymbol); err != nil {
		log.Printf("Failed to place buy order for %s: %v", LongSymbol, err)
		return
	}

	if err := placeMarketOrderForSell(ShortSymbol); err != nil {
		log.Printf("Failed to place sell order for %s: %v", ShortSymbol, err)
		return
	}

	log.Println("Both positions opened successfully. Waiting for the configured close time before square off.")
	if err := utils.TelegramSend(BOTTOKEN, CHATID, "Positions opened: BUY HDFCBANK and SELL SBIN"); err != nil {
		log.Printf("Telegram send failed: %v", err)
	}

	isWorkDone := make(chan time.Time)
	go utils.ApplicationClosing(ClosingHour, ClosingMinutes, ClosingSeconds, isWorkDone)

	<-isWorkDone
	close(isWorkDone)

	log.Println("Configured close time reached. Squaring off both positions.")
	if err := utils.TelegramSend(BOTTOKEN, CHATID, "Configured close time reached. Squaring off both positions."); err != nil {
		log.Printf("Telegram send failed: %v", err)
	}

	if err := squareOffPosition(LongSymbol); err != nil {
		log.Printf("Failed to square off long position for %s: %v", LongSymbol, err)
	}

	if err := squareOffPosition(ShortSymbol); err != nil {
		log.Printf("Failed to square off short position for %s: %v", ShortSymbol, err)
	}

	log.Println("Algo completed. Both positions squared off and the application is closing.")
	if err := utils.TelegramSend(BOTTOKEN, CHATID, "Algo completed. Both positions squared off and application is closing."); err != nil {
		log.Printf("Telegram send failed: %v", err)
	}
}
