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
	///
	///
	///
	///
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	// posso essere inizializzate più variabili nel for
	for i, j := 0, 0; i < 5; i, j = i+1, j+3 { // i += 1, j += 3
		fmt.Printf("i: %v\n", i)
		fmt.Printf("j: %v\n", j)
	}

	g := 0
	// il for fa solo la comparazione (codice più pulito da vedere)
	for g < 3 { // è uguale a scrivere "for ; i < 3 ; {"
		fmt.Println(g)
		g++
	}

Pippo:
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if j > 2 {
				// break // interrompe solo il ciclo più interno, ovvero "for j := 0; j < 5; j++ {"
				break Pippo // in questo modo termina tutto il loop esterno che si trova su di Pippo
			}
			fmt.Printf("i: %v, j: %v\n", i, j)
		}
	}

	s := []int{6, 3, -8}
	fmt.Printf("s: %v\n", s)
	// For loop in range, funziano anche per gli array
	for k, v := range s { // k = chiave, v = valore
		fmt.Printf("s[%v]: %v\n", k, v)
	}

	// posso iterare anche le mappe
	for k, v := range statePopulations {
		fmt.Printf("s[\"%v\"]: %v\n", k, v)
	}

	str := "Hello Go!"
	// anche le stringhe
	for k, v := range str {
		fmt.Printf("str[%v]: %v\n", k, string(v))
	}

	// in questo modo prendo solo i VALORI
	for _, v := range statePopulations {
		fmt.Printf("%v\n", v)
	}
	// in questo modo prendo solo le CHIAVI
	for k := range statePopulations {
		fmt.Printf("%v\n", k)
	}
}
