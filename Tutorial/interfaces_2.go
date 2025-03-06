package main

import "fmt"

/*
Il codice illustra un concetto importante in Go: la differenza tra i receiver dei metodi
(value receiver vs pointer receiver) e come influenzano l'implementazione delle interfacce.

1. Method Set:
   - Quando si definisce un tipo, il suo *method set* include tutti i metodi, indipendentemente
		 dal tipo di receiver (value o pointer).
   - Tuttavia, quando si implementa un'interfaccia, il method set dipende dal tipo di receiver:
     - Se si usa un *value type* per implementare l'interfaccia, solo i metodi con *value
		   receiver* fanno parte del method set.
     - Se si usa un *pointer type*, il method set include sia i metodi con *value receiver* che
		   quelli con *pointer receiver*.

2. Problema nel codice:
   - Nel codice, "myWriteCloser" ha un metodo "Write" con *pointer receiver* e un metodo "Close"
	   con *value receiver*.
   - Quando si tenta di assegnare un'istanza di "myWriteCloser" (non un puntatore) a una variabile
	   di tipo "WriterCloser", Go genera un errore perché il method set non include "Write" (che ha
		 un pointer receiver).

3. Soluzioni:
   - Usare un *pointer type* per implementare l'interfaccia: "var wc WriterCloser = &myWriteCloser{}".
   - Oppure, cambiare il receiver di "Write" in *value receiver* se non è necessario modificare lo
	   stato del tipo.

4. Conclusione:
   - Se un metodo di un'interfaccia ha un *pointer receiver*, l'interfaccia deve essere implementata
	   con un *pointer type*.
   - Se tutti i metodi hanno **value receiver**, è possibile usare sia value type che pointer type.

Questo comportamento è fondamentale per capire come Go gestisce le interfacce e i receiver dei metodi.
*/

func main() {
	var wc WriterCloser = myWriteCloser{}
	fmt.Println(wc)
}

type Writer interface {
	Write([]byte) (int, error)
}
type Closer interface {
	Close() error
}
type WriterCloser interface {
	Writer
	Closer
}
type myWriteCloser struct{}

func (mwc myWriteCloser) Write(data []byte) (int, error) {
	return 0, nil
}
func (mwc myWriteCloser) Close() error {
	return nil
}
