package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Implementar una versión del problema de los Filósofos Comensales.
// Hay 5 filósofos y 5 tenedores (recursos). Cada filósofo necesita 2 tenedores para comer.

type tenedor struct{ mu sync.Mutex }

func filosofo(id int, izq, der *tenedor, wg *sync.WaitGroup) {
	defer wg.Done()

	pensar(id)

	// Orden global: tomar primero el tenedor con menor dirección
	var primero, segundo *tenedor
	if id == 4 {
		primero, segundo = der, izq
	} else {
		primero, segundo = izq, der
	} 
	primero.mu.Lock()
	segundo.mu.Lock()

	comer(id)

	segundo.mu.Unlock()
	primero.mu.Unlock()
	
	fmt.Printf("[filósofo %d] satisfecho\n", id)
}

func pensar(id int) {
	fmt.Printf("[filósofo %d] pensando...\n", id)
	time.Sleep(time.Duration(100+id*50) * time.Millisecond)

}

func comer(id int) {
	fmt.Printf("[filósofo %d] COMIENDO\n", id)
	time.Sleep(200 * time.Millisecond)

}

func main() {
	const n = 5
	var wg sync.WaitGroup
	wg.Add(n)

	// crear tenedores
	forks := make([]*tenedor, n)
	for i := 0; i < n; i++ {
		forks[i] = &tenedor{}
	}

	// lanzar filósofos
	for i := 0; i < n; i++ {
		izq := forks[i]
		der := forks[(i+1)%n]
		go filosofo(i, izq, der, &wg)

	}

	wg.Wait()
	fmt.Println("Todos los filósofos han comido sin deadlock.")
}
