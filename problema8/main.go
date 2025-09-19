package main

import (
	"fmt"
	"time"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		time.Sleep(500 * time.Millisecond)
		ch <- x * x
	}()
	return ch
}

func fanIn(chs ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var done = make(chan struct{})
		var total = len(chs)

		for _, ch := range chs {
			go func(c <-chan int) {
				for v := range c {
					out <- v
				}
				done <- struct{}{}
			}(ch)
		}

		// esperar que todas las goroutines terminen
		for i := 0; i < total; i++ {
			<-done
		}
	}()
	return out
}

func main() {
// ✅ Crear varios futuros
	f1 := asyncCuadrado(3)
	f2 := asyncCuadrado(4)
	f3 := asyncCuadrado(5)

	fmt.Println("=== Opción 1: esperar cada futuro secuencialmente ===")
	fmt.Println("f1:", <-f1)
	fmt.Println("f2:", <-f2)
	fmt.Println("f3:", <-f3)

	// ✅ Crear nuevos futuros para fan-in
	f4 := asyncCuadrado(6)
	f5 := asyncCuadrado(7)
	f6 := asyncCuadrado(8)

	fmt.Println("=== Opción 2: fan-in ===")
	for v := range fanIn(f4, f5, f6) {
		fmt.Println("resultado fan-in:", v)
	}

}
