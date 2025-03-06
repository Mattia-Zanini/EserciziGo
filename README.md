# Go

#### Premessa
Questa non sarà la miglior guida per Go però ho cercato di fare del mio meglio

### Ordine dei file da seguire:
(i file si trovano nella cartella tutorial)

1) variabili.go
2) costanti.go
3) array_and_slices.go
4) maps.go
5) structs.go
6) statements_loops.go
7) defer_panic_recover.go
8) defer_panic.md
9) pointers.go
10) functions.go
11) interfaces.go
12) interfaces_2.go
13) go_routine.md
14) go-routine.go & Esempio race 2.md
15) go_max_proc.go
16) go concurrency best practices.md
17) channels.go
18) select_1.go
19) select_2.go

### Introduzione
Go è un linguaggio di programmazione forte e staticamente tipizzato, simile a Java e C++, in cui il tipo di una variabile non può cambiare nel tempo e deve essere definito in fase di compilazione. Pur offrendo alcuni meccanismi per bypassare il sistema di tipi, nella maggior parte dei casi il linguaggio resta rigoroso e sicuro. È stato progettato per essere semplice, con una sintassi che riduce la verbosità tipica di altri linguaggi fortemente tipizzati, grazie anche a un compilatore che deduce molti dettagli per l'utente.

Go si concentra su tempi di compilazione rapidi, favorendo un ciclo design-build-test efficace. È dotato di garbage collection per la gestione automatica della memoria, con continui miglioramenti per ridurre i tempi di pausa, e include primitive per la concorrenza integrate nel linguaggio stesso. Inoltre, il compilatore produce un unico eseguibile standalone che include tutte le dipendenze necessarie, semplificando il deployment e la gestione delle versioni. Infine, una comunità attiva sostiene lo sviluppo del linguaggio, garantendo un continuo progresso e facilitando l'adozione da parte dei nuovi sviluppatori.

## Riassunto

## Cos'è Go?
- Go è un linguaggio di programmazione open-source e multipiattaforma.
- Progettato per creare applicazioni ad alte prestazioni.
- Tipizzato staticamente e compilato per semplicità ed efficienza.
- Sviluppato da Robert Griesemer, Rob Pike e Ken Thompson nel 2007.
- La sintassi di Go è simile a C++.

## A cosa serve Go?
- Sviluppo web (server-side).
- Creazione di programmi basati su rete.
- Sviluppo di applicazioni aziendali multipiattaforma.
- Sviluppo cloud-native.

## Perché usare Go?
- Facile da imparare e usare.
- Esecuzione e compilazione veloci.
- Supporto integrato per la concorrenza.
- Gestione automatica della memoria.
- Compatibile con più piattaforme (Windows, Mac, Linux, Raspberry Pi, ecc.).

## Confronto tra Go, Python e C++

| Caratteristica                   | Go                              | Python                           | C++                             |
|---------------------------------|--------------------------------|--------------------------------|--------------------------------|
| Tipizzazione                     | Statically typed               | Dynamically typed               | Statically typed               |
| Velocità di esecuzione           | Veloce                         | Lento                           | Veloce                         |
| Compilazione                     | Compilato                      | Interpretato                    | Compilato                      |
| Velocità di compilazione         | Veloce                         | Interpretato                    | Lento                          |
| Concorrenza                       | Goroutines & channels         | Nessun supporto integrato       | Threads                        |
| Garbage Collection                | Sì                             | Sì                              | No                             |
| Supporto alla programmazione OOP | No classi e oggetti            | Sì                              | Sì                              |
| Ereditarietà                      | No                             | Sì                              | Sì                              |

## Note
- **Tempo di compilazione** si riferisce alla traduzione del codice in un programma eseguibile.
- **Concorrenza** significa eseguire più operazioni fuori ordine o simultaneamente senza influenzare il risultato finale.
- **Tipizzazione statica** significa che i tipi di variabile sono noti al momento della compilazione.

## Packages

