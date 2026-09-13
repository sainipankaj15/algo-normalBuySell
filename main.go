package main

import (
	"log"
	"os"
	"time"

	utils "github.com/sainipankaj15/All-In-One-Broker/commanUtilsAcrossBroker"
)

var AlgoName = "Algo_EmaAndRsi"

func main() {

	// Step 1 : Setting the log file name
	fileName := "Log_" + AlgoName + "_" + TargetSymbol + "_" + utils.CurrentDate() + "_" + utils.CurrentTime() + "_" + ".txt"
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Step 2 : Setup your broker websocket if required
	// if err := connectFyersWebSocket(); err != nil {
	// 	log.Printf("Fyers websocket connection failed: %v", err)
	// } else {
	// 	log.Println("Fyers websocket connected")
	// }

	// Step 3 : OptionChainMap fetch : Symbol Fetch :
	// OptionChainMap, err = fyers.GetOptionChainMap_Fyers(TargetSymbol, 20, fyers.ADMIN_FYERS)
	// if err != nil {
	// 	log.Fatal("Error while getting the option chain map", err)
	// }
	// fyers.PrintOptionChainMap(OptionChainMap)

	// Step 4 : Setup everything which you want to do before market start

	// Step 5 : Everything is Ready Now, will wait for my application Starting
	utils.ApplicationStart(StartingHour, StartingMinutes, StartingSeconds)

	// Step 6 : Now do everything which you want to do after market start Such as live LTP fetch, Algo logic etc
	// go liveLTPGoRoutine()

	// Step 9 : Main logic of Algo
	//go algo()

	// Step 6 : Create a channel for tracking when to close this program(Main Program)
	isWorkDone := make(chan time.Time)
	go utils.ApplicationClosing(ClosingHour, ClosingMinutes, ClosingSeconds, isWorkDone)

	// Step 11 : Holding there to close the program
	<-isWorkDone
	close(isWorkDone)

	// Step 10 : Function who will close all open position before closing : In case if Target and SL not hits
	//exitAllPoistionAtClosing()

	// Step 12 : Do the Stuff which you want to do before closing the application such as closing the redis connection, closing the broker websocket connection etc

}

func init() {

	// Do whatever you want to do before main function such as redis connection, Loading Env variable etc
}
