# Programmazione Concorrente in Go: Le Go Routine

## Introduzione
La programmazione concorrente è uno degli argomenti più discussi in Go, specialmente tra chi sta imparando il linguaggio. In questo documento, esploreremo il concetto di **Go routine** e come queste consentono di creare applicazioni efficienti e altamente concorrenti.

## Creazione di Go Routine
Le Go routine sono il fondamento della concorrenza in Go. Ecco come crearne una:

1. **Creare una Go Routine**:
   - Utilizza la keyword `go` prima di una chiamata di funzione per eseguirla come Go routine.
   - Esempio:
     ```go
     func sayHello() {
         fmt.Println("Hello")
     }

     func main() {
         go sayHello() // Esegue sayHello in una Go routine
         time.Sleep(100 * time.Millisecond) // Attesa per permettere alla Go routine di completare
     }
     ```

2. **Thread Leggeri (Green Threads)**:
   - Le Go routine sono thread leggeri gestiti dal runtime di Go.
   - A differenza dei thread del sistema operativo, che sono pesanti (1MB di stack), le Go routine iniziano con stack piccoli e sono economiche da creare e distruggere.
   - Lo scheduler di Go mappa le Go routine sui thread del sistema operativo, permettendo un uso efficiente delle CPU.

## Sincronizzazione
Quando si lavora con Go routine, la sincronizzazione è cruciale per coordinare i task. Due strumenti chiave sono:

1. **Wait Groups**:
   - Utilizzati per attendere il completamento di un gruppo di Go routine.
   - Esempio:
     ```go
     var wg sync.WaitGroup
     wg.Add(1)
     go func() {
         defer wg.Done()
         sayHello()
     }()
     wg.Wait()
     ```

2. **Mutex**:
   - Prevengono le condizioni di gara assicurando che solo una Go routine alla volta acceda ai dati condivisi.
   - Esempio:
     ```go
     var mu sync.Mutex
     var counter int
     go func() {
         mu.Lock()
         counter++
         mu.Unlock()
     }()
     ```

## Parallelismo
Mentre la concorrenza permette di gestire più task, il parallelismo consente di eseguirli simultaneamente su più core della CPU. Per ottenere il parallelismo:
- Assicurati che la tua applicazione Go utilizzi più core (non limitarti al Go Playground, che usa un solo core).
- Usa Go routine e primitive di sincronizzazione per distribuire il lavoro sui core.

## Best Practices
1. **Evitare Condizioni di Gara**:
   - Usa mutex o channel per proteggere le risorse condivise.
2. **Limitare la Creazione di Go Routine**:
   - Sebbene le Go routine siano leggere, evita di crearne un numero eccessivo.
3. **Utilizzare Strumenti**:
   - Sfrutta il race detector di Go (`-race`) per identificare e risolvere problemi di concorrenza.

## Conclusione
Le Go routine offrono un modo potente ed efficiente per gestire la concorrenza in Go. Comprendendo la sincronizzazione, il parallelismo e le best practice, puoi costruire applicazioni altamente concorrenti e performanti.