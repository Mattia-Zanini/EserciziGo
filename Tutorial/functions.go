package main

import (
	"fmt"
)

type greeter struct {
	greeting string
	name     string
}

// Un metodo è una funzione associata a un tipo, spesso una struct. Nell'esempio, la struct greeter
// ha due campi (greeting e name), e il metodo greet() viene definito con un receiver (g greeter).
// Questo permette di chiamare il metodo su un'istanza della struct (g.greet()), come se fosse un
// suo comportamento.
// In questo caso, il metodo greet() utilizza un value receiver (g greeter), quindi lavora su una
// copia della struct greeter, non sulla struct originale.
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// Un pointer receiver permette di modificare direttamente i campi della struct.
func (g *greeter) greet2() {
	fmt.Println(g.greeting, g.name)
	g.name = "Piero"
}

func main() {
	greetings("Ciao", "Giancarlo")
	sum(1, 2, 3, 4, 5)
	sum2("The sum is: ", 1, 2, 3, 4, 5)
	fmt.Println("La sdrogo somma è: ", sum4(1, 2, 3, 4, 5, 6, 7))
	fmt.Println("")

	d, err := divide(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// In Go le funzioni possono essere usate come variabili
	f := func() { // VERSIONE VERBOSA: "var f func() = func() {"
		fmt.Println("Hello World!")
	}
	f()
	fmt.Println("")

	var div func(float64, float64) (float64, error)
	div = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Non si può dividere per zero")
		}
		return a / b, nil
	}
	d, err = div(3, 0.0)
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(d)
	fmt.Println("")

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	g.greet2()
	fmt.Println("Il nuovo nome è:", g.name)
}

// Se i parametri sono dello stesso tipo posso scrivere il tipo alla fine e in questo modo indico che
// tutti i parametri separati da virgola prima sono di quel determinato tipo
func greetings(greetings, name string) {
	fmt.Println(greetings, name)
}

// i variadic parameters permettono di passare un numero arbitrario di argomenti a una funzione.
// Nel codice, values ...int indica che tutti gli interi passati vengono raccolti in una slice
// chiamata values, semplificando la gestione di più parametri senza dichiararli singolarmente.
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("La somma è: ", result)
}

// Di variadic parameters ce ne può essere solo uno solo e deve essere l'ultimo
func sum2(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

// In Go, i parametri variadici devono essere gli ultimi della lista, perché il runtime non
// può distinguere dove terminano e dove iniziano eventuali altri parametri se posti dopo di essi.
// func sum3(values ...int, msg string)

func sum4(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// Si possono usare valori di ritorno nominati, che permettono di dichiarare una variabile di risultato direttamente nel tipo di ritorno. Questa variabile è disponibile nella funzione e viene restituita automaticamente.
func sum5(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

// Una funzione può restituire più valori, come un risultato e un errore. Questo evita di usare
// panic quando si verificano condizioni problematiche. Ad esempio, la funzione divide restituisce
// un errore se il divisore è zero, invece di bloccare l'applicazione. Il chiamante verifica
// l'errore prima di usare il risultato, seguendo uno stile idiomatico di Go.
// In questo caso il primo valore indica il risultato dell'operazione, mentre il secondo è un
// errore (nil se tutto è andato bene).
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Non si può dividere per zero")
	}
	return a / b, nil
}
