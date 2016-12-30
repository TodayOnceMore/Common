package Sort

// 通用的接口
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// 冒泡排序
func BubbleSort(data Interface) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		ordered := true
		for j := i + 1; j < n; j++ {
			if data.Less(j, i) {
				data.Swap(i, j)
				ordered = false
			}
		}

		if ordered {
			break
		}
	}
}

// 快速排序
func QuickSort(data []int) {
	quicksort(data, 0, len(data)-1)
}

func quicksort(data []int, start, end int) {
	if start < end {
		i, j, x := start, end, data[start]
		for i < j {
			for (i < j) && (data[j] >= x) {
				j--
			}

			if i < j {
				data[i] = data[j]
				i++
			}

			for (i < j) && (data[i] < x) {
				i++
			}

			if i < j {
				data[j] = data[i]
				j--
			}
		}
		data[i] = x
		quicksort(data, start, i-1)
		quicksort(data, i+1, end)
	}
}

// 直接插入排序
func InsertSort(data Interface) {
	n := data.Len()
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0 && data.Less(j+1, j); j-- {
			data.Swap(j, j+1)
		}
	}
}

// 希尔排序
func ShellSort(data Interface) {
	n := data.Len()
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i += gap {
			for j := i - gap; j >= 0 && data.Less(j+gap, j); j -= gap {
				data.Swap(j, j+gap)
			}
		}
	}
}

// 简单选择排序
func SelectSort(data Interface) {
	n := data.Len()
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if data.Less(j, minIndex) {
				minIndex = j
			}
		}
		data.Swap(i, minIndex)
	}
}

// 堆排序
func HeadSort(data Interface) {
	headSort(data, 0, data.Len())
}

func headSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a
	for index := (hi - 1) / 2; index >= 0; index-- {
		siftDown(data, index, hi, first)
	}

	for index := hi - 1; index > 0; index-- {
		data.Swap(first, first+index)
		siftDown(data, lo, index, first)
	}
}

func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

// 归并排序
func MergeSort(data []int) []int {
	n := len(data)
	if n <= 1 {
		return data
	}

	n2 := n / 2
	left := MergeSort(data[:n2])
	right := MergeSort(data[n2:])
	return merge(left, right)
}

func merge(left, right []int) (data []int) {
	l, r, nl, nr := 0, 0, len(left), len(right)
	for (l < nl) && (r < nr) {
		if left[l] < right[r] {
			data = append(data, left[l])
			l++
		} else {
			data = append(data, right[r])
			r++
		}
	}

	data = append(data, left[l:]...)
	data = append(data, right[r:]...)
	return
}
