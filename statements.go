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

	v := 9
	switch {
	case v <= 10:
		fmt.Println("v <= 4")
	case v <= 20:
		fmt.Println("v <= 7")
	default:
		fmt.Println("default case")
	}

}
