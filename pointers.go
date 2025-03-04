package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	var a [3]int = [3]int{8, 2, 23}
	var b *int = &a[0]
	var c *int = &a[1]
	// var d *int = &a[0] + 4 // In Go NON è possibile utilizzare l'aritmetica dei puntatori, a differenza di C e C++
	fmt.Printf("%v %p %p\n", a, b, c)
	// In Go, l'aritmetica dei puntatori è stata esclusa dal linguaggio per mantenere la semplicità,
	// che è uno dei principi fondamentali del suo design. Tuttavia, se si ha assolutamente bisogno di
	// eseguire operazioni su puntatori, è possibile utilizzare il pacchetto "unsafe". Questo pacchetto
	// permette di eseguire operazioni non controllate dal runtime di Go, consentendo manipolazioni
	// avanzate della memoria. Tuttavia, il nome "unsafe" è appropriato, poiché il suo utilizzo può
	// portare a comportamenti non sicuri e deve essere limitato a scenari specifici.

	//

	// un puntatore non inizializzato viene inizializzato automaticamente con il valore "nil"
	var ms *myStruct
	fmt.Printf("%v, %T\n", ms, ms)
	// In Go, oltre all'uso dell'operatore &, possiamo inizializzare un puntatore a un oggetto
	// utilizzando la funzione incorporata new. Tuttavia, con new non possiamo utilizzare la sintassi
	// di inizializzazione degli oggetti, quindi l'oggetto creato avrà tutti i suoi campi impostati
	// ai valori zero predefiniti. Questo significa che new alloca memoria per l'oggetto ma non
	// consente di inizializzarne direttamente i campi.
	ms = new(myStruct)
	(*ms).foo = 3
	fmt.Println((*ms).foo)

	// In Go, non è necessario dereferenziare esplicitamente un puntatore per accedere ai campi di una
	// struttura. Il compilatore riconosce automaticamente che ci riferiamo all'oggetto sottostante,
	// semplificando la sintassi e migliorando la leggibilità del codice.
	ms.foo = 57
	fmt.Println(ms.foo)
}
