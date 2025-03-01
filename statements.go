package main

import (
	"fmt"
)

func main() {
	// l'if deve essere SEMPRE racchiuso tra le {}
	if true {
		fmt.Println("1 code line") // anche se c'è solo una riga di codice bisogna comunque usare le {}
	}

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}
	if pop, ok := statePopulations["Florida"]; ok {
		// questo è l'inizializzatore dell'if, la variabile "ok" viene utilizzata come valore booleano per l'if
		// "pop" e "ok" sono utilizzabili SOLO all'interno dell'if
		fmt.Println(pop)
	}
	// fmt.Println(ok) // Errore
	// fmt.Println(pop) // Errore

	// Lo switch su Go ha la possibilità di eseguire più comfronti in un solo case.
	// Inoltre i "break" sono impliciti
	n := 3
	fmt.Printf("n: %v\n", n)
	switch n {
	case 1, 4, 7:
		fmt.Println("n = 1, 4 o 7")
	case 2, 3, 5:
		fmt.Println("n = 2, 3 o 5")
	default:
		fmt.Println("default case")
	}

	// esistono anche gli inizializzatori
	switch p := 4 + 2; p { // p = 6
	case 1, 4, 7:
		fmt.Println("p = 1, 4 o 7")
	case 2, 3, 5:
		fmt.Println("p = 2, 3 o 5")
	default:
		fmt.Println("default case")
	}

	// Usando questa sintassi i casi possono sovrapporsi
	i := 4
	switch {
	case i <= 4:
		fmt.Println("i <= 4")
	case i <= 7:
		fmt.Println("i <= 7")
	default:
		fmt.Println("default case")
	}

	// Mentre i "fallthrough" sono espliciti e il fallthrough non esegue la logica dei confronti
	// quindi è a carico del programmatore controllare se la logica del programma viene rispettata
	v := 9
	fmt.Printf("v: %v\n", v)
	switch {
	case v <= 10:
		fmt.Println("v <= 4")
		fallthrough
	case v >= 20:
		fmt.Println("v >= 20")
	default:
		fmt.Println("default case")
	} // alla fine vengono stampati entrambi i primi due print

	// In Go, interface{} (nota anche come empty interface) è un'interfaccia speciale che può
	// contenere un valore di qualsiasi tipo. Questo è possibile perché in Go tutte le variabili
	// implementano automaticamente un'interfaccia vuota. Poiché interface{} può contenere qualsiasi
	// tipo, è utile per scrivere codice generico o lavorare con dati di tipo dinamico.
	var x interface{} = 42 // x può contenere un intero
	x = "ciao"             // ora x contiene una stringa
	x = 3.14               // ora x contiene un float64
	fmt.Printf("x: %v, %T\n", x, x)

	// Lo switch di tipo (switch i.(type)) è una struttura che permette di controllare dinamicamente
	// il tipo effettivo di una variabile contenuta in un'interface{}.
	var y interface{} = 1 // i contiene un int
	fmt.Printf("y: %v, %T\n", y, y)
	switch y.(type) {
	case int:
		fmt.Println("y is an int")
	case float64:
		fmt.Println("y is a float64")
	case string:
		fmt.Println("y is string")
	default:
		fmt.Println("y is another type")
	}

}
