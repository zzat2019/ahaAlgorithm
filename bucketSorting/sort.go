package bucketSorting

import (
	"fmt"
)

func BucketSorting() {
	var num int
	fmt.Scanln(&num)

	arr := [1000]int{}

	var temp int
	for i := 0; i < num; i++ {
		fmt.Scanln(&temp)
		arr[temp]++
	}

	for k, v := range arr {
		for i := 1; i <= v; i++ {
			fmt.Println(k)
		}
	}
}

func BubbleSort() {
	var num int
	fmt.Scanln(&num)
	arr := make([]int, num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&arr[i])
	}

	for i := 0; i < num-1; i++ {
		for j := 0; j < num-i-1; j++ {
			if arr[j] < arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	fmt.Println(arr)

}