Il sito ufficiale di Go offre una sezione dedicata ai **package**, che contiene la documentazione di tutte le librerie integrate nel linguaggio. Quando si installa Go, si ottengono automaticamente questi pacchetti, che coprono vari ambiti come **archiviazione, crittografia, database, HTML e gestione del traffico di rete**. Tuttavia, **Go non include librerie per GUI**, poiché è principalmente progettato per lo sviluppo di **server e applicazioni web**. Esistono alcuni progetti per l'uso di Go in applicazioni mobili e client-side, ma al momento non sono ufficialmente supportati.

### **Programma "Hello World" in Go**
Ecco un semplice programma **Hello World** scritto in **Go**:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

---

### **Descrizione del programma**
Il programma segue la struttura tipica di un'applicazione in Go:

1. **Dichiarazione del pacchetto (`package main`)**:  
   Ogni programma in Go è organizzato in pacchetti. Il pacchetto **main** è speciale perché indica il punto di ingresso dell'applicazione.

2. **Importazione della libreria (`import "fmt"`)**:  
   La libreria **fmt** (pronunciata "fumpt" nella comunità Go) è utilizzata per la formattazione e l'output di stringhe.

3. **Definizione della funzione `main()`**:  
   La funzione **main** è il punto di avvio del programma. In essa viene chiamata la funzione `fmt.Println()` per stampare `"Hello, World!"` sul terminale.

Se il codice viene eseguito correttamente, stamperà:

```
Hello, World!
```

Se ci fosse un errore di sintassi, il compilatore genererà un messaggio utile per individuare e correggere il problema.

---

In Go, i comandi `go run`, `go build` e `go install` servono per gestire il ciclo di vita di un programma, dalla scrittura del codice alla distribuzione dell'eseguibile. Ecco una descrizione di ciascuno:

1. **`go run`**: Questo comando compila ed esegue immediatamente il codice sorgente specificato, senza generare un file eseguibile permanente. È utile durante lo sviluppo per testare rapidamente il codice.

2. **`go build`**: Compila il codice sorgente e le sue dipendenze, producendo un file eseguibile nel directory corrente. Questo eseguibile può essere distribuito e eseguito indipendentemente dal codice sorgente.

3. **`go install`**: Compila e installa il pacchetto, posizionando l'eseguibile risultante nella directory specificata dalla variabile d'ambiente `GOBIN`. Se `GOBIN` non è impostata, l'eseguibile viene collocato in una directory predefinita all'interno del percorso di Go. Questo consente di eseguire il programma da qualsiasi posizione nel sistema.

4. **`go get url`**: Scarica pacchetti da repository remoti.

---

