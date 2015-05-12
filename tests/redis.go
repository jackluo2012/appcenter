package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

type Stock struct {
	CompanyName string `redis:"company_name"`
	OpenPrice   string `redis:"open_price"`
	AskPrice    string `redis:"ask_price"`
	ClosePrice  string `redis:"close_price"`
	BidPrice    string `redis:"bid_price"`
}

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalf("Couldn't connect to Redis: %v\n", err)
	}
	defer conn.Close()

	stockData := map[string]*Stock{
		"GOOG": &Stock{CompanyName: "Google Inc.", OpenPrice: "803.99", AskPrice: "795.50", ClosePrice: "802.66", BidPrice: "793.36"},
		"MSFT": &Stock{AskPrice: "N/A", OpenPrice: "28.30", CompanyName: "Microsoft Corpora", BidPrice: "28.50", ClosePrice: "28.37"},
	}

	for sym, row := range stockData {
		if _, err := conn.Do("HMSET", redis.Args{sym}.AddFlat(row)...); err != nil {
			log.Fatal(err)
		}
	}

	for sym := range stockData {
		values, err := redis.Values(conn.Do("HGETALL", sym))
		if err != nil {
			log.Fatal(err)
		}
		var stock Stock
		if err := redis.ScanStruct(values, &stock); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s: %+v", sym, &stock)
	}
}
