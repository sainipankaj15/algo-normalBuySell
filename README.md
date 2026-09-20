# Algo_NormalBuySell

A simple intraday trading bot written in Go that opens a long position in one stock and a short position in another, waits until the configured close time, then squares off both positions and exits.

## Overview

This project implements a basic long/short market-open strategy:

- Long symbol: HDFCBANK
- Short symbol: SBIN
- Quantity: 1
- Start time: 09:30:02
- Close time: 09:35:02

At startup, the bot:

1. Creates a log file for the session
2. Sends a Telegram notification
3. Waits until the configured market start time
4. Places a market buy order for the long symbol
5. Places a market sell order for the short symbol
6. Waits until the configured close time
7. Squares off both positions
8. Sends a completion notification and exits

## Project Structure

- `main.go` - bot flow and execution logic
- `utils.go` - order placement and square-off functions
- `variable.go` - strategy configuration values
- `constants.go` - Telegram-related constants
- `go.mod` - Go module definition and dependency list

## Requirements

- Go 1.22+
- A valid broker integration configured through the underlying trading library
- Telegram bot access if you want notifications enabled

## Installation

```bash
go mod tidy
go run .
```

## Configuration

The strategy parameters are defined in `variable.go`:

- `LongSymbol`
- `ShortSymbol`
- `Quantity`
- `StartingHour`, `StartingMinutes`, `StartingSeconds`
- `ClosingHour`, `ClosingMinutes`, `ClosingSeconds`

The Telegram bot details are defined in `constants.go` and should not be committed in a public repository.

## Important Security Note

Never hardcode live secrets, tokens, chat IDs, or account credentials in a public Git repository.

Before publishing or sharing the code:

- remove hardcoded credentials
- move them to environment variables or a local `.env` file
- add `.env` and sensitive files to `.gitignore`
- rotate any token that was ever committed

## Notes

- This bot is intended for educational and testing purposes.
- Trading in live markets involves risk.
- The project is not financial advice.

