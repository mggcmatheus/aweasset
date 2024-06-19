package infra

import "time"

func CurrentMillis() int64 {
	// Obtém a hora atual
	now := time.Now()
	// Converte para timestamp em milissegundos
	millis := now.UnixNano() / int64(time.Millisecond)
	return millis
}
