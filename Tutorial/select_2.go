package main

import (
	"fmt"
	"time"
)

// Definizione delle costanti per i livelli di log
const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

// Definizione della struttura logEntry per rappresentare una voce di log
type logEntry struct {
	time     time.Time // Timestamp della voce di log
	severity string    // Livello di gravità (INFO, WARNING, ERROR)
	message  string    // Messaggio di log
}

// Creazione di un channel buffered per memorizzare le voci di log
var logCh = make(chan logEntry, 50)

// E' un channel che trasmette segnali usando struct{}. Non trasporta dati,
// ma serve solo per notificare eventi ( signal only channel ).
// È una pratica comune per segnalare eventi senza sprecare memoria (zero memory allocation).
var doneCh = make(chan struct{})

func main() {
	// Avvio della Go routine per il logger
	go logger()

	// Invio di due voci di log al channel
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is about to shutdown"}

	doneCh <- struct{}{} // serve per chiudere il logger in modo controllato.

	// "struct{}{}" È una struct vuota che non contiene campi. In Go, questa struttura non
	// richiede allocazione di memoria, rendendola ideale per canali usati solo come segnali.
	// "struct{}{}" -> "struct{}" struct vuota senza campi, "{}" inizializzo la struct
}

// Funzione logger che gestisce le voci di log dal channel
func logger() {
	for {
		// Il select blocca l'esecuzione fino a quando non riceve
		// un messaggio da uno dei channel monitorati.
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}

// Un ultimo concetto importante è l'uso del default in un select statement. Se si aggiunge un
// caso default, il select non è più bloccante. In pratica:
//
// - Se c'è un messaggio pronto su uno dei channel monitorati, il select esegue il
//   caso corrispondente.
// - Se nessun messaggio è disponibile, viene eseguito il blocco default.
//
// Questo è utile per creare un select non bloccante, ad esempio per eseguire altre operazioni
// se nessun messaggio è disponibile. Senza il default, il select rimane in attesa indefinita fino
// a quando non arriva un messaggio.
// In sintesi, il default permette di gestire situazioni in cui non si vuole bloccare
// l'esecuzione del programma.
