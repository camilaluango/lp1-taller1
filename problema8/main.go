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
		// simular trabajo
		time.Sleep(500 * time.Millisecond)

		ch <- x * x
	}()
	return ch
}

func main() {
	// TODO: crea varios futuros y recolecta sus resultados: f1, f2, f3

	// TODO: Opción 1: esperar cada futuro secuencialmente

	
	// TODO: Opción 2: fan-in (combinar múltiples canales)
	// Pista: crea una función fanIn que recibe múltiples <-chan int y retorna un único <-chan int
	// que emita todos los valores. Requiere goroutines y cerrar el canal de salida cuando todas terminen.
	
}
