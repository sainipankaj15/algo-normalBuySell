package main

import "time"

var UserFyersID string = "XP03754"
var ZerodhaUserID string = "FC8173"

// ################ Long / Short Equity Strategy Details : Start #####################
var LongSymbol string = "HDFCBANK"
var ShortSymbol string = "SBIN"
var Quantity int = 1
var TargetSymbolExpiryDay time.Weekday = time.Wednesday
var TargetSymbolToken int32 = 26009
var TargetSymbolStrikeGap = 100

// Timing for closing the application
var (
	ClosingHour    int = 15
	ClosingMinutes int = 10
	ClosingSeconds int = 01
)

// Timing for starting the application : Basically resume the application
var (
	StartingHour    int = 9
	StartingMinutes int = 30
	StartingSeconds int = 02
)

var (
	ExitingAllHour    int = 15
	ExitingAllMinutes int = 27
	ExitingAllSeconds int = 01
)
