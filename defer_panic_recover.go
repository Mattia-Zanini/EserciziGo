package main

import (
	"fmt"
	"log"
)

// Il defer solitamente viene utilizzato per liberare/chiudere le risorse
func main() {
	f1()
	fmt.Println("")
	f2()
	fmt.Println("")
	f3()
	fmt.Println("")
	f4()
	fmt.Println("")

	// In Go NON esistono le eccezioni come in molti altri linguaggi, perché situazioni spesso
	// considerate eccezionali altrove sono viste come normali in un'applicazione Go.
	// Ad esempio, se si prova ad aprire un file che non esiste, questo è un caso prevedibile e
	// gestibile tramite valori di errore anziché eccezioni. Tuttavia, esistono situazioni in cui
	// l'applicazione non può continuare a funzionare: questi casi sono davvero eccezionali. In Go,
	// invece di usare il termine "eccezione", che ha connotazioni specifiche in altri linguaggi, si
	// utilizza il termine "panic", poiché l'applicazione va in uno stato critico e non sa come proseguire.

	/*
		a, b := 1, 0
		ans := a / b // se si esegue, il runtime genera un panic
		fmt.Println(ans)
	*/

	/*
		fmt.Println("start")
		panic("è successo qualcosa") // stampa il messaggio riguardante il motivo del panic
		fmt.Println("end")
	*/

	// I panic in Go non devono essere necessariamente fatali. Diventano tali solo se si propagano
	// fino al runtime di Go, che, non sapendo come gestire un'applicazione in panico, la termina.

	/*
		fmt.Println("start")
		defer fmt.Println("è stato posticipato")
		panic("è successo qualcosa")
		fmt.Println("end")
	*/
	// OUTPUT:
	// start
	// è stato posticipato
	// panic: è successo qualcosa

	// I panic si verificano dopo l'esecuzione delle istruzioni defer.
	// L'ordine di esecuzione è il seguente:
	// prima viene eseguita la funzione principale, poi vengono eseguite le istruzioni defer,
	// successivamente vengono gestiti eventuali panic e, infine, viene gestito il valore di ritorno.

	//

	// La funzione main stampa "start", chiama la funzione panicker e poi stampa "end".
	// La funzione panicker stampa "about to panic", genera un panic e contiene una funzione defer
	// che intercetta il panic usando recover, loggando l'errore prima che l'esecuzione si interrompa.
	// Quando il programma viene eseguito, viene stampato "start" e "about to panic", poi il panic
	// interrompe il normale flusso di esecuzione e attiva la funzione defer. Il recupero con recover
	// permette alla funzione main di continuare l'esecuzione e stampare "end", ma la funzione panicker
	// non riprende dopo il panic. Il punto chiave è che la funzione che genera e gestisce il panic si
	// ferma, mentre le funzioni chiamanti possono continuare, permettendo al programma di non terminare
	// in modo anomalo.
	fmt.Println("start")
	panicker()
	// panicker2()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	// funzione anonima
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

// Se durante il recupero di un panic ci si rende conto che l'errore non può essere gestito,
// è possibile rilanciare il panic. Per farlo, si può semplicemente leggere l'errore e generare
// un nuovo panic con un messaggio personalizzato. In questo caso, all'interno del gestore del
// panic, viene lanciato un nuovo panic, causando l'interruzione del programma prima che main
// possa stampare l'ultima istruzione. Il risultato è che si ottiene il traceback completo
// dell'errore, mostrando il punto esatto in cui il panic si è verificato. Se non è possibile
// gestire un errore, si può quindi rilanciarlo per consentire la gestione a un livello superiore
// dello stack di chiamate.
func panicker2() {
	fmt.Println("about to panic")
	// funzione anonima
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func f1() {
	fmt.Println("f1, start")
	fmt.Println("f1, middle")
	fmt.Println("f1, end")
}
func f2() {
	// Il defer "sposta" le funzioni, espressioni prima del ritorno della funzione (ovvero quando la funzione ritorna un valore)
	fmt.Println("f2, start")
	defer fmt.Println("f2, middle")
	fmt.Println("f2, end")
}
func f3() {
	// il defer fa eseguire le funzioni con la politica LIFO
	defer fmt.Println("f3, start")
	defer fmt.Println("f3, middle")
	defer fmt.Println("f3, end")
}

func f4() {
	a := "start"
	defer fmt.Println(a)
	// stamperà "start", perchè quando uso il defere su una funzione, essa prende i valori degli
	// argomenti al momento di quando viene chiamato il defer, non quando viene eseguito
	a = "end"
}
