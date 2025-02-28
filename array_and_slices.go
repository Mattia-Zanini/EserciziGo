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

	e := []int{6, 21, 19, 52, 0, 8, -3}
	f := e[:]   // slice of all elements
	g := e[3:]  // slice dal quarto elemento fino alla fine (quarto elemento incluso). [52 0 8 -3]
	h := e[:6]  // slice dei primi 6 elementi (escluso e[6]). [6 21 19 52 0 8]
	i := e[3:6] // slice composta dal quarto, quinto e sesto elemento (escluso e[6]). [52 0 8]
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	// Posso creare uno slice anche da un array
	e2 := [...]int{6, 21, 19, 52, 0, 8, -3} // e2 è un array
	f2 := e2[:]                             // f2 è uno slice
	fmt.Printf("e2: %v, %T\n", e2, e2)
	fmt.Printf("f2: %v, %T\n", f2, f2)

	// Si può creare uno slice tramite il metodo make() che accetta 2 o 3 operandi
	v1 := make([]int, 3)
	fmt.Printf("v1: %v, %T\n", v1, v1)
	v2 := make([]int, 3, 20) // il terzo argomento è la capacità
	fmt.Printf("v2: %v, %T, Lunghezza: %v, Capacità: %v\n", v2, v2, len(v2), cap(v2))
	// Solo gli slice hanno una capacità perchè sono array di dimensione variabile,
	// mentre gli array NON hanno una capacità, hanno solo la lunghezza

	v3 := []int{}
	fmt.Printf("v3: %v, %T, Lunghezza: %v, Capacità: %v\n", v3, v3, len(v3), cap(v3))
	v3 = append(v3, 2) // in questo modo aggiungo alla fine dello slice un elemento, tipo il metodo push_back dei vettori su C++
	fmt.Printf("v3: %v, %T, Lunghezza: %v, Capacità: %v\n", v3, v3, len(v3), cap(v3))
	v3 = append(v3, 3, 6, -1, 0) // tutti gli argomenti dopo il primo rappresentano gli elementi da inserire
	fmt.Printf("v3: %v, %T, Lunghezza: %v, Capacità: %v\n", v3, v3, len(v3), cap(v3))

	// Con il metodo append si possono concatenare le slice usando l'operatore "spread"
	v3 = append(v3, []int{1, 2, 3}...)
	fmt.Printf("v3: %v, %T, Lunghezza: %v, Capacità: %v\n", v3, v3, len(v3), cap(v3))
	v5 := []int{9, -6, 20, 65, 33}
	v3 = append(v3, v5...)
	fmt.Printf("v3: %v, %T, Lunghezza: %v, Capacità: %v\n", v3, v3, len(v3), cap(v3))
	v3 = append([]int{-99, -98}, v3...)
	fmt.Printf("v3: %v, %T, Lunghezza: %v, Capacità: %v\n", v3, v3, len(v3), cap(v3))

}
