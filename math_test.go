package main

import "testing"

func TestSoma(t *testing.T) {
	// Dado
	a := 15
	b := 15
	// Quando

	total := Soma(a, b)
	// Entao
	if total != 30 {
		t.Errorf("Resultado da soma eh invalido: Resultado: %d. Esperado: %d", total, 30)
	}
}
