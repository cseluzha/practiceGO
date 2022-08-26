package main

import "fmt"

func main() {
	numeros := []int{2, 3, 1, 3, -8, 6, 6, 3, 4, 7, 9}
	quickSort(numeros)
	fmt.Println(numeros)
}

func partition(arr []int) (primeIdx int) {
	primeIdx = 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[len(arr)-1] {
			arr[i], arr[primeIdx] = arr[primeIdx], arr[i]
			primeIdx++
		}
	}
	arr[primeIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[primeIdx]
	return
}

func quickSort(arr []int) {
	if len(arr) > 1 {
		primeIdx := partition(arr)
		quickSort(arr[:primeIdx])   // Izquierda
		quickSort(arr[primeIdx+1:]) // Derecha
	}
}
