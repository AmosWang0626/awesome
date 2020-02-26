package sort_utils

const ArrLength = 80000

/*
冒泡排序
*/
func BubbleSorting(arr *[ArrLength]int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		temp := 0 // 加个标示, 如果执行下边循环，没有进行过交换，就说明数组有序了
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				temp++
			}
		}
		if temp == 0 {
			break
		}
	}

	//fmt.Println("排序结果", arr)
}

/*
插入排序
*/
func InsertSorting(arr *[ArrLength]int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		inx := i - 1
		val := arr[i]
		for inx >= 0 && val < arr[inx] {
			arr[inx+1] = arr[inx]
			inx--
		}
		if inx+1 != i {
			arr[inx+1] = val
		}
	}

	//fmt.Println("排序结果", arr)
}

/*
选择排序
*/
func SelectSorting(arr *[ArrLength]int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {

		minIndex := i // 最小数坐标
		for j := i; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if minIndex == i {
			continue
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	//fmt.Println("排序结果", arr)
}

/*
快速排序
*/
func QuickSorting(arr *[ArrLength]int) {
	sort(0, len(arr)-1, arr)
	//fmt.Println("排序结果", arr)
}

// 确定枢轴位置，一分为二，左右分别递归排序
func sort(low, high int, arr *[ArrLength]int) {
	if low < high {
		pivotPoint := partition(low, high, arr)
		sort(low, pivotPoint-1, arr)
		sort(pivotPoint+1, high, arr)
	}
}

// 单趟排序，确定枢轴位置
func partition(low, high int, arr *[ArrLength]int) int {
	pivotVal := arr[low] // 默认取第一个数
	for low < high {
		for low < high && arr[high] >= pivotVal {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] <= pivotVal {
			low++
		}
		if high == low { // 优化
			break
		}
		arr[high] = arr[low]
	}

	arr[low] = pivotVal
	return low
}
