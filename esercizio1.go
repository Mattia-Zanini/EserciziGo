package main

import (
	"fmt"
	"sync"
)

func main() {
	// Creazione di un WaitGroup per sincronizzare la goroutine
	wg := sync.WaitGroup{}
	// Creazione di un canale per comunicare il risultato dalla goroutine
	ch := make(chan int)

	// Stringa da analizzare
	stringa := "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"
	// Carattere da contare nella stringa
	carattere := 'a'
	// Variabile per memorizzare il conteggio finale
	counter := 0

	// Aggiunta di un task al WaitGroup
	wg.Add(1)
	// Avvio di una goroutine per contare le occorrenze del carattere
	go func(s string, r rune, ch chan<- int) {
		count := 0
		// Iterazione su ogni carattere della stringa
		for i := range s {
			// Confronto del carattere corrente con quello cercato
			if rune(s[i]) == r {
				count++
			}
		}
		// Invio del risultato attraverso il canale
		ch <- count
		// Chiusura del canale dopo l'invio
		close(ch)
		// Segnalazione del completamento del task
		wg.Done()
	}(stringa, carattere, ch)

	// Ricezione del risultato dal canale
	counter = <-ch
	// Attesa del completamento della goroutine
	wg.Wait()
	// Stampa del risultato finale
	fmt.Printf("Stringa: %v\nNumero di caratteri '%v': %v\n", stringa, string(carattere), counter)
}
