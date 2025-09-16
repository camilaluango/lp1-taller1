package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Implementar un pool de workers que procesa trabajos y retorna resultados.
// Usa un canal de trabajos y otro de resultados. Cierra canales correctamente.

type trabajo struct {
	ID int
	X  int
}

type resultado struct {
	ID       int
	X        int
	Procesado int
}

func worker(id int, jobs <-chan trabajo, results chan<- resultado, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(200 * time.Millisecond)
		r := resultado{
			ID: j.ID, 
			X: j.X, 
			Procesado: j.X * 2,
		}
		fmt.Printf("[worker %d] procesa trabajo %d -> %d\n", id, j.ID, r.Procesado)
		results <- r
	}
	fmt.Printf("[worker %d] no hay más trabajos\n", id)
}

func main() {
	nTrabajos := 12
	nWorkers := 3

	jobs := make(chan trabajo)
	results := make(chan resultado)

	var wg sync.WaitGroup

	wg.Add(nWorkers)
	for i := 1; i <= nWorkers; i++ {
		go worker(i, jobs, results, &wg)

	}

	go func() {
		for i := 1; i <= nTrabajos; i++ {
			jobs <- trabajo{ID: i, X: i}

		}
		close(jobs) // importante: cerrar para que los workers terminen
	}()

	var wgCollect sync.WaitGroup
	wgCollect.Add(1)
	go func() {
		defer wgCollect.Done()
		wg.Wait()      // esperar que todos los workers terminen
		close(results) // entonces cerrar resultados
	}()

	// Consumidor principal de resultados
	for r := range results {
		fmt.Printf("[main] resultado: trabajo %d -> %d\n", r.ID, r.Procesado)
	}

	wgCollect.Wait()
	fmt.Println("Pool finalizado.")
}
