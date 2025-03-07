package main

import (
	"fmt"
	"time"
)

const (
	COTTA = iota
	GUARNITA
	DECORATA
)

func printTempo() {
	secondi := 1
	for {
		time.Sleep(time.Second)
		fmt.Printf("Sono passati %v s\n", secondi)
		secondi++
	}
}

type torta struct {
	id    int
	stato int
}

var cotteCh = make(chan torta, 2)    // canale con buffer per le torte cotte da guarnire
var guarniteCh = make(chan torta, 2) // canale con buffer per le torte guarnite da decorare
var finito = make(chan struct{})     // canale per segnalare che sono state finite tutte le torte

func main() {
	// PRIMO pasticciere
	go func(c chan<- torta) {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)

			id_torta := i + 1
			c <- torta{id: id_torta, stato: COTTA}
			fmt.Printf("CUCINATA la torta [%v]\n", id_torta)
		}
		close(cotteCh)
	}(cotteCh)

	// SECONDO pasticciere
	go func(c <-chan torta, g chan<- torta) {
		for tortaSingola := range c {
			time.Sleep(time.Second * 4)

			tortaSingola.stato = GUARNITA
			g <- tortaSingola
			fmt.Printf("GUARNITA la torta [%v]\n", tortaSingola.id)
		}
		close(guarniteCh)
	}(cotteCh, guarniteCh)

	// TERZO pasticciere
	go func(g <-chan torta, f chan<- struct{}) {
		for tortaSingola := range g {
			time.Sleep(time.Second * 8)

			fmt.Printf("DECORATA la torta [%v]\n", tortaSingola.id)
		}
		f <- struct{}{} // segnalo che il terzo pasticciere ha finito di decorare tutte le torte
		close(finito)
	}(guarniteCh, finito)
	go printTempo() // stampa i secondi passati

	// aspetto che il terzo pasticciere finisca
	// Il select blocca l'esecuzione fino a quando non riceve un messaggio dai canali monitorati.
	select {
	case <-finito:
		fmt.Println("Tutti i pasticcieri hanno finito")
		break
	}
}
