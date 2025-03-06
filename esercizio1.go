package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	stringa := "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"
	var carattere rune = 'a'
	var counter int = 0

	wg.Add(1)
	go func(s string, r rune, ch chan<- int) {
		count := 0
		for i := 0; i < len(s); i++ {
			if rune(s[i]) == r {
				count++
			}
		}
		ch <- count
		close(ch)
		wg.Done()
	}(stringa, carattere, ch)

	counter = <-ch
	wg.Wait()
	fmt.Printf("Stringa: %v\nNumero di caratteri '%v': %v\n", stringa, string(carattere), counter)
}
