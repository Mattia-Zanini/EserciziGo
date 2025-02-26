package main

import "fmt"

// iota è un contatore che si può utilizzare quando si stanno usando costanti numerate e comincia sempre da 0
const (
	// const-block scope
	c = iota // = 0
	d = iota // = 1
	e = iota // = 2
)

// iota è sensibile allo scope, quindi il risultato di iota sarà diverso a seconda degli scope
const f = iota // = 0

// Non dà errore perchè il compilatore è in grado di vedere il pattern di assegnamenti e capirà
// che h = iota e i = iota
const (
	g = iota
	h
	i

	// Posso anche utilizzare il valore di iota, per esempio:
	j = iota + 10 // = 3 + 10 = 13

	// Inoltre applicando questo utilizzo di iota e il fatto che il compilatore capisce quando
	// inserire iota, li si può utilizzare per avere un offset fissato
	k // = 14
	l // = 15
)

// Possiamo sfruttare l'ultimo punto per ottenere delle potenze come costanti, perchè la funzione
// per elevare alla potenza fa parte del package math, però il valore delle costanti deve essere
// noto a tempo di compilazione e visto che si può un po' "giocare" con iota
const (
	_  = iota             // si fa questo stratagemma per evitare il valore 0 di iota
	KB = 1 << (10 * iota) // = 1 * 2^10
	MB                    // = 1 * 2^20
	GB                    // = 1 * 2^30
	TB                    // = 1 * 2^40
)

// Un altro esempio è quello di definire dei flags da inserire in un singolo byte
const (
	isAdmin            = 1 << iota // = 00000001 (rappresentato in un byte)
	isHeadquarters                 // = 00000010
	canSeeFinancials               // = 00000100
	canSeeAfrica                   // = 00001000
	canSeeAsia                     // = 00010000
	canSeeEurope                   // = 00100000
	canSeeNorthAmerica             // = 01000000
	canSeeSouthAmerica             // = 10000000
)

func main() {
	// Ricordo che se la prima lettera è maiuscola allora la variabile/costante viene esportata è sarà
	// visibile nche dagli altri pacchetti, quindi se vogliamo una costante locale semplicemente mettiamo
	// la prima lettere minuscola
	// Una costante deve avere un valore definito a tempo di compilazione, come constexpr di C++
	const myConst int = 4
	// const myConst2 int = myFunc() // Dà errore di compilazione perchè il valore di myFunc() non è definito a tempo di compilazione
	// Posso fare lo shadowing anche di variabili costanti

	// Il compilatore, come per le variabili può capire (più o meno) il tipo di esse
	const a = 42
	var b int16 = 6
	fmt.Printf("%v, %T\n", a+b, a+b)
	// "più o meno" perchè in realtà il compilatore vede il codice in questa maniera:
	// fmt.Printf("%v, %T\n", 42 + b, 42 + b)

	fmt.Printf("c: %v, %T\n", c, c)
	fmt.Printf("d: %v, %T\n", d, d)
	fmt.Printf("e: %v, %T\n", e, e)
	fmt.Printf("j: %v, %T\n", j, j)
	fmt.Printf("k: %v, %T\n", k, k)
	fmt.Printf("l: %v, %T\n", l, l)

	fmt.Printf("KB: %v, %T\n", KB, KB)
	fmt.Printf("MB: %v, %T\n", MB, MB)
	fmt.Printf("GB: %v, %T\n", GB, GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("roles: %b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
}
