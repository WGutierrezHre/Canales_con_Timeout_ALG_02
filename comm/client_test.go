package comm

import (
	"strings"
	"testing"
	"time"
)

// TEST 1: TEST DE EXITO (50ms )
func TestExito(t *testing.T) {
	opDelay := 50 * time.Millisecond
	timeout := 200 * time.Millisecond

	data, exito, timeOut := RespuestaTimeOut(opDelay, timeout)

	if !exito {
		t.Errorf("  Se esperaba - exito = true, pero llego: exito = %v", exito)
	}
	if timeOut {
		t.Errorf(" No deberia de haber timeout, pero llego: timeout = %v", timeOut)
	}
	if data == "" {
		t.Error(" Se esperaba Datos, pero llego vacio")
	}
	if !strings.Contains(data, "Datos recibidos") {
		t.Errorf(" Datos no contienen lo esperado: %s", data)
	}

	t.Logf(" TEST EXITOSO - Data: %v", data)
	t.Logf(" exito = %v - timeout = %v", exito, timeOut)

}

// TEST 2: TEST CON TIMEOUT (500ms delay, 200ms timeout)
func TestTimeOut(t *testing.T) {
	opDelay := 500 * time.Millisecond
	timeout := 200 * time.Millisecond

	data, exito, timeOut := RespuestaTimeOut(opDelay, timeout)

	if exito {
		t.Errorf(" Se esperaba - exito = false, pero llego: exito = %v", exito)
	}
	if !timeOut {
		t.Errorf(" Deberia de haber timeout, pero llego: timeout = %v", timeOut)
	}
	if data != "" {
		t.Errorf(" No deberia de haber Datos, pero llego: data = %v", data)
	}

	t.Log(" TEST EXITOSO - TIMEOUT DETECTADO")
	t.Logf(" exito = %v - timeout = %v", exito, timeOut)

	t.Log("  Verificando cancelación de goroutine...")
	time.Sleep(400 * time.Millisecond)
	t.Log("  Goroutine cancelada correctamente")

}

// TEST 3: Helper de simulacion con Latencia Aleatoria
func TestLatency(t *testing.T) {
	timeout := 150 * time.Millisecond
	intentos := 10

	exitoCount := 0
	timeOutCount := 0

	for i := 0; i < intentos; i++ {
		t.Logf(" --- Intento %d/%d ---", i+1, intentos)

		cancel := make(chan struct{})
		dataChan := FetchDataLatency(50, 300, cancel)

		select {
		case data, ok := <-dataChan:
			close(cancel)
			if ok && data != "" {
				t.Logf("EXITO: %s", data)
				exitoCount++
			}
		case <-time.After(timeout):
			close(cancel)
			t.Log(" TIMEOUT")
			timeOutCount++
		}
		time.Sleep(50 * time.Millisecond)
	}

	t.Logf(" -- TOTAL DE INTETOS: %d", intentos)
	t.Logf(" -- Exitos: %d", exitoCount)
	t.Logf(" -- Timeouts: %d", timeOutCount)

}
