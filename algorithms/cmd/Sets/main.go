package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	A := mapset.NewSet(1, 2, 3, 4)
	B := mapset.NewSet(4, 5, 6, 7)
	C := mapset.NewSet(3, 4, 6, 7, 10)
	fmt.Println(A.Union(B))
	fmt.Println(A.Union(C))
	fmt.Println(B.Intersect(C))
	fmt.Println(B.Intersect(C).Difference(A))
	fmt.Println(A.Union(B).Union(C).Cardinality())
}
