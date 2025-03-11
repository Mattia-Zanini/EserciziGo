package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// creo i canali
	eur_usd := make(chan float64)
	gbp_usd := make(chan float64)
	jpy_usd := make(chan float64)

	go simulateMarketData(eur_usd, "EUR/USD")
	go simulateMarketData(gbp_usd, "GBP/USD")
	go simulateMarketData(jpy_usd, "JPY/USD")

	go selectPair(eur_usd, gbp_usd, jpy_usd)

	time.Sleep(time.Minute)
}

func simulateMarketData(c chan<- float64, cambioValuta string) {
	min := 1.0
	max := 1.5
	if cambioValuta == "JPY/USD" {
		min = 0.006
		max = 0.009
	}

	for {
		newValue := min + rand.Float64()*(max-min)
		c <- newValue
		time.Sleep(time.Second)
	}
}

func selectPair(eu, gu, ju <-chan float64) {
	for {
		select {
		case euValue := <-eu:
			if euValue > 1.20 {
				fmt.Println("Vendita di [EUR/USD] a", euValue)
				time.Sleep(time.Second * 4)
				fmt.Println("Venduto [EUR/USD]")
			}
		case guValue := <-gu:
			if guValue < 1.35 {
				fmt.Println("Acquisto di [GBP/USD] a", guValue)
				time.Sleep(time.Second * 3)
				fmt.Println("Acquistato [GBP/USD]")
			}
		case juValue := <-ju:
			if juValue < 0.0085 {
				fmt.Println("Acquisto di [JPY/USD] a", juValue)
				time.Sleep(time.Second * 3)
				fmt.Println("Acquistato [JPY/USD]")
			}
		}
	}
}
