````markdown
# ETH Price API

A simple Go API to fetch the current price and trading volume of ETH from popular exchanges: MEXC and Binance.

## About

This service queries public APIs of MEXC and Binance exchanges and returns the current ETH price in USDT, trading volume, and the time of the request. It can be used as part of trading bots, analytics systems, or just for learning how to work with APIs and Go.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/eth-price-api.git
cd eth-price-api
```

2. Run the server:

```bash
go run .
```

The API will be available at `http://localhost:8080`

## Usage

Available endpoints:

* `/ETHPriceMEXC` — ETH price and volume on MEXC
* `/ETHPriceBinance` — ETH price and volume on Binance

Example request:

```bash
curl http://localhost:8080/ETHPriceMEXC
```

Example response:

```json
{
  "Crypto": "ETHUSDT",
  "PriceUSDT": 1800.54,
  "Exchange": "MEXC",
  "TimeNow": "Fri, 31 Jul 2025 11:00:00 UTC",
  "Volume": 12345.67
}
```

## Project Structure

* `main.go` — entry point, HTTP route setup
* `handlers.go` — API request handlers
* `structs.go` — data structures for parsing and responding

## Contributing

* Feel free to submit pull requests with improvements
* Report issues and bugs
* Add support for more cryptocurrencies and exchanges

## License

This project is licensed under the MIT License — free to use and modify.
