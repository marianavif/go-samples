package sort

// recursive bubblesort

func RecursiveBubbleSort(arr []int, n int) []int {
	if n == 1 {return arr}
	for i:=0; i < n-1; i++{
		if arr[i] > arr[i+1] {
			arr[i],arr[i+1] = swap(arr[i],arr[i+1])
		}	
	}
	return RecursiveBubbleSort(arr, n-1)
}

// iterative bubblesort

func IterativeBubbleSort(arr []int) []int {
	for i:=0; i < len(arr); i++ {
		for j:=0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = swap(arr[j], arr[j+1])
			}
		}
	}
	return arr
}

// super optimized iterative bubblesort (cocktailsort)

func CocktailSort(arr []int) []int {
	left, hasSwapped := 0, false
	right := len(arr)-1

	for left < right {
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = swap(arr[i], arr[i+1])
				hasSwapped = true
			}
		}
		right -= 1
		
		for i:= right; i > left; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = swap(arr[i], arr[i-1])
				hasSwapped = true
			}
		}
		left += 1

		if !hasSwapped {
			return arr
		} else {
			hasSwapped = false
		}
	}
	return arr
}

func swap(x,y int) (int, int) {
	aux := x
	x = y
	y = aux
	return x,y
}
