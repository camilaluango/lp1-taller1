package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Objetivo: Implementar Productor–Consumidor con canales.
// Un productor genera N valores y los envía por un canal; varios consumidores los procesan.
// Practicar cierre de canal y uso de WaitGroup.

func productor(n int, out chan<- int) {
	defer close(out) 
	for i := 1; i <= n; i++ {
		v := rand.Intn(100)
		fmt.Printf("[productor] envía %d\n", v)
		out <- v
		time.Sleep(time.Duration(rand.Intn(400)+100) * time.Millisecond) // simula trabajo}
}
}

func consumidor(id int, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("[consumidor %d] recibe %d\n", id, v)
		time.Sleep(time.Duration(rand.Intn(400)+100) * time.Millisecond) // simula trabajo
		
	}
	fmt.Printf("[consumidor %d] canal cerrado, termina\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	valores := 10
	consumidores := 3

	ch := make(chan int, 4) 

	var wg sync.WaitGroup
	wg.Add(consumidores)
	for i := 1; i <= consumidores; i++ {
		go consumidor(i, ch, &wg)

	}

	go productor(valores, ch)

	wg.Wait()
	fmt.Println("Listo: todos los consumidores terminaron.")
}
