package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

var counter int = 0
var globalWG = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	go sayHello()                      // Esegue sayHello in una Go routine
	time.Sleep(100 * time.Millisecond) // Attesa per permettere alla Go routine di completare

	// Abbiamo istruito la funzione main per avviare una Go routine, ma l'applicazione termina non
	// appena la funzione main completa la sua esecuzione. Di conseguenza, non appena la Go routine
	// viene avviata, il programma termina senza dare alla funzione sayHello il tempo di eseguire e
	// stampare il messaggio.

	//

	// Il codice utilizza una funzione anonima eseguita in una goroutine, che accede alla variabile
	// msg grazie ai closures, un concetto di Go (e di molti altri linguaggi) che permette alle
	// funzioni di "catturare" e utilizzare variabili definite nel loro scope esterno.
	//
	// Questo permette alla goroutine di usare msg anche se viene eseguita su
	// uno stack separato. Tuttavia, ciò crea una dipendenza tra msg e la goroutine, con il rischio
	// che main termini prima che la goroutine completi l'esecuzione.
	// La time.Sleep(100 * time.Millisecond) evita il problema, ma in programmi più complessi sarebbe
	// meglio gestire la sincronizzazione con sync.WaitGroup o canali per garantire un corretto
	// accesso ai dati condivisi.
	var msg string = "Ciao"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)

	//

	// Il problema del codice è che la goroutine dipende dalla variabile msg, ma il valore di msg può
	// cambiare prima che la goroutine la usi. Poiché il Go scheduler non interrompe subito il thread
	// principale, msg viene riassegnata a "Pera" prima che la goroutine stampi il suo valore,
	// causando un race condition. Creare una dipendenza da variabili esterne nelle goroutine è
	// rischioso perché il comportamento non è prevedibile, portando a risultati incoerenti.
	fmt.Println("\nSecondo messaggio")
	msg = "Mela"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Pera"
	time.Sleep(100 * time.Millisecond)

	//

	// Per evitare il problema della race condition, si può passare il valore della variabile come
	// argomento alla goroutine. In questo modo, il valore viene copiato, evitando dipendenze dalla
	// variabile originale (viene copiato perchè una stringa viene passata per valore alla funzione).
	// Questo garantisce che la goroutine stampi sempre il valore previsto al momento della chiamata.
	// Tuttavia, l'uso di time.Sleep per sincronizzare l'esecuzione non è una buona pratica, perché
	// lega le prestazioni dell'applicazione al tempo reale, rendendola meno affidabile.
	fmt.Println("\nTerzo messaggio")
	msg = "Anguria"
	go func(m string) {
		fmt.Println(m)
	}(msg)
	msg = "Melone"
	time.Sleep(100 * time.Millisecond)

	//
	//

	// Un'alternativa migliore è utilizzare un WaitGroup della libreria sync, che permette di
	// sincronizzare più goroutine tra loro.
	//
	// Nel codice, il WaitGroup viene inizializzato e incrementato "Add(1)" per indicare che una
	// goroutine deve essere attesa. All'interno della goroutine, dopo aver stampato il valore di msg,
	// viene chiamato Done(), che segnala la conclusione dell'esecuzione e decrementa il contatore del
	// WaitGroup. Infine, il Wait() nel main assicura che il programma non termini prima che la
	// goroutine abbia completato la sua esecuzione.
	//
	// Questa soluzione garantisce che il valore passato alla goroutine sia quello corretto e evita
	// problemi di temporizzazione legati all'uso di sleep.
	var wg = sync.WaitGroup{}
	fmt.Println("\nQuarto messaggio")
	msg = "Uva"
	wg.Add(1)
	go func(m string) {
		fmt.Println(m)
		wg.Done()
	}(msg)
	msg = "Vino"
	wg.Wait()

	// Il codice crea una race condition perché le goroutine sayHello2() e Increment() accedono alla
	// variabile counter senza sincronizzazione, causando un comportamento imprevedibile. L'output
	// varia a ogni esecuzione, con numeri ripetuti o saltati.
	//
	// Problema: Nessuna sincronizzazione tra le goroutine che leggono e scrivono counter.
	// Soluzione: Usare mutex, canali, o operazioni atomiche per garantire accesso sicuro alla variabile.
	fmt.Println("\nEsampio race\n")
	for i := 0; i < 10; i++ {
		globalWG.Add(2)
		go sayHello2()
		go Increment()
	}
	globalWG.Wait()

	// Leggi "Esempio race 2.md"
	fmt.Println("\nEsampio race 2\n")
	globalWG = sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		globalWG.Add(2)
		go sayHello2()
		go Increment()
	}
	globalWG.Wait()
}

func sayHello2() {
	fmt.Printf("Hello #%v\n", counter)
	globalWG.Done()
}
func Increment() {
	counter++
	globalWG.Done()
}

func sayHello3() {
	m.RLock()
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	globalWG.Done()
}
func Increment2() {
	m.Lock()
	counter++
	m.Unlock()
	globalWG.Done()
}
