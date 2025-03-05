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

// Funzione logger che gestisce le voci di log dal channel
func logger() {
	// Ciclo for range per iterare sui messaggi nel channel
	for entry := range logCh {
		// Stampa formattata della voce di log
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}

func main() {
	// Avvio della Go routine per il logger
	go logger()

	// Invio di due voci di log al channel
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is about to shutdown"}

	// Attesa per permettere al logger di elaborare i messaggi
	time.Sleep(1000 * time.Microsecond)
}

// DOMANDA: Quand è che si chiude la goroutine del logger?
//
// Ricorda, un'applicazione termina non appena l'ultima istruzione della funzione main completa la
// sua esecuzione. Quindi, quando la sleep termina, l'applicazione si chiude e tutte le risorse
// vengono rilasciate, poiché il runtime di Go restituisce tutte le risorse che stava utilizzando
// al sistema operativo.
// Ciò significa che la nostra Go routine del logger viene chiusa forzatamente. Non c'è uno
// spegnimento graduale per questa Go routine. Viene semplicemente interrotta perché la funzione
// main ha terminato.

// Una soluzione possibile è di aggiungere una funzione defer anonima che vada a chiudere il channel,
// ( la si può per esempio inserire dopo la chiamata della goroutine "go logger" ) questo garantisce
// uno spegnimento graduale, chiudendo intenzionalmente il channel e controllando la terminazione
// della Go routine. È una soluzione valida e accettabile per questo caso d'uso.

// Un'altra possibili soluzione è sul file "select_2.go"
