package sort

// recursive bubblesort

func recursiveBubbleSort(arr []int32, n int) []int32 {
	if n == 1 {return arr}
	for i:=0; i < n-1; i++{
		if arr[i] > arr[i+1] {
			arr[i],arr[i+1] = swap(arr[i],arr[i+1])
		}	
	}
	return recursiveBubbleSort(arr, n-1)
}

func iterativeBubbleSort(arr []int32) []int32 {
	for i:=0; i < len(arr); i++ {
		for j:=0; j < len(arr) - i - 1; j++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = swap(arr[i], arr[i+1])
			}
		}
	}
	return arr
}

func swap(x,y int32) (int32, int32) {
	aux := x
	x = y
	y = aux
	return x,y
}
/*

func main() {
	arr := []int32{5,4,3,2,1,0,6}
	recursiveArr := recursiveBubbleSort(arr, len(arr))
	iterativeArr := iterativeBubbleSort(arr)
	fmt.Println(recursiveArr, iterativeArr)
}

*/