package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// l'unico modo per creare un channel è utilizzando la funzione make()
	ch := make(chan int) // i channel sono fortemente tipizzati e quindi posso inviare/ricevere dati SOLO del tipo per cui ho creato il channel
	wg.Add(2)
	// Receiving goroutine
	go func() {
		i := <-ch // estraggo dati dal channel
		fmt.Println(i)
		wg.Done()
	}()
	// Sending goroutine
	go func() {
		i := 42
		ch <- i // inserisco un intero nel channel, lo passo per copia
		i = 7
		wg.Done()
	}()
	wg.Wait()

	//

	// fatal error: all goroutines are asleep - deadlock!
	// Il channel creato (ch := make(chan int)) è unbuffered, il che significa che può contenere
	// solo un messaggio alla volta. Il problema nasce perché una Go routine riceve solo un messaggio,
	// mentre tre Go routine tentano di inviarne multipli. Dopo il primo messaggio, le Go routine si
	// bloccano perché non c'è nessuno a ricevere i messaggi successivi, causando un deadlock.
	/*
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		for i := 0; i < 3; i++ {
			wg.Add(2)
			go func() {
				ch <- 42
				wg.Done()
			}()
		}
		wg.Wait()
	*/

	//

	fmt.Println("\nEsempio 2")
	// Queste sono due goroutine che agiscono sia come lettori che come scrittori su un channel.
	// Una Go routine riceve un messaggio (42), lo stampa e ne invia un altro (7), mentre l'altra
	// goroutine invia e riceve messaggi, stampando 7. Questo scambio di messaggi funziona, ma spesso
	// è preferibile assegnare ruoli specifici (lettura o scrittura) a ciascuna Go routine per una
	// gestione più chiara dei channel.
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println("Go #1:", i)
		ch <- 8
		wg.Done()
	}()
	go func() {
		ch <- 21
		fmt.Println("Go #2:", <-ch)
		wg.Done()
	}()
	wg.Wait()

	//

	// SOLUZIONE
	// Il problema del codice precedente riguardava l'uso di channel bidirezionali, che rendevano
	// ambiguo il flusso di dati tra Go routine. La soluzione proposta prevede l'uso di channel
	// unidirezionali ( "<-chan int" per ricevere e "chan<- int" per inviare ), che rendono esplicito
	// il ruolo di ogni Go routine. Questo approccio migliora la chiarezza e la sicurezza del codice,
	// evitando errori come l'invio o la ricezione di dati in modo improprio.

	//Il vantaggio principale è la chiarezza: separando i ruoli di mittente e destinatario, il flusso
	// di dati diventa più intuitivo e meno soggetto a errori. Inoltre, il compilatore di Go gestisce
	// in modo polimorfico i channel: un channel bidirezionale può essere passato come argomento a una
	// funzione che accetta un channel unidirezionale, senza bisogno di conversione esplicita. Questa
	// particolarità è specifica dei channel e semplifica la gestione del flusso di dati tra Go routine,
	// mantenendo il codice sicuro e leggibile.
	fmt.Println("\nEsempio 3")
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 17
		wg.Done()
	}(ch)
	wg.Wait()

	//
	//

	// Il codice risolve il problema di deadlock utilizzando un channel buffered "make(chan int, 50)",
	// che può memorizzare più messaggi. Questo permette di inviare più messaggi senza bloccare il
	// mittente, evitando il deadlock. Tuttavia, il secondo messaggio (23) viene perso perché il
	// ricevitore elabora solo il primo messaggio (42). I channel buffered migliorano l'efficienza
	// in scenari concorrenti, ma non gestiscono automaticamente i messaggi persi. La soluzione è
	// utile per evitare blocchi, ma richiede attenzione per garantire che tutti i messaggi vengano
	// elaborati correttamente.
	fmt.Println("\nEsempio 4")
	// 50 indica la capacità del buffer, questo permette di inviare più messaggi
	// senza bloccare il mittente, finché il buffer non è pieno.
	ch = make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		/*
			// Se aggiungo queste righe allora non perdo il secondo valore
			i = <-ch
			fmt.Println(i)
		*/
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 78
		ch <- 30 // viene perso, però il programma non si blocca
		wg.Done()
	}(ch)
	wg.Wait()

	// I channel buffered in Go sono progettati per gestire situazioni in cui mittente e destinatario
	// operano a frequenze diverse. Ad esempio, in un sistema di acquisizione dati come i sismografi,
	// i sensori potrebbero inviare dati in burst (es. una trasmissione ogni ora) per risparmiare
	// energia. Durante questi burst, il mittente è inondato di dati, mentre il destinatario potrebbe
	// impiegare più tempo per elaborarli. Un channel buffered consente di memorizzare temporaneamente
	// i dati in arrivo, evitando che il mittente si blocchi mentre il destinatario li processa.
	// Questo approccio è ideale quando uno dei due lati (mittente o destinatario) ha bisogno di più
	// tempo per completare le operazioni, mantenendo il flusso di dati fluido ed efficiente.

	//
	//

	// Il codice utilizza un ciclo for range per iterare sui messaggi di un channel, permettendo
	// alla Go routine ricevente di gestire più messaggi. Tuttavia, si verifica un deadlock: il ciclo
	// continua a cercare nuovi messaggi anche dopo che il mittente ha smesso di inviarli, poiché il
	// ricevente non ha modo di sapere quando non ci sono più messaggi in questo caso perchè per
	// esempio in una slice hai un numero finito di elementi (se fai un for range loop si di esso),
	// mentre nel channel potresti avere potenzialmente un numero infinito di messaggi.
	// Questo causa un blocco nella Go routine ricevente, che rimane in attesa di messaggi che non
	// arriveranno. Il problema evidenzia la necessità di un meccanismo per segnalare la fine dei
	// messaggi e far terminare correttamente il ciclo.
	/*
		fmt.Println("\nEsempio 5")
		ch = make(chan int, 50)
		wg.Add(2)
		go func(ch <-chan int) {
			for v := range ch {
				fmt.Println(v)
			}
			wg.Done()
		}(ch)
		go func(ch chan<- int) {
			ch <- 33
			ch <- 1
			wg.Done()
		}(ch)
		wg.Wait()
	*/

	//

	// Il problema del deadlock nel ciclo for range su un channel viene risolto chiudendo il channel
	// dopo che tutti i messaggi sono stati inviati. Utilizzando la funzione close(ch), il mittente
	// segnala al ricevitore che non ci sono più messaggi in arrivo. Quando il ciclo for range rileva
	// che il channel è chiuso, termina correttamente, evitando il deadlock. Questo approccio
	// garantisce che tutte le Go routine si chiudano in modo ordinato, senza blocchi, e che i
	// messaggi vengano elaborati correttamente.
	fmt.Println("\nEsempio 6")
	ch = make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 87
		ch <- 22
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()

	// LATO MITTENTE
	// informazione cruciale nella gestione dei channel: non si può inviare un messaggio su un channel
	// chiuso, quando un channel viene chiuso con "close()", tentare di inviare ulteriori messaggi
	// causa un panic.
	// ATTENZIONE: non esiste un modo per riaprire un channel o verificare se è chiuso prima di
	// inviare messaggi. l'unica soluzione è usare una funzione recover per gestire il panic in caso
	// di errore. Sul lato RICEVENTE, il ciclo for range termina automaticamente quando il channel è
	// chiuso, MA sul lato MITTENTE, la chiusura del channel deve essere gestita con molta cautela per
	// evitare panic.

	//

	// Il ricevente con il for range loop riesce a capire quando il channel viene chiuso, in realtà
	// infatti si possono ottenere due parametri dal channel, utilizzando la sintassi "v, ok := <-ch"
	// si ottiengo i valori all'interno del channel e si può verificare se il channel è aperto o chiuso
	fmt.Println("\nEsempio 7")
	ch = make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		// VERSIONE MANUALE
		// Il ciclo for range è più semplice per scenari standard, mentre la sintassi manuale offre
		// maggiore controllo in situazioni complesse, come quando si gestiscono Go routine separate
		// per ogni messaggio. Entrambi gli approcci sono validi, a seconda delle esigenze.
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 13
		ch <- 99
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
