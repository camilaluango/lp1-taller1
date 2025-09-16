package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Provocar un deadlock con dos mutex y dos goroutines que adquieren
// recursos en orden distinto. Luego evitarlo imponiendo un orden global.
// NOTA: La versión con deadlock se quedará bloqueada: ejecútala, observa y luego cambia a la versión segura.

func deadlock() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("G1: Lock mu1") 
		mu1.Lock()

		time.Sleep(100 * time.Millisecond) // fuerza entrelazado
		fmt.Println("G1: Lock mu2") 
		mu2.Lock()

		fmt.Println("G1: listo")
		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		defer wg.Done()
		fmt.Println("G2: Lock mu2") 
		mu2.Lock()

		time.Sleep(100 * time.Millisecond)
		fmt.Println("G2: Lock mu1") 
		mu1.Lock()

		fmt.Println("G2: listo")
		mu1.Unlock()
		mu2.Unlock()
	}()

	// ADVERTENCIA: esto no retornará por el deadlock
	wg.Wait()
}

func seguroOrdenado() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	// Regla: siempre adquirir mu1 luego mu2 (mismo orden en todos)
	lockEnOrden := func(a, b *sync.Mutex) func() func() {
		// retorna: lock():unlock()
		return func() func() {
			a.Lock()
			b.Lock()
			return func() {
				b.Unlock()
				a.Unlock()
			}
		}
	}

	go func() {
		defer wg.Done()
		unlock := lockEnOrden(&mu1, &mu2)()
		defer unlock()
		fmt.Println("G1: trabajo con mu1->mu2")
		time.Sleep(100 * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		unlock := lockEnOrden(&mu1, &mu2)()
		defer unlock()
		fmt.Println("G2: trabajo con mu1->mu2")
		time.Sleep(100 * time.Millisecond)
	}()

	wg.Wait()
	fmt.Println("Seguro: sin deadlock.")
}

func main() {
	fmt.Println("=== Elige una sección para ejecutar ===")
	
	// deadlock()      // <- provocará interbloqueo
	seguroOrdenado()   // <- versión segura
}
