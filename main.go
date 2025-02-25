package main

import "fmt"

var y int = 3 // al package level posso definire le variabili solo in questo modo

// Posso definire più variabili al package level senza utilizzando un blocco 'var', queste variabili
// NON sono correlate, questa NON è una struct
var (
	nome        string  = "Mario"
	cognome     string  = "Rossi"
	eta         int     = 20
	temperatura float32 = 20.4
)

var pippo int = 8
var pluto int = 6

// il nome controlla la visibilità di una variabile, per esempio le variabili definite fin'ora
// in minuscolo sono visualizzabili SOLO dal pacchetto main

var P int = 0 // mentre questa variabile è visualizzabile anche dagli altri pacchetti globalmente

func main() {
	fmt.Printf("Hello World!\n")

	// var v int // il compilatore dà errore se viene dichiarata/definita una variabile locale ma non viene utilizzata
	var i int
	i = 42
	var g float32 = 5
	x := 7. // il compilatore può capire che tipo di variabile è x da come la stiamo definendo
	fmt.Println(i)
	fmt.Println(g)
	fmt.Println(x)
	fmt.Printf("%v, %T\n", g, g) // %v: stampa il valore di g, %T: stampa il tipo di g
	fmt.Printf("%v, %T\n", x, x) // Non è possibile definire x come float32 tramite la definizione ":"

	// i := 2 // Non posso ridefinire la variabile che è già stata definita

	// Posso però ridefinire la variabile che ho definito al package level e ha precedenza quella con
	// lo scope più interiore, in questo caso vince la variabile pippo che è dentro il main.
	// Questo fenomeno si chiama SHADOWING, dove la variabile al package level viene nascosta
	var pippo int = 10
	fmt.Printf("pippo: %v\n", pippo)

	// Viene nascosta MA non sostituita, per esempio se faccio cosi
	fmt.Printf("pluto: %v\n", pluto) // stampa 6
	var pluto int = 13
	fmt.Printf("pluto: %v\n", pluto) // stampa 13
}
