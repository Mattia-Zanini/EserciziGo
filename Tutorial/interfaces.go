package main

import (
	"bytes"
	"fmt"
	"io"
)

//  In Go, la convenzione di nomenclatura per le interfacce dipende dal numero di metodi che
// contengono. Per le interfacce con un solo metodo, il nome dovrebbe essere il nome del metodo
// seguito dal suffisso "er". Ad esempio, un'interfaccia con un metodo Write si chiamerà "Writer",
// mentre una con un metodo Read si chiamerà "Reader".

// Per le interfacce con più metodi, il nome dovrebbe riflettere il comportamento complessivo
// dell'interfaccia, descrivendo cosa fa. L'importante è che il nome sia significativo e
// rappresenti chiaramente lo scopo dell'interfaccia.

//

// Le interfacce in Go sono tipi che definiscono comportamenti, non dati e al loro interno si
// dichiarano metodi che i tipi devono implementare. A differenza degli struct, che contengono dati,
// le interfacce specificano azioni ( es: Write() ).
// Servono per standardizzare comportamenti tra tipi diversi.
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// La struct ConsoleWriter implementa questa interfaccia in modo implicito, senza
// bisogno di una dichiarazione esplicita. Il metodo Write converte la slice di byte
// in una stringa e la stampa a console.
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// In Go, qualsiasi tipo con metodi può implementare un'interfaccia. Ad esempio, un alias come
// IntCounter per int può avere un metodo Increment che soddisfa l'interfaccia Incrementer.
// Questo mostra che non servono solo struct: qualsiasi tipo personalizzato con metodi può
// implementare interfacce, purché si abbia il controllo sulla sua definizione. Non è però
// possibile aggiungere metodi a tipi primitivi o definiti in altri pacchetti.
type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type Closer interface {
	Close() error
}

// E' possibile comporre interfacce combinando più interfacce esistenti,
// un concetto potente che favorisce la scalabilità.
type WriterCloser interface {
	Writer
	Closer
}
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data) // Scrivo i dati nel buffer
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	// Se il buffer contiene più di 8 byte, li scrive in blocchi di 8
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}
func (bwc *BufferedWriterCloser) Close() error {
	// Svuota il buffer rimanente, scrivendo i dati in blocchi di 8 byte
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}), // Inizializza un nuovo buffer vuoto
	}
}

// Il vantaggio è che nel main, si può creare una variabile di tipo Writer e assegnarle un'istanza
// di ConsoleWriter. Questo permette di chiamare il metodo Write senza conoscere il tipo concreto,
// sfruttando il polimorfismo e senza che il chiamante conosca i dettagli dell'implementazione.
func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello podcast listeners, this is a test"))
	wc.Close()

	// Questo è un tentativo di convertire wc nel tipo concreto *BufferedWriterCloser.
	// Se la conversione ha successo, è possibile accedere ai campi e ai metodi specifici di
	// BufferedWriterCloser, come il buffer interno. Questo approccio è utile quando si ha
	// bisogno di accedere a dettagli specifici di un'implementazione che non sono esposti
	// dall'interfaccia.
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	// Il codice dimostra come eseguire una conversione di tipo sicura in Go. Utilizzando la sintassi
	// "r, ok := wc.(io.Reader)", si tenta di convertire wc (di tipo WriterCloser) in un'interfaccia
	// io.Reader. Se la conversione ha successo, ok è true e r contiene il valore convertito;
	// se fallisce, ok è false e r è nil, evitando errori a runtime. Questo approccio permette di
	// verificare se un tipo implementa un'interfaccia specifica in modo sicuro. Nel codice, la
	// conversione fallisce, ma l'applicazione continua, stampando "Conversione fallita".
	// In questo modo si evita di mandare il sistema in panic
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversione fallita")
	}
	fmt.Println("")

	// L'interfaccia vuota in Go è un'interfaccia senza metodi, definita con interface{}. Qualsiasi
	// tipo, anche i primitivi, può essere assegnato a essa. Tuttavia, non avendo metodi, non si può
	// fare nulla direttamente con la variabile, rendendo necessaria una conversione di tipo o l'uso
	// del pacchetto reflect per identificarne il tipo.
	var obj interface{} = NewBufferedWriterCloser()
	if wc, ok := obj.(WriterCloser); ok {
		wc.Write([]byte("Questo e' il secondo test"))
		wc.Close()
	} else {
		fmt.Println("Conversione fallita")
	}
}
