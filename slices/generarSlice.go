package slices

import "math/rand"

func GenSliceInt(tamaño int, tope int) []int {
	generado := make([]int, tamaño)
	for i := range generado {
		generado[i] = rand.Intn(tope) + 1
	}
	return generado
}

func GenSliceChar(tamaño int) []byte {
	generado := make([]byte, tamaño)
	for i := range generado {
		generado[i] = byte(rand.Intn(126-32) + 32)
	}
	return generado
}

func GenSliceStr(tamaño int, tamañoStr int) []string {
	generado := make([]string, tamaño)
	for i := range generado {
		generado[i] = string(GenSliceChar(rand.Intn(tamañoStr)))
	}
	return generado
}
