package main

import "fmt"

func DynamicArrays() {
	//withData := []int{0, 1, 2, 3, 4, 5}

	//withData[1] = 5
	//x := withData[1]
	//fmt.Println(x)

	//newSlice := withData[2:4]
	//fmt.Println(newSlice)

	//newSlice := withData[:2]
	//fmt.Println(newSlice)

	//newSlice := withData[2:]
	//fmt.Println(newSlice)

	//newSlice := withData[:]
	//fmt.Println(newSlice)

	//a := []int{1, 3}
	//a = append(a, 4, 2)
	//fmt.Println(a)

	//nextSlice := []int{100, 101, 102}
	//newSlice := append(withData, nextSlice...)
	//fmt.Println(newSlice)
}

func find(num int, nums ...int) {

	fmt.Printf("type of nums is %T\n", nums)

	for i, v := range nums {

		if v == num {

			fmt.Println(num, "found at index", i, "in", nums)

			return

		}

	}

	fmt.Println(num, "not found in ", nums)

}
