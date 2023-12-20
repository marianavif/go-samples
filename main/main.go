package main 

import (
	
	"fmt"
	sort "github.com/marianavif/go-samples/sort"

)

func main() {
	arr := []int{5,4,-3,2,-1,0,6}
	recursiveArr := sort.RecursiveBubbleSort(arr, len(arr))
	iterativeArr := sort.IterativeBubbleSort(arr)
	cocktailArr := sort.CocktailSort(arr)
	fmt.Println(recursiveArr, iterativeArr, cocktailArr)
}