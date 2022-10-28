package alunos

import "fmt"

/*
func SomaInteiro(m map[string]int64) int64 {
	var soma int64
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}
*/
/*Função que junta as duas (SomaInteiro, SomaFloat)funções numa só*/

func SomaGenerica[T int64 | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	/*fmt.Println(SomaInteiro(map[string]int64{"a": 1, "b": 2, "c": 3}))*/
	/*fmt.Println(SomaFloat(map[string]float64{"a": 1.1, "b": 22.2, "c": 3.1}))*/
	fmt.Println("Hello")
	fmt.Println(SomaGenerica(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SomaGenerica(map[string]float64{"a": 1.1, "b": 22.2, "c": 3.1}))
}
