package main 

func quickSort(arr []int, left int, right int) {
	i := left
	j := right
	tmp := 0
	pivot := arr[(left + right) / 2];

	for ; i <= j; {
		for ; arr[i] < pivot;  {
			i++
		}

		for ; arr[j] > pivot; {
			j--;
		}


		if i <= j {
			tmp = arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
			i++;
			j++
		}
	}

	if (left < j) {
		quickSort(arr, left, j)
	}
	if (i < right) {
		quickSort(arr, i, right)
	}
}
