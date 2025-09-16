package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Lanzar varias goroutines que imprimen mensajes y esperar a que todas terminen.

func worker(id int, veces int, wg *sync.WaitGroup) {
	defer wg.Done()
	

	for i := 1; i <= veces; i++ {
		fmt.Printf("[worker %d] hola %d\n", id, i)
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	numGoroutines := 4
	veces := 3

	for id := 1; id <= numGoroutines; id++ {
		wg.Add(1)
		go worker(id, veces, &wg)
	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	fmt.Println("Listo: todas las goroutines terminaron.")
}
