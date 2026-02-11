package comm

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func FetchData(delay time.Duration, cancel <-chan struct{}) <-chan string {
	dataChan := make(chan string, 1)

	go func() {
		defer close(dataChan)

		log.Printf(" -- GOROUTINE Iniciada - esperando %v", delay)

		timer := time.NewTimer(delay)
		defer timer.Stop()

		select {
		case <-timer.C:
			select {
			case dataChan <- fmt.Sprintf(" Datos recibidos despues de %v", delay):
				log.Print(" -- EXITO Datos enviados correctamente")

			case <-cancel:
				log.Print(" -- TIMEOUT - Operacion Cancelada.")
			}
		case <-cancel:
			log.Print(" -- TIMEOUT - Operacion NO Completada.")
		}

	}()
	return dataChan
}

func FetchDataLatency(minMs, maxMs int, cancel <-chan struct{}) <-chan string {
	latency := time.Duration(rand.Intn(maxMs-minMs)+minMs) * time.Millisecond
	log.Printf(" -- LATENCIA Simulada: %v", latency)
	return FetchData(latency, cancel)
}

func RespuestaTimeOut(opDelay, timeout time.Duration) (string, bool, bool) {

	cancel := make(chan struct{})
	defer close(cancel)

	dataChan := FetchData(opDelay, cancel)
	select {
	case data, ok := <-dataChan:
		if !ok {
			log.Print(" -- ERROR Canal cerrado inesperadamente")
			return "", false, false
		}
		log.Printf(" -- EXITO - %s", data)
		return data, true, false

	case <-time.After(timeout):
		log.Printf(" -- TIMEOUT - Operacion Excedida: %v", timeout)
		return "", false, true

	}
}
