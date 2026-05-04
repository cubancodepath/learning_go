package main

import (
	"fmt"
	"math"
	"sync"
)

func lanzarProceso() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 1; j <= 10; j++ {
				ch <- j + (id * 100)
			}
		}(i)
	}

	done := make(chan bool)

	go func() {
		for num := range ch {
			fmt.Printf("Leído: %d\n", num)
		}
		done <- true
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	<-done
	fmt.Println("Proceso finalizado: Todos los valores han sido impresos.")
}

func usingSelect() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var wg sync.WaitGroup

	// 1. Lanzamos las goroutines primero
	wg.Add(2)
	go func() {
		defer wg.Done()
		for j := 1; j <= 10; j++ {
			ch1 <- j + 100
		}
	}()
	go func() {
		defer wg.Done()
		for j := 1; j <= 10; j++ {
			ch2 <- j + 200
		}
	}()

	// 2. Necesitamos cerrar los canales cuando terminen de escribir
	// Lo hacemos en otra goroutine para no bloquear la lectura
	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()

	// 3. El bucle de lectura con select
	for {
		// El select intentará leer de lo que esté disponible
		select {
		case num, ok := <-ch1:
			if !ok {
				ch1 = nil // "Anulamos" el canal para que select lo ignore
			} else {
				fmt.Printf("Leído from ch1: %d\n", num)
			}
		case num, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Printf("Leído from ch2: %d\n", num)
			}
		}

		// Si ambos canales están en nil, ya terminamos
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("Fin del select")
}

func generateMap() map[int]float64 {
	fmt.Println("Generando map...")
	m := make(map[int]float64, 100_000)

	for i := range 100_000 {
		m[i] = math.Sqrt(float64(i))
	}
	fmt.Println("Mapa generado.")
	return m
}

func main() {
	var cachedMapFunc = sync.OnceValue(generateMap)

	for i := 0; i <= 10000; i += 1000 {
		mapa := cachedMapFunc()

		fmt.Printf("Raíz de %d: %.4f\n", i, mapa[i])
	}

	lanzarProceso()
	usingSelect()
}
