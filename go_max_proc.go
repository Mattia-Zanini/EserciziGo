package main

import (
	"fmt"
	"runtime"
)

// GOMAXPROCS controlla il numero di thread del sistema operativo disponibili per le goroutine.
// Per impostazione predefinita, Go assegna un numero di thread pari ai core disponibili sulla
// macchina. Tuttavia, questo valore può essere modificato per migliorare le prestazioni.

//   - Impostare GOMAXPROCS a 1 forza l'esecuzione su un singolo thread, eliminando il parallelismo
//     ma mantenendo la concorrenza.

//   - Aumentare GOMAXPROCS può migliorare le prestazioni, ma valori troppo alti possono introdurre
//     overhead di memoria e problemi di scheduling.

//   - Usare GOMAXPROCS(-1) permette di interrogare il valore attuale senza modificarlo.

// In fase di sviluppo, è consigliabile usare GOMAXPROCS > 1 per individuare le race condition.
// Prima della produzione, testare con diversi valori per trovare il compromesso ottimale tra
// prestazioni e stabilità.
func main() {
	runtime.GOMAXPROCS(100)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}
