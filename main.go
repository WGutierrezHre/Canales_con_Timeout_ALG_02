package main

import (
	"ejercicio3/comm"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)

	// CASO CON TIMEOUT
	data, exito, timeout := comm.RespuestaTimeOut(4*time.Second, 2*time.Second)
	log.Printf(" Resultado: data = %q - exito = %v - timeout = %v", data, exito, timeout)

	time.Sleep(5 * time.Second)

	// CASO SIN TIMEOUT
	data, exito, timeout = comm.RespuestaTimeOut(1*time.Second, 3*time.Second)
	log.Printf(" Resultado: data = %q - exito = %v - timeout = %v", data, exito, timeout)

	time.Sleep(5 * time.Second)

}
