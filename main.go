package main

import (
	"fmt"
	"strconv"
)

var y int = 3 // al package level posso definire le variabili solo in questo modo

// Posso definire più variabili al package level senza utilizzando un blocco 'var', queste variabili
// NON devono per forza essere correlate, questa NON è una struct, sto semplicemente evitando di
// scrivere "var" per definire ogni variabile, scrivendolo una sola volta
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
// Basta che solo la prima lettere sia maiuscola perchè sia visualizzabile globalmente dagli altri pacchetti
var Papere int = -1

// Inoltre c'è una convenzione (è preferibile rispettare questa cosa ma non è necessaria) che le
// variabili con nomi lunghi rappresentano variabile che vivono per più tempo, quindi una variabile
// con un nome corto 'dovrebbe' vivere meno rispetto ad una variabile che ha un nome lungo e descrittivo

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

	// Esistono anche metodi per convertire un tipo in un altro
	b := float32(i)
	fmt.Printf("%v, %T\n", b, b)
	b2 := int(b) // Questa conversione fa perdere informazioni, ad esempio le cifre dopo la virgola
	fmt.Println(b2)

	b3 := string(i) // Questa conversione converte i in una stringa, però in formato ASCII
	fmt.Printf("%v, %T\n", b3, b3)

	// Per convertire i in una stringa vera e propria, quindi che diventi "42" allora bisogna usare
	// il metodo strconv.Itoa() del pacchetto strconv.
	// Itoa() sta per "I" -> integer, to "a" -> ascii
	s := strconv.Itoa(i)
	fmt.Printf("%v, %T\n", s, s)

	// NON si possono fare conversioni implicite, bisogna SEMPRE specificare la conversione

	// Si può definire una variabile booleana per esempio in questo modo
	bo := 1 == 2
	fmt.Printf("%v, %T\n", bo, bo)

	// Quando una variabile primitiva è definita essa è inizializzata a 0
	var boo bool // viene inizializzata a false (= 0)
	fmt.Printf("%v, %T\n", boo, boo)

	// Esistono vari tipi di interi: int8, int16, int32, int64, uint8, uint16, uint32
	// Però non si possono fare operazioni tra interi diversi, bisogna sempre specificare la conversione
	// quindi bisogna convertire un intero nel tipo dell'altro
	// var n1 int = 5
	// var n2 int8 = 3
	// fmt.Println(n1 + n2) // Dà errore, a meno che non converta uno dei due

	// Operatori booleani:
	// &	AND
	// |	OR
	// ^	OR esclusiva
	// &^	AND NOT
}
