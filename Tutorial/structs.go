package main

import (
	"fmt"
)

// Vale sempre la sintassi dei nomi, l'esistenza di questa struct è visibile dai pacchetti esterni
// PERO' non vedono i campi interni della struct perchè iniziano con una lettera minuscola.
// Se volessi rendere visibile dei campi ai pacchetti esterni (considero sempre il caso che la struct
// inizi con la lettera maiuscola) allora basta che nomino una struct con la prima lettera maiuscola
type Doctor struct {
	number     int
	actorName  string
	companions []string
	surname    string
}

func main() {
	// NON è necessario inizializzare tutti i campi della struct
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		// quando assegno una slice (alla variabile di una struct) devo specificare di che tipo è
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		}}
	fmt.Println(aDoctor)
	// Per accedere ad un elemento della struct utilizzo il punto '.'
	fmt.Println(aDoctor.actorName)

	// E' possibile inizializzare una struct utilizzando una sintassi posizionale, però non è molto
	// consigliata in quanto se avvengono dei cambiamenti alla definizione della struct potrebbe
	// portare ad errori logici del programma oppure a situazioni non desiderate e volute.
	// Utilizzando la sintassi posizionale bisogna per forza inizializzare tutti i campi della struct

	// Doctor è una struct che possiede un nome, esistono anche le struct anonime
	anonimousStruct := struct{ name string }{name: "Pippo"}
	fmt.Println(anonimousStruct)

	// A differenza delle mappe e degli slice, le struct vengono passate per copia
	aS2 := anonimousStruct
	aS2.name = "Carlo"
	fmt.Printf("anonimousStruct: %v\n", anonimousStruct)
	fmt.Printf("aS2: %v\n", aS2)

	// In Go NON esiste l'eredità ovvero la "Is a" proprietà, però esiste la composizione
	type Animal struct {
		name   string
		origin string
	}
	type Bird struct {
		Animal   // la struct Animal viene incapsulata/incorporata
		speedKmh int
		canFly   bool
	}

	b := Bird{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speedKmh = 42
	b.canFly = true

	// "_" è un piccolo artificio per definire una variabile che non viene utilizzata
	// A differenza di prima che ho creato una struct vuota e poi assegnato i valori ai vari campi
	// quando devo usare la sintassi per iniziaizzare i campi in questo modo, devo specificare la
	// struct incapsulata/incorporata, dove invece nel caso di prima il singolo assegnamento
	// veniva delegato
	_ = Bird{
		Animal:   Animal{name: "Papero", origin: "Laghetto"},
		speedKmh: 10,
		canFly:   true,
	}
	fmt.Printf("b: %v\n", b)

	// esistono anche i tags per i campi delle struct i quali fornisco una descrizione testuale del campo,
	// un uso tipico è aggiungere specifiche o vincoli per la persistenza o la serializzazione.
	// I tag struct entrano in gioco come una forma di metadati allegati ai campi struct.
	// Non sono utilizzati direttamente da Go stesso, ma forniscono un potente meccanismo per le
	// librerie di terze parti per analizzare le tue struct e comportarsi di conseguenza.

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
}
