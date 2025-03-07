package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type cliente struct {
	nome string
}
type veicolo struct {
	tipo string
}

func (c cliente) noleggia() {
	randomNumber := rand.Intn(3)
	autoNoleggiate[veicoli[randomNumber].tipo]++
	m.Unlock()
	fmt.Printf("%v ha noleggiato un/una %v\n", c.nome, veicoli[randomNumber].tipo)
}
func stampa() {}

var (
	veicoli        = [...]veicolo{{tipo: "Berlina"}, {tipo: "SUV"}, {tipo: "Station Wagon"}}
	autoNoleggiate = map[string]int{}
	wg             = sync.WaitGroup{}
	m              = sync.Mutex{}
)

func main() {
	// inizializzo la mappa
	for _, v := range veicoli {
		autoNoleggiate[v.tipo] = 0
	}

	clienti := [...]cliente{
		{nome: "Luca"},
		{nome: "Sophie"},
		{nome: "Paolo"},
		{nome: "Bob"},
		{nome: "Carlo"},
		{nome: "Franco"},
		{nome: "Mario"},
		{nome: "Luigi"},
		{nome: "Leonardo"},
		{nome: "Tomas"},
	}

	for _, c := range clienti {
		wg.Add(1)
		m.Lock()
		go func() {
			c.noleggia()
			wg.Done()
		}()
	}
	wg.Wait()
}