In Go, i percorsi `/usr/local/go` e `~/go` (dove `~` rappresenta la home directory dell'utente) hanno ruoli distinti:

1. **`/usr/local/go`**: Questo è il percorso predefinito in cui viene installato il runtime di Go e i suoi strumenti principali. Contiene l'intero sistema Go, inclusi compilatore, linker e altre utilità necessarie per lo sviluppo con Go.

2. **`~/go`**: Questo è il percorso predefinito per il workspace dell'utente, definito dalla variabile d'ambiente `GOPATH`. All'interno di questa directory, si trovano tre sottodirectory principali:
   - `src`: contiene il codice sorgente dei tuoi progetti Go.
   - `pkg`: ospita i pacchetti compilati (librerie).
   - `bin`: include gli eseguibili compilati.

Questa struttura facilita l'organizzazione e la gestione dei progetti Go.

È importante notare che, a partire da Go 1.11, è stato introdotto il supporto per i moduli, che permette di gestire le dipendenze e il codice sorgente al di fuori del tradizionale `GOPATH`. Tuttavia, la comprensione della struttura del workspace rimane utile, soprattutto quando si lavora su progetti legacy o si interagisce con strumenti che si basano su `GOPATH`.

---

## Sintassi di Go

Un file Go è composto dai seguenti elementi:
- Dichiarazione del pacchetto
- Importazione dei pacchetti
- Funzioni
- Istruzioni ed espressioni

## Esempio di base
```go
package main
import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

## Spiegazione del codice
- **Riga 1**: Ogni programma Go è parte di un pacchetto. Il programma appartiene al pacchetto `main`.
- **Riga 2**: `import "fmt"` consente di importare il pacchetto `fmt`.
- **Riga 4**: `func main()` definisce la funzione principale del programma.
- **Riga 5**: `fmt.Println()` stampa a schermo il testo specificato.

📌 **Nota**: In Go, qualsiasi codice eseguibile deve appartenere al pacchetto `main`.

---

## Istruzioni in Go
Un'istruzione in Go termina con un *a capo* oppure con un punto e virgola `;`.

Esempio:
```go
package main
import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

La parentesi graffa `{` non può trovarsi all'inizio di una riga.

---

## Codice compatto
È possibile scrivere codice più compatto, ma non è raccomandato per leggibilità:
```go
package main; import "fmt"; func main() { fmt.Println("Hello World!") }
```
---

## Variabili e Tipi in Go

### Dichiarazione delle Variabili
In Go, le variabili possono essere dichiarate sia a livello di pacchetto che a livello locale. 

### Dichiarazione al Package Level
Le variabili dichiarate a livello di pacchetto devono essere definite con `var`:
```go
var y int = 3
```
È possibile dichiarare più variabili utilizzando un blocco `var`:
```go
var (
    nome        string  = "Mario"
    cognome     string  = "Rossi"
    eta         int     = 20
    temperatura float32 = 20.4
)
```
Le variabili dichiarate con iniziale minuscola sono visibili solo nel pacchetto corrente, mentre quelle con iniziale maiuscola sono esportabili globalmente.

Esempi:
```go
var pippo int = 8  // visibile solo nel pacchetto
var P int = 0      // visibile globalmente
```

## Scope e Shadowing
Se una variabile locale ha lo stesso nome di una variabile globale, la variabile locale ha la precedenza (shadowing):
```go
var pippo int = 8
func main() {
    var pippo int = 10 // Nasconde la variabile globale
    fmt.Println(pippo) // Stampa 10
}
```

## Tipi di Dati e Conversioni
Go non supporta conversioni implicite tra tipi, quindi le conversioni devono essere fatte esplicitamente:
```go
var i int = 42
b := float32(i) // Conversione esplicita
fmt.Printf("%v, %T\n", b, b)
```
Le conversioni errate possono causare perdita di dati, ad esempio:
```go
b2 := int(20.5) // Perde la parte decimale
fmt.Println(b2) // Stampa 20
```

### Conversione di Numeri in Stringhe
La conversione diretta non restituisce una stringa leggibile:
```go
b3 := string(i) // Converte in ASCII
fmt.Printf("%v, %T\n", b3, b3)
```
Per ottenere una conversione corretta, si usa `strconv.Itoa()`:
```go
s := strconv.Itoa(i) // Converte in "42"
fmt.Printf("%v, %T\n", s, s)
```

## Booleani e Operatori
Le variabili booleane assumono i valori `true` o `false`. Se non inizializzate, il valore predefinito è `false`:
```go
var boo bool // false di default
fmt.Printf("%v, %T\n", boo, boo)
```
Operatori booleani:
- `&` → AND
- `|` → OR
- `^` → XOR
- `&^` → AND NOT

## Numeri Complessi
Go supporta numeri complessi:
```go
var c complex64 = 1 + 2i
fmt.Printf("%v, %T\n", c, c)
fmt.Printf("Parte reale: %v\n", real(c))
fmt.Printf("Parte immaginaria: %v\n", imag(c))
```

## Stringhe e Array di Byte
Le stringhe sono immutabili e rappresentano caratteri UTF-8:
```go
str := "sdrogo"
// str[4] = "a" // Errore di compilazione
```
È possibile convertirle in array di byte:
```go
bs := []byte(str)
fmt.Printf("%v, %T\n", bs, bs)
```

## Rune
Le rune rappresentano caratteri in UTF-32:
```go
var r rune = 'a'
fmt.Printf("%v, %T\n", r, r)
```