package main

import "fmt"

func main() {
	// Un array può contenere solo un tipo di elementi ed ha una lunghezza fissa
	voti := [3]int{18, 24, 30} // [numero di elementi]tipo{ n elementi }

	// non è necessario scrivere il numero di elementi che si inserisce nell'array
	voti2 := [...]int{19, 20, 31} // in questo modo gli dico di creare un array di dimensione giusta per contenere il numero di elementi che ho definito
	fmt.Printf("%v, %T\n", voti, voti)
	fmt.Printf("%v\n", voti2)

	var studenti [3]string // creo un array di 3 elementi, vuoto
	fmt.Printf("%v\n", studenti)
	studenti[0] = "Pino"
	studenti[2] = "Gino"
	fmt.Printf("%v\n", studenti)
	fmt.Printf("Numero di studenti: %v\n", len(studenti))

	// Si possono creare matrici/array multidimensionali usando questa sintassi
	var matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

	var matrix2 [3][3]int
	matrix2[0] = [3]int{1, 1, 1}
	matrix2[1] = [3]int{2, 2, 2}
	matrix2[2] = [3]int{3, 3, 3}

	fmt.Printf("Matrice: %v\n", matrix)
	fmt.Printf("Matrice 2: %v\n", matrix2)

	// Inoltre l'assegnamento, inizializzazione o passaggio per funzione come argomento, di un array
	// viene fatto per valore, quindi viene copiato l'intero array
	a := [...]int{1, 2, 3}
	b := a // sto copiando l'intero array
	b[1] = 5
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	// quando si crea un array, si possono inizializzare solo alcuni elementi
	arr := [5]int{1: 10, 2: 40} // Iniziallizzo solo il secondo e il terzo elemento, gli altri sono a zero
	fmt.Printf("b: %v\n", arr)  // [0 10 40 0 0]

	c := []int{1, 2, 3} // inizializzo uno "slice", che è simile ad un array per certi versi
	fmt.Printf("%v, %T, Lunghezza: %v\n", c, c, len(c))

	// A differenza dell'array, con le slice l'assegnamento, creazione e
	// passaggio per funzione sono per RIFERIMENTO
	d := c
	d[0] = 7
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
}
