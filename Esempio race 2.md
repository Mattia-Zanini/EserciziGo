Per correggere il problema della race condition, è necessario sincronizzare le operazioni concorrenti. Una possibile soluzione è usare un **mutex** (mutual exclusion), che agisce come un lucchetto per evitare accessi simultanei a una risorsa condivisa.

```go
package main

import (
	"fmt"
	"time"
)

var counter int = 0
var globalWG = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		globalWG.Add(2)
		go sayHello()
		go Increment()
	}
	globalWG.Wait()
}
func sayHello() {
	m.RLock()
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	globalWG.Done()
}
func Increment() {
	m.Lock()
	counter++
	m.Unlock()
	globalWG.Done()
}
```

### **Uso del Mutex e RWMutex**
- Un **mutex semplice** permette solo due stati: **locked** (bloccato) e **unlocked** (sbloccato). Se un thread tenta di accedere alla risorsa mentre è bloccata, deve aspettare che venga sbloccata.  
- Un **RWMutex (Read-Write Mutex)** è una variante che:
  - Permette a più lettori di accedere contemporaneamente alla risorsa.
  - Impedisce la scrittura quando ci sono lettori attivi.
  - Garantisce che quando un thread scrive, nessun altro possa leggere o scrivere fino alla fine della modifica.  

### **Come viene usato nel codice**
- La funzione `sayHello()` legge la variabile `counter`, quindi usa un **read lock** (`m.RLock()`) per garantire che la lettura sia sicura.
- La funzione `Increment()` modifica `counter`, quindi usa un **write lock** (`m.Lock()`) per impedire accessi concorrenti.  
- Dopo ogni operazione, i lock vengono rilasciati (`RUnlock()` per letture, `Unlock()` per scritture).  

### **Problema rimasto**
Anche se il mutex impedisce modifiche concorrenti alla variabile `counter`, il codice presenta ancora un comportamento **non del tutto sincronizzato**, perché la lettura e l’incremento non sono coordinati nel ciclo principale. Di conseguenza, i valori stampati possono risultare fuori ordine o incongruenti, anche se il programma evita errori di accesso concorrente.