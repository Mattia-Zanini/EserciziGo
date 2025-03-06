package main

import "fmt"

func main() {
	// Ho creato una mappa che ha le chiavi di tipo stringa e i valori di tipo int
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}
	fmt.Println(statePopulations)

	// Non posso utilizzare come chiave uno slice
	// m := map[[]int]string{} // Dà errore di compilazione

	// Posso utilizzare però un array come chiave
	m := map[[2]int]string{}
	fmt.Println(m)

	// Posso accedere alle entry tramite la chiave
	fmt.Println(statePopulations["Texas"])
	// Per aggiungere una entry sempre con l'operatore []
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)

	// Per eliminare una entry utilizzo il metodo delete()
	delete(statePopulations, "California")
	fmt.Println(statePopulations)
	// Se si prova ad accedere ad una entry che non esiste o che è stata eliminata si ottiene 0
	fmt.Println(statePopulations["California"]) // = 0

	// Si può interrogare la mappa per capire se una entry esiste ed ottenere il suo valore usando questa sintassi
	californiaPopulations, ok := statePopulations["California"]
	texasPopulations, ok := statePopulations["Texas"]
	fmt.Println(californiaPopulations, ok) // = 0 false
	fmt.Println(texasPopulations, ok)      // = 27862596 true
	// se vogliamo sapere solo se esiste una chiave, il valore booleano è salvato nella variabile ok
	fmt.Println(ok) // = true

	// Come per le slice, le mappe vengono copiate per riferimento
	a := map[int]string{
		1: "ciao",
		2: "pollo",
		3: "riso",
	}
	fmt.Printf("a (prima della creazione e manipolazione di b): %v\n", a)
	b := a
	delete(b, 2)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

}
