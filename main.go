package main

import (
	"fmt"
	"strconv"
)

func main() {
	// nilai := ValueBesar([]int{1,2,3,8,9,3,2,1})
	nilai := ValueBesar([]int{1,2,3,8,9,2,1})
	// nilai := ValueBesar([]int{1,2})
	// nilai := ValueBesar([]int{1,2,1,2,2,1})
	fmt.Println("nilai terbesar: ", nilai)
}

func ValueBesar(nilai []int) string {
	
	ascending := []int{}
	ascending = append(ascending, nilai[0])
	terakhirasc := ascending[len(ascending)-1]

	if len(nilai) < 2{
		return "invalid"
	}
	
	for i, _ := range nilai {
		if nilai[i] == terakhirasc + 1 {
			ascending = append(ascending, nilai[i])
		}
		terakhirasc = ascending[len(ascending)-1]
	}

	descending := []int{}
	descending = append(descending, nilai[len(nilai)-1])
	terakhirdsc := descending[len(descending)-1]

	for i := len(nilai)-1; i > 0 ; i--{
		if nilai[i] == terakhirdsc + 1{
			descending = append(descending, nilai[i])
		}
		if nilai[i] != terakhirdsc {
			terakhirdsc = nilai[i] 
		}else{
			terakhirdsc = descending[len(descending)-1]
		}
	}

	hasil := 0
	for _, j := range ascending {
		if j == descending[len(descending)-1] {
			hasil = j
		}
	}

	return strconv.Itoa(hasil)
}