package main

import "testing"

func TestSoma(t *testing.T) {
	// Given
	a := 15
	b := 15
	// When

	total := Soma(a, b)
	// Then
	if total != 30 {
		t.Errorf("Resultado da soma eh invalido: Resultado: %d. Esperado: %d", total, 30)
	}
}

func TestSubstracao(t *testing.T) {
	// Given
	a := 15
	b := 3
	//When
	sub := Substracao(a, b)
	//Then
	if sub != 12 {
		t.Errorf("Resultado da substracao eh invalido: Resultado: %d. Esperado: %d", sub, 12)
	}

}

func TestMultiplicacao(t *testing.T) {
	//Given
	a := 3
	b := 2

	//When
	prod := Multiplicacao(a, b)

	//Then
	if prod != 6 {
		t.Errorf("Resultado da multiplicacao eh invalido: Resultado: %d. Esperado: %d", prod, 6)
	}

}
